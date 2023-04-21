package blockchain

type Blockchain struct {
	Chain []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, "Genesis", "0")
	return &Blockchain{
		Chain: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(len(bc.Chain), data, bc.Chain[len(bc.Chain)-1].Hash)
	bc.Chain = append(bc.Chain, newBlock)
}
