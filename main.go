package main

import (
	"fmt"
	"log"

	"github.com/hosseinmirzapur/goblockchain/block"
	"github.com/hosseinmirzapur/goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchainFunc()
	walletFunc()
}

func blockchainFunc() {
	my_blockchain_address := "MY_BLOCKCHAIN_ADDRESS"
	blockchain := block.NewBlockchain(my_blockchain_address)
	blockchain.Print()
}

func walletFunc() {
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	walletM := wallet.NewWallet() // Miner

	t := wallet.NewTransaction(walletA.PrivKey(), walletA.PubKey(), walletA.BlockchAddr(), walletB.BlockchAddr(), 1.0)

	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchAddr())
	isAdded := blockchain.AddTransaction(
		walletA.BlockchAddr(),
		walletB.BlockchAddr(),
		1.0,
		walletA.PubKey(),
		t.GenSign(),
	)
	fmt.Println("Added?", isAdded)

	// Mining
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchAddr()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchAddr()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchAddr()))
}
