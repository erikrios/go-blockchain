package blockchain

import (
	"testing"
	"time"

	"github.com/erikrios/go-blockchain/utils"
)

func TestNewBlock(t *testing.T) {
	t.Run("it should return non nil object, when the function is invoked", func(t *testing.T) {
		senderName := utils.RandStr(5) + utils.RandStr(3)
		receiverName := utils.RandStr(10) + utils.RandStr(4)
		amount := 200

		data := NewData(senderName, receiverName, uint64(amount))
		block := NewBlock(utils.RandStr(64), *data, time.Now())
		if block == nil {
			t.Fatal("data shouldn't nil")
		}
	})
}
