package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type TenancyAgreement struct {
	ID        string `json:"id"`
	Property  string `json:"property"`
	Landlord  string `json:"landlord"`
	Tenant    string `json:"tenant"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, IntToHex(b.Timestamp)}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
