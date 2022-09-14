package blockchain

import (
	"reflect"
	"time"
)

type Blockchain struct {
	Chain      []Block
	Difficulty uint8
}

func NewBlockchain(difficulty uint8) *Blockchain {
	return &Blockchain{
		Difficulty: difficulty,
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

func (b *Blockchain) GetLatestBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *Blockchain) AddBlock(block Block) {
	prevHash := b.GetLatestBlock().Hash
	block.PrevHash = prevHash
	block.MineBlock(b.Difficulty)
	b.Chain = append(b.Chain, block)
}

func (b *Blockchain) IsChainValid() bool {
	if len(b.Chain) == 0 {
		return false
	}

	if b.Chain[0].PrevHash != createGenesisBlock().PrevHash {
		return false
	}

	if !reflect.DeepEqual(b.Chain[0].Data, createGenesisBlock().Data) {
		return false
	}

	for i := 1; i < len(b.Chain); i++ {
		currentBlock := b.Chain[i]
		prevBlock := b.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}

	return true
}
