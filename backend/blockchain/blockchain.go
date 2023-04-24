package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"strings"
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
	Blocks     []*Block
	LatestHash string
}

func NewBlock(prevBlockHash []byte, data string) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func (b *Block) SetHash() {
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, IntToHex(b.Timestamp)}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis Block")
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks:     []*Block{NewGenesisBlock()},
		LatestHash: "",
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.LatestHash = string(newBlock.Hash)
}

func (bc *Blockchain) String() string {
	var lines []string
	for _, block := range bc.Blocks {
		lines = append(lines, block.String())
	}
	return strings.Join(lines, "\n")
}

func (b *Block) String() string {
	return "Timestamp: " + strconv.FormatInt(b.Timestamp, 10) + "\n" +
		"Data: " + string(b.Data) + "\n" +
		"PrevBlockHash: " + string(b.PrevBlockHash) + "\n" +
		"Hash: " + string(b.Hash) + "\n"
}
