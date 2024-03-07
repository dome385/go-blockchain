package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockChain initialisiert eine neue Blockchain mit einem Block
func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init")
	return bc
}

// CreateBlock fügt der Blockchain einen Block hinzu mit append.
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// NewBlock initialisiert einen neuen Block
func NewBlock(nonce int, previousHash string) *Block {
	// Alternativ folgender Code:
	/* b := new(Block)
	b.timestamp = time.Now().UnixNano()
	return b */
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) Print() {
	fmt.Println("--------------------------------------------")
	fmt.Printf("| timestamp                 %d\n", b.timestamp)
	fmt.Printf("| nonce                     %d\n", b.nonce)
	fmt.Printf("| previous_hash             %s\n", b.previousHash)
	fmt.Printf("| transactions              %s\n", b.transactions)
	fmt.Println("--------------------------------------------")
}

func (bc *Blockchain) PrintBlockChain() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func main() {
	blockChain := NewBlockChain()
	blockChain.PrintBlockChain()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.PrintBlockChain()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.PrintBlockChain()
}
