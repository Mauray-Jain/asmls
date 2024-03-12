package lsp

import "encoding/json"

type Request struct {
	Jsonrpc string           `json:"jsonrpc"`
	Id      *json.RawMessage `json:"id"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

func (req Request) IsNotif() bool {
	return req.Id == nil
}
