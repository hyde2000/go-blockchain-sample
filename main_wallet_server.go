package main

import (
	"flag"
	"go-blockchain-sample/server"
)

func main() {
	port := flag.Uint("port", 8888, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway")
	flag.Parse()

	app := server.NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
