package blockchain

import (
	"sync"
)

type Blockchain struct {
	Chain []*Block
}

var once sync.Once
var blockchainInstance *Blockchain

func GetBlockchain() *Blockchain {
	once.Do(func() {
		blockchainInstance = &Blockchain{
			Chain: []*Block{genesisBlock()},
		}
	})
	return blockchainInstance
}

func genesisBlock() *Block {
	tenancyAgree := TenancyAgree{
		ID:          "0",
		Property:    "N/A",
		Tenant:      "N/A",
		StartDate:   "N/A",
		EndDate:     "N/A",
		Description: "Genesis Block",
	}
	genesis := &Block{
		Index:        0,
		Timestamp:    "2023-01-01T00:00:00Z",
		TenancyAgree: tenancyAgree,
	}
	genesis.Hash = genesis.calculateHash()
	return genesis
}
