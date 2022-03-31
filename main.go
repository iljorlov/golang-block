package main

import (
	"fmt"

	"github.com/iljorlov/golang-block/blockchain"
)


func main() {

	// chain := blockchain.InitBlockChain()
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First add")
	chain.AddBlock("Second add")
	chain.AddBlock("Third add")

	for _, block := range chain.blocks{
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("===")
	}

}