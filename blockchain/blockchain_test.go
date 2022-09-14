package blockchain

import (
	"reflect"
	"testing"
	"time"
)

const newBlockchainLength = 1

func TestNewBlockhain(t *testing.T) {
	t.Run("it should return non nil object, when the function is invoked", func(t *testing.T) {
		got := NewBlockchain()
		if got == nil {
			t.Fatal("data shouldn't nil")
		}

		if len(got.Chain) != newBlockchainLength {
			t.Fatalf("got %d, want %d", len(got.Chain), newBlockchainLength)
		}
	})
}

func TestCreateGenesisBlock(t *testing.T) {
	t.Run("it should return valid genesis block, when createGenesisBlock msthod is invoked", func(t *testing.T) {
		data := NewData("", "", 0)
		expected := NewBlock(*data, time.Now())

		got := createGenesisBlock()

		if got.PrevHash != expected.PrevHash {
			t.Fatalf("got %s, want %s", got.PrevHash, expected.PrevHash)
		}

		if !reflect.DeepEqual(got.Data, expected.Data) {
			t.Fatalf("got %+v, want %+v", got.Data, expected.Data)
		}
	})
}
