package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
}

type Transaction struct {
	ID         string
	Tenant     string
	PropertyID string
	Details    PropertyDetails
}

type PropertyDetails struct {
	Price          float64
	StartDate      string
	EndDate        string
	RentPrice      float64
	MaintenanceFee float64
}

func CreateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(len(block.Transactions)) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateNewBlock(prevBlock Block, transactions []Transaction) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Format(time.RFC3339)
	newBlock.Transactions = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = CreateHash(newBlock)

	return newBlock
}
