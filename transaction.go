package main

import (
	"fmt"
	"strings"
)

// Transaction related ode

type Transaction struct {
	SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	Value                      float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s   Tx Info   %s\n", strings.Repeat("=", 20), strings.Repeat("=", 20))
	fmt.Printf("sender: %s\n", t.SenderBlockchainAddress)
	fmt.Printf("recipient: %s\n", t.RecipientBlockchainAddress)
	fmt.Printf("value: %.1f\n", t.Value)
}
