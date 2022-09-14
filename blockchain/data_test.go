package blockchain

import (
	"testing"

	"github.com/erikrios/go-blockchain/utils"
)

func TestNewData(t *testing.T) {
	t.Run("it should return non nil object, when the function is invoked", func(t *testing.T) {
		senderName := utils.RandStr(5) + utils.RandStr(3)
		receiverName := utils.RandStr(10) + utils.RandStr(4)
		amount := 200

		data := NewData(senderName, receiverName, uint64(amount))
		if data == nil {
			t.Fatal("data shouldn't nil")
		}
	})
}
