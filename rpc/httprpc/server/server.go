package main

import (
	"fmt"
	"net/http"
	"net/rpc"

	"../common"
)

func main() {
	var ms = new(common.MathService)
	rpc.Register(ms)
	rpc.HandleHTTP()
	fmt.Println("启动服务...")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("服务已停止!")
}
