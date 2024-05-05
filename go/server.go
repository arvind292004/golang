package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *int) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	*quo = args.A / args.B
	return nil
}

func main() {
	rpc.RegisterName("Arith", new(Arith))

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go rpc.ServeConn(conn)
	}
}
