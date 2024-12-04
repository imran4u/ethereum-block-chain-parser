package data

import (
	"sync"
)

// Transaction represents an Ethereum transaction
type Transaction struct {
	Hash  string // Transaction hash
	From  string // Sender address
	To    string // Receiver address
	Value string // Amount transferred
	Block int    // Block number in which the transaction was confirmed
}

// Parser interface
type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}

// TxParser implements the Parser interface and stores data in memory
type TxParser struct {
	Mu            sync.RWMutex
	CurrentBlock  int
	Subscriptions map[string]bool
	Transactions  map[string][]Transaction
}

// >>> Parser methods
// NewTxParser creates a new TxParser instance
func NewTxParser() *TxParser {
	return &TxParser{
		CurrentBlock:  0,
		Subscriptions: make(map[string]bool),
		Transactions:  make(map[string][]Transaction),
	}
}

// GetCurrentBlock returns the last parsed block number
func (p *TxParser) GetCurrentBlock() int {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.CurrentBlock
}

// Subscribe adds an Ethereum address to the observer list
func (p *TxParser) Subscribe(address string) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	if _, exists := p.Subscriptions[address]; exists {
		return false
	}
	p.Subscriptions[address] = true
	return true
}

// GetTransactions returns the list of transactions for a subscribed address
func (p *TxParser) GetTransactions(address string) []Transaction {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.Transactions[address]
}
