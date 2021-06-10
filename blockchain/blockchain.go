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

func (self *blockchain) NewBlock(proof string) *block {
	return self.NewBlockWithHash(proof, "")
}

func (self *blockchain) NewBlockWithHash(proof string, previousHash string) *block {
	prevHash := previousHash

	if len(prevHash) == 0 {
		prevHash = Hash(self.LastBlock())
	}

	block := block{
		Index:        len(self.Chain) + 1,
		Timestamp:    time.Now().UnixNano(),
		Transactions: self.CurrentTransactions,
		Proof:        proof,
		PreviousHash: prevHash,
	}

	self.CurrentTransactions = make([]transaction, 0)
	self.Chain = append(self.Chain, block)

	return &block
}

func (self *blockchain) NewTransaction(sender string, recipient string, amount int) int {
	self.CurrentTransactions = append(self.CurrentTransactions, transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	})

	return self.LastBlock().Index + 1
}

func (self *blockchain) LastBlock() *block {
	return &self.Chain[len(self.Chain)-1]
}

func Hash(block *block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", block))))
}
