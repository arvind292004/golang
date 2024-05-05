package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer client.Close()

	args := &Args{7, 8}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("Error calling Multiply:", err)
		return
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quo int
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		fmt.Println("Error calling Divide:", err)
		return
	}

	fmt.Printf("Arith: %d / %d = %d\n", args.A, args.B, quo)
}
