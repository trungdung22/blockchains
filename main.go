package main

import (
	"fmt"

	"./blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first block chain")
	chain.AddBlock("second block chain")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data in Block: %s\n", block.Data)
	}
}
