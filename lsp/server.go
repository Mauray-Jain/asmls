package lsp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/Mauray-Jain/asmls/methods"
)

type MethodHandler func(params *json.RawMessage) (result any, err error)
type NotifHandler func(params *json.RawMessage) (err error)

type Server struct {
	Reader    io.Reader
	Writer    io.Writer
	logger    log.Logger
	Documents map[string]string
	Methods   map[string]MethodHandler
	Notifs    map[string]NotifHandler
}

func NewServer(r io.Reader, w io.Writer) Server {
	return Server{
		Reader:    r,
		Writer:    w,
		Documents: map[string]string{},
		Methods: map[string]MethodHandler{
			"initialize": methods.Initialize,
		},
		Notifs: map[string]NotifHandler{},
	}
}

func (server *Server) write(obj Response) error {
	objJson, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(server.Writer, "Content-Length: %d\r\n\r\n", len(objJson))
	if err != nil {
		return err
	}
	_, err = server.Writer.Write(objJson)
	return err
}

// https://github.com/sourcegraph/jsonrpc2/blob/master/stream.go#L138
func (server *Server) read(req *Request) error {
	var contentLen int64
	reader := bufio.NewReader(server.Reader)

	for {
		line, err := reader.ReadString('\r')
		if err != nil {
			return err
		}
		newLineChar, err := reader.ReadByte()
		if err != nil {
			return err
		}
		if newLineChar != byte('\n') {
			return errors.New(`JsonRPC: \r should be followed by \n`)
		}
		if line == "\r" {
			break
		}
		if strings.HasPrefix(line, "Content-Length: ") {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "Content-Length: ")
			var err error
			contentLen, err = strconv.ParseInt(line, 10, 64)
			if err != nil {
				return err
			}
			// Don't break here as there maybe other headers as well and \r\n just before content isn't handled
		}
	}

	if contentLen == 0 {
		return errors.New("JsonRPC: No `Content-Length: ` header found")
	}
	return json.NewDecoder(io.LimitReader(reader, contentLen)).Decode(req)
}

func (server *Server) HandleMethod(req Request) {
	handler, ok := server.Methods[req.Method]

	if !ok {
		err := server.write(NewResErr(req.Id, ErrMethodNotFound))
		server.logger.Printf("Method %s not found", req.Method)
		if err != nil {
			server.logger.Println("Error while responding: ", err.Error())
		}
		return
	}

	var res Response
	result, err := handler(req.Params)
	if err != nil {
		res = NewResErr(req.Id, err)
	} else {
		res = NewResponse(req.Id, result)
	}

	err = server.write(res)
	if err != nil {
		server.logger.Println("Error while responding: ", err.Error())
	}
}
