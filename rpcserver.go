package net

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/http"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

//func (t *T) MethodName(argType T1, replyType *T2) error
func (t *Arith) Multiply(args *Args, quo *Quotient) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return fmt.Errorf("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}
	go http.Serve(l,nil)
}
