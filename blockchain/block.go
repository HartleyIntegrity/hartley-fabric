package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

func NewBlock(index int, data string, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		Data:         data,
		PreviousHash: previousHash,
		Hash:         "",
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	hashData := string(b.Index) + b.Timestamp + b.Data + b.PreviousHash
	hash := sha256.New()
	hash.Write([]byte(hashData))
	return hex.EncodeToString(hash.Sum(nil))
}
