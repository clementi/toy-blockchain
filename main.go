package main

import (
	"crypto/sha256"
	"fmt"
	"toy-blockchain/blockchain"
)

func main() {
	chain := blockchain.New()

	initialHash := fmt.Sprintf("%x", sha256.Sum224([]byte("initial hash")))

	block := chain.NewBlock("proof", initialHash)

	chain.NewTransaction("me", "you", 35)

	block2 := chain.NewBlock("more proof", blockchain.Hash(block))
	block3 := chain.NewBlock("yet more proof", blockchain.Hash(block2))

	fmt.Printf("%v\n", block)
	fmt.Printf("%v\n", block2)
	fmt.Printf("%v\n", block3)

	fmt.Printf("%v\n", chain)
}
