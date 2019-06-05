package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"toy-blockchain/blockchain"
)

func main() {
	chain := blockchain.New()

	initialHash := fmt.Sprintf("%x", sha256.Sum256([]byte("initial hash")))

	block := chain.NewBlockWithHash("proof", initialHash)

	for i := 0; i < 10; i++ {
		chain.NewTransaction(fmt.Sprintf("me%d", i), fmt.Sprintf("you%d", i), rand.Intn(1000))
	}

	block2 := chain.NewBlock("more proof")

	for i := 0; i < 5; i++ {
		chain.NewTransaction(fmt.Sprintf("me%d", i), fmt.Sprintf("you%d", i), rand.Intn(1000))
	}

	block3 := chain.NewBlock("yet more proof")

	fmt.Printf("%v\n", block)
	fmt.Printf("%v\n", block2)
	fmt.Printf("%v\n", block3)

	fmt.Printf("%v\n", chain)
}
