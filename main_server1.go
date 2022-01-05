package main

import (
	"flag"
	"go-blockchain-sample/server"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := server.NewBlockchainServer(uint16(*port))
	app.Run()
}
