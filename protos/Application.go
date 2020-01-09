package main

import (
	"fmt"
	"jaemon-go/protos/common"
)

func main() {
	var name string = "answer"
	var data []byte = []byte(name)

	var response = &common.Response{
		Code: 10000,
		Message: "success",
		Result: data,
	}

	fmt.Println(response.Message)

	fmt.Println(string(response.Result))

	fmt.Print("Start Appliaction...")
}
