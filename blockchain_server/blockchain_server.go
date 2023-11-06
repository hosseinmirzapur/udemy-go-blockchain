package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(
		http.ListenAndServe("0.0.0.0:"+fmt.Sprintf("%d", bcs.port),
			nil,
		),
	)
}
