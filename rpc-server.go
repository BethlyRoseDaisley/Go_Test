package main

import (
    //"fmt"
    "log"
    "net"
    "net/http"
    "net/rpc"
    //"os"
	//"time"
	"errors"
)


type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int
//乘
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
//除
func (t *Arith) Divide(args *Args, quo *Quotient) error {

	if args.B == 0 {
		return errors.New("divide by zero")
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
		log.Fatal("listen error:", e)
	}
	
	http.Serve(l, nil)

}