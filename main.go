package main

import (
	"fmt"
	"time"

	"github.com/erikrios/go-blockchain/blockchain"
	"github.com/erikrios/go-blockchain/utils"
)

func main() {
	myBlockchain := blockchain.NewBlockchain(5)

	firstData := blockchain.NewData(utils.RandStr(5), utils.RandStr(8), 5000)
	firstBlock := blockchain.NewBlock(*firstData, time.Now())

	secondData := blockchain.NewData(utils.RandStr(10), utils.RandStr(15), 1500)
	secondBlock := blockchain.NewBlock(*secondData, time.Now())

	myBlockchain.AddBlock(*firstBlock)
	myBlockchain.AddBlock(*secondBlock)

	// myBlockchain.Chain[1].Data.Amount = 10000
	// myBlockchain.Chain[1].Hash = myBlockchain.Chain[1].CalculateHash()

	fmt.Printf("%+v\n", myBlockchain)

	fmt.Println("Is valid", myBlockchain.IsChainValid())
}
