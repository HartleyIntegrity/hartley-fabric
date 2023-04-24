package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int          `json:"index"`
	Timestamp    string       `json:"timestamp"`
	TenancyAgree TenancyAgree `json:"tenancyAgree"`
	Hash         string       `json:"hash"`
	PrevHash     string       `json:"prevHash"`
}

type TenancyAgree struct {
	ID          string `json:"id"`
	Property    string `json:"property"`
	Tenant      string `json:"tenant"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
}

func (b *Block) calculateHash() string {
	indexStr := strconv.Itoa(b.Index)
	record := strings.Join([]string{indexStr, b.Timestamp, b.PrevHash}, "|")
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBlock(prevBlock *Block, tenancyAgree TenancyAgree) *Block {
	block := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Format(time.RFC3339),
		TenancyAgree: tenancyAgree,
		PrevHash:     prevBlock.Hash,
	}
	block.Hash = block.calculateHash()
	return block
}
