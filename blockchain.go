package main

import (
	"log"
	"strings"
	"time"
)

// Blockchain related code

type Blockchain struct {
	TransactionPool   []*Transaction `json:"transaction_pool"`
	Chain             []*Block       `json:"chain"`
	BlockchainAddress string         `json:"blockchain_address"`
}

func NewBlockchain(bcAddr string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.BlockchainAddress = bcAddr
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Chain {
		block.Print()
	}
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddTransaction(
	sender string,
	recipient string,
	value float32,
) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.TransactionPool {
		transactions = append(transactions, NewTransaction(t.SenderBlockchainAddress, t.RecipientBlockchainAddress, t.Value))
	}

	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, prevHash [32]byte,
	transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{
		Nonce:        nonce,
		PreviousHash: prevHash,
		Timestamp:    time.Now().UnixNano(),
		Transactions: transactions,
	}
	guessHash := guessBlock.Hash()
	return string(guessHash[:])[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	prevhash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, prevhash, transactions, MINING_DIFFICULTY) {
		nonce++
	}

	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.BlockchainAddress, MINING_REWARD)

	nonce := bc.ProofOfWork()
	prevhash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, prevhash)
	log.Println("action=mining, status=success")

	return true
}

func (bc *Blockchain) CalculateTotalAmount(address string) float32 {
	var totalAmount float32 = 0

	for _, block := range bc.Chain {
		for _, transaction := range block.Transactions {
			if transaction.RecipientBlockchainAddress == address {
				totalAmount += transaction.Value
			}
			if transaction.SenderBlockchainAddress == address {
				totalAmount -= transaction.Value
			}
		}
	}

	return totalAmount
}
