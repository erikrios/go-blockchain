package blockchain

import "testing"

func TestNewData(t *testing.T) {
	data := NewData("Bob", "Alice", 200)
	if data == nil {
		t.Fatal("data shouldn't nil")
	}
}
