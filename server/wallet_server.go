package server

import (
	"bytes"
	"encoding/json"
	"go-blockchain-sample/models"
	"go-blockchain-sample/utils"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{
		port:    port,
		gateway: gateway,
	}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := models.NewWallet()
		m, _ := myWallet.MarshallJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	var t TransactionRequest

	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&t); err != nil {
			log.Printf("ERROR: %v\n", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		valueFloat32 := float32(value)
		// Prepare response
		w.Header().Add("Content-Type", "application/json")
		transaction := models.NewWalletTransaction(privateKey, publicKey, *t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, valueFloat32)
		signature, err := transaction.GenerateSignature()
		if err != nil {
			log.Println("ERROR: generate signature error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		signatureStr := signature.String()

		bt := &models.TransactionRequest{
			SenderBlockchainAddress:    t.SenderBlockchainAddress,
			RecipientBlockchainAddress: t.RecipientBlockchainAddress,
			SenderPublicKey:            t.SenderPublicKey,
			Value:                      &valueFloat32,
			Signature:                  &signatureStr,
		}
		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)
		// Do transaction
		res, err := http.Post(ws.Gateway()+"/transactions", "application/json", buf)
		if err != nil {
			log.Println("ERROR: cannot post")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if res.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		} else {
			io.WriteString(w, string(utils.JsonStatus("fail")))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
		return
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
