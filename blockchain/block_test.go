package blockchain

import (
	"testing"
	"time"

	"github.com/erikrios/go-blockchain/utils"
)

const hashLen = 64

func TestNewBlock(t *testing.T) {
	t.Run("it should return non nil object, when the function is invoked", func(t *testing.T) {
		senderName := utils.RandStr(5) + utils.RandStr(3)
		receiverName := utils.RandStr(10) + utils.RandStr(4)
		amount := 200

		data := NewData(senderName, receiverName, uint64(amount))
		block := NewBlock(*data, time.Now())
		if block == nil {
			t.Fatal("data shouldn't nil")
		}
	})
}

func TestCalculateHash(t *testing.T) {
	t.Run("it should return 64 byte length, when CalculateHash method is invoked", func(t *testing.T) {
		senderName := utils.RandStr(5) + utils.RandStr(3)
		receiverName := utils.RandStr(10) + utils.RandStr(4)
		amount := 200

		data := NewData(senderName, receiverName, uint64(amount))
		block := NewBlock(*data, time.Now())
		got := block.CalculateHash()

		if len(got) != hashLen {
			t.Fatalf("got %d, want %d", len(got), hashLen)
		}
	})
}
