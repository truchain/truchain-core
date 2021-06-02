package rpc

import (
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func RunServer() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(RPCService), "")

	http.Handle("/rpc", s)
}
