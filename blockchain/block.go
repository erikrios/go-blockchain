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
	Nonce     uint64
}

func NewBlock(
	data Data,
	timestamp time.Time,
) *Block {
	block := &Block{
		Data:      data,
		Timestamp: timestamp,
	}

	block.Hash = block.CalculateHash()

	return block
}

func (b *Block) CalculateHash() string {
	jsonBytes, _ := json.Marshal(b.Data)
	unixNanoTimestamp := b.Timestamp.UnixNano()

	blockStr := fmt.Sprintf(
		"%s%s%d%d",
		b.PrevHash,
		string(jsonBytes),
		unixNanoTimestamp,
		b.Nonce,
	)

	h := sha256.New()
	h.Write([]byte(blockStr))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func (b *Block) MineBlock(difficulty uint8) {
	start := time.Now()

	diffBytes := make([]byte, difficulty)
	for i := range diffBytes {
		diffBytes[i] = '0'
	}

	diffStr := string(diffBytes)

	for {
		b.Hash = b.CalculateHash()

		if hashLeading := b.Hash[0:difficulty]; hashLeading == diffStr {
			break
		}

		b.Nonce++
	}

	elapsed := time.Since(start)

	fmt.Printf("Block mined, took: %s, hash result: %s\n", elapsed, b.Hash)
}
