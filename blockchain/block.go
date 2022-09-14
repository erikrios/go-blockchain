package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	PrevHash  string
	Data      Data
	Timestamp time.Time
	Hash      string
}

func NewBlock(
	prevHash string,
	data Data,
	timestamp time.Time,
) *Block {
	block := &Block{
		PrevHash:  prevHash,
		Data:      data,
		Timestamp: timestamp,
	}

	block.Hash = block.CalculateHash()

	return block
}

func (b *Block) CalculateHash() string {
	jsonBytes, _ := json.Marshal(b.Data)
	unixNanoTimestamp := b.Timestamp.UnixNano()

	blockStr := fmt.Sprintf("%s%s%d", b.PrevHash, string(jsonBytes), unixNanoTimestamp)

	h := sha256.New()
	h.Write([]byte(blockStr))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}
