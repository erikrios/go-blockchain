package blockchain

import "time"

type Block struct {
	PrevHash  string
	Data      Data
	Timestamp time.Time
}

func NewBlock(
	prevHash string,
	data Data,
	timestamp time.Time,
) *Block {
	return &Block{
		PrevHash:  prevHash,
		Data:      data,
		Timestamp: timestamp,
	}
}
