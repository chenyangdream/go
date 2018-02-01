package main

import (
	"fmt"
	"net/rpc"

	"../common"
)

func main() {
	var args = common.Args{7, 8}
	var result = common.Result{}
	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接RPC服务器失败：", err)
	}

	divCall := client.Go("MathService.Divide", args, &result, nil)
	<-divCall.Done
}
