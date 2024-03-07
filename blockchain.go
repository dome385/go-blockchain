package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// NewBlockChain initialisiert eine neue Blockchain mit einem Block
func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// CreateBlock fügt der Blockchain einen Block hinzu mit append.
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

// NewBlock initialisiert einen neuen Block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	// Alternativ folgender Code:
	/* b := new(Block)
	b.timestamp = time.Now().UnixNano()
	return b */
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	/* fmt.Println(string(m)) */
	return sha256.Sum256([]byte(m))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (b *Block) Print() {
	fmt.Println("--------------------------------------------")
	fmt.Printf("| timestamp                 %d\n", b.timestamp)
	fmt.Printf("| nonce                     %d\n", b.nonce)
	fmt.Printf("| previous_hash             %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.PrintTransaction()
	}
	fmt.Println("--------------------------------------------")
}

func (bc *Blockchain) PrintBlockChain() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) PrintTransaction() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address     %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient                     %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" Value                     %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func main() {

	// Neue Blockchain erstellen
	blockChain := NewBlockChain()
	/* blockChain.PrintBlockChain() */

	blockChain.AddTransaction("A", "B", 1.0)
	// Letzter Hash dem neuem Block hinzufügen
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	/* blockChain.PrintBlockChain() */

	// Letzter Hash dem neuen Block hinzufügen
	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.PrintBlockChain()

}
