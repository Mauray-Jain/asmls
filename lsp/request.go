package lsp

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Jsonrpc string           `json:"jsonrpc"`
	Id      *json.RawMessage `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

func (req Request) IsNotif() bool {
	return req.Id == nil
}

func (req Request) String() string {
	return fmt.Sprintf("Method: %s, Id: %v, Params: %v", req.Method, req.Id, string(*req.Params))
}
