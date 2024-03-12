package lsp

import "encoding/json"

type Response struct {
	Jsonrpc string           `json:"jsonrpc"`
	Id      *json.RawMessage `json:"id,omitempty"`
	Result  any              `json:"result,omitempty"`
	Error   *ResErr          `json:"error,omitempty"`
}

type ResErr struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (err *ResErr) Error() string {
	return err.Message
}

var (
	ErrParseError           = &ResErr{Code: -32700, Message: "Parse Error"}
	ErrInvalidRequest       = &ResErr{Code: -32600, Message: "Invalid Request"}
	ErrMethodNotFound       = &ResErr{Code: -32601, Message: "Method Not Found"}
	ErrInvalidParams        = &ResErr{Code: -32602, Message: "Invalid Params"}
	ErrInternalError        = &ResErr{Code: -32603, Message: "Internal Error"}
	ErrServerNotInitialized = &ResErr{Code: -32002, Message: "Server Not Initialized"}
	ErrUnknownErrorCode     = &ResErr{Code: -32001, Message: "Unknown Error Code"}
	ErrRequestFailed        = &ResErr{Code: -32803, Message: "Request Failed"}
	ErrServerCancelled      = &ResErr{Code: -32802, Message: "Server Cancelled"}
	ErrRequestCancelled     = &ResErr{Code: -32800, Message: "Request Cancelled"}
)

func NewResponse(id *json.RawMessage, result any) Response {
	return Response{
		Jsonrpc: "2.0",
		Id:      id,
		Result:  result,
		Error:   nil,
	}
}

func NewResErr(id *json.RawMessage, err error) Response {
	errObj, ok := err.(*ResErr)
	if !ok {
		errObj = &ResErr{
			Code:    0,
			Message: err.Error(),
		}
	}
	return Response{
		Jsonrpc: "2.0",
		Id:      id,
		Result:  nil,
		Error:   errObj,
	}
}
