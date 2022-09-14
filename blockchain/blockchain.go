package blockchain

import "time"

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain: []Block{
			createGenesisBlock(),
		},
	}
}

func createGenesisBlock() Block {
	data := NewData("", "", 0)
	block := NewBlock(*data, time.Now())
	return *block
}
