package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JSONRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type JSONRPCResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Result  BlockResponse `json:"result,omitempty"`
	Error   interface{}   `json:"error,omitempty"`
	ID      int           `json:"id"`
}

type BlockResponse struct {
	Number    string `json:"number"`
	TimeStamp string `json:"timestamp"`
}

func SendJSONRPCRequest(url string, method string, params []interface{}) (*JSONRPCResponse, error) {
	request := JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response JSONRPCResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
