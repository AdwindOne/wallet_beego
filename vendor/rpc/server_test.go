package rpc

import (
	"errors"
	"net"
	"net/http"
	"testing"
)

type Service1Request struct {
	A int
	B int
}

type Service1Response struct {
	Result int
}

type Service1 struct {
}

func (t *Service1) Multiply(r *http.Request, req *Service1Request, res *Service1Response) error {
	res.Result = req.A * req.B
	return nil
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B * 100
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestServer(t *testing.T) {
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", ":8000"); err != nil {
		t.Errorf("TestServer error %+v", err)
	}
	server := NewServer()
	server.Register(new(Arith))
	NewHTTPServer(server, []string{"*"}).Serve(listener)
	// curl -X POST  -d '{"id": 1, "method": "Arith.Multiply", "params":[{"A":1, "B":3}]}' http://localhost:8000/
}
