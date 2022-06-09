package main

import (
	"fmt"
	"log"
	"net/rpc"

	"go-rpc/rpcproto"
)

func main() {
	sumRequest := &rpcproto.SumRequest{
		A: 1,
		B: 2,
	}
	sumReply := &rpcproto.SumReply{}
	client, err := rpc.DialHTTP("tcp", ":1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer client.Close()

	err = client.Call("SumService.Sum", sumRequest, sumReply)

	if err != nil {
		log.Fatal("call error:", err)
	}

	fmt.Println("Result:", sumReply.Result)
}
