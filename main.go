package main

import (
	"fmt"

	"github.com/imran4u/ethereum-block-chain/data"
	p "github.com/imran4u/ethereum-block-chain/parser"
)

func main() {
	// Initialize the TxParser
	parser := data.NewTxParser()

	// Fetch the latest block
	currentBlock, err := p.FetchCurrentBlockNumber()
	if err != nil {
		fmt.Printf("Error fetching block number: %v\n", err)
		return
	}
	parser.CurrentBlock = currentBlock
	fmt.Printf("Current Block: %d\n", currentBlock)

	// Subscribe to an address
	address := "0ximranAbc123"
	if parser.Subscribe(address) {
		fmt.Printf("Address %s subscribed successfully.\n", address)
	}

	// Fetch transactions for the subscribed address
	transactions, err := p.FetchTransactions(address, currentBlock)
	if err != nil {
		fmt.Printf("Error fetching transactions: %v\n", err)
		return
	}

	// Store the transactions in the parser
	parser.Mu.Lock()
	parser.Transactions[address] = transactions
	parser.Mu.Unlock()

	// Output the transactions for the subscribed address
	fmt.Printf("Transactions for address %s:\n", address)
	for _, tx := range transactions {
		fmt.Printf("Transaction Hash: %s, From: %s, To: %s, Value: %s, Block: %d\n", tx.Hash, tx.From, tx.To, tx.Value, tx.Block)
	}
}
