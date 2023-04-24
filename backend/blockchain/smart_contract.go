package blockchain

func (bc *Blockchain) AddBlock(tenancyAgree TenancyAgree) {
	newBlock := CreateBlock(bc.Chain[len(bc.Chain)-1], tenancyAgree)
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) UpdateBlock(id string, tenancyAgree TenancyAgree) {
	// TODO: Implement the function to update a block
}

func (bc *Blockchain) DeleteBlock(id string) {
	// TODO: Implement the function to delete a block
}
