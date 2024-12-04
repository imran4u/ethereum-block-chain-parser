package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	cons "github.com/imran4u/ethereum-block-chain/constant"
	data "github.com/imran4u/ethereum-block-chain/data"
)

// MakeJSONRPCRequest sends a JSON-RPC request to the Ethereum endpoint
func MakeJSONRPCRequest(method string, params []interface{}) (map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	resp, err := http.Post(cons.ETHEREUM_RPCURL, "application/json", bytes.NewReader(reqBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// FetchCurrentBlockNumber retrieves the latest block number from the Ethereum network
func FetchCurrentBlockNumber() (int, error) {
	result, err := MakeJSONRPCRequest("eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}

	// The result is a hex string, so we need to parse it
	blockNumberHex := result["result"].(string)
	blockNumber, err := strconv.ParseInt(blockNumberHex[2:], 16, 64)
	if err != nil {
		return 0, err
	}
	return int(blockNumber), nil
}

// FetchTransactions retrieves transactions for a given Ethereum address
func FetchTransactions(address string, block int) ([]data.Transaction, error) {
	// This is a simplified fetch, in a real scenario, you'd call the "eth_getLogs" or "eth_getTransactionByBlockNumber" API
	// For now, we'll simulate fetching transactions
	transactions := []data.Transaction{
		{
			Hash:  "0x123",
			From:  "0xabc",
			To:    address,
			Value: "1000000000000000000", // 1 Ether in Wei
			Block: block,
		},
		{
			Hash:  "0x456",
			From:  address,
			To:    "0xdef",
			Value: "500000000000000000", // 0.5 Ether in Wei
			Block: block,
		},
	}
	return transactions, nil
}
