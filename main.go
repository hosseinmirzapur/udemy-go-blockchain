package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	my_blockchain_address := "MY_BLOCKCHAIN_ADDRESS"
	blockchain := NewBlockchain(my_blockchain_address)
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 3.4)
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount("A"))
	fmt.Printf("My %.1f\n", blockchain.CalculateTotalAmount(my_blockchain_address))
}
