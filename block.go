package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Block related Code

type Block struct {
	Nonce        int            `json:"nonce"`
	PreviousHash [32]byte       `json:"previous_hash"`
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.Nonce)
	fmt.Printf("prevHash: %x\n", b.PreviousHash)
	fmt.Printf("timestamp: %d\n", b.Timestamp)
	for _, t := range b.Transactions {
		t.Print()
	}
	fmt.Println(strings.Repeat("-", 20))
}
