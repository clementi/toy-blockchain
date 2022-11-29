package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type block struct {
	Index        int
	Timestamp    int64
	Transactions []transaction
	Proof        string
	PreviousHash string
}

type blockchain struct {
	Chain               []block
	CurrentTransactions []transaction
}

func New() *blockchain {
	return &blockchain{
		Chain:               make([]block, 0),
		CurrentTransactions: make([]transaction, 0),
	}
}

func (chain *blockchain) NewBlock(proof string) *block {
	return chain.NewBlockWithHash(proof, "")
}

func (chain *blockchain) NewBlockWithHash(proof string, previousHash string) *block {
	prevHash := previousHash

	if len(prevHash) == 0 {
		prevHash = Hash(chain.LastBlock())
	}

	block := block{
		Index:        len(chain.Chain) + 1,
		Timestamp:    time.Now().UnixNano(),
		Transactions: chain.CurrentTransactions,
		Proof:        proof,
		PreviousHash: prevHash,
	}

	chain.CurrentTransactions = make([]transaction, 0)
	chain.Chain = append(chain.Chain, block)

	return &block
}

func (chain *blockchain) NewTransaction(sender string, recipient string, amount int) int {
	chain.CurrentTransactions = append(chain.CurrentTransactions, transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	})

	return chain.LastBlock().Index + 1
}

func (chain *blockchain) LastBlock() *block {
	return &chain.Chain[len(chain.Chain)-1]
}

func Hash(block *block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", block))))
}
