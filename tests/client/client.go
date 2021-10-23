package main

import (
	internelclient "finger2011/simpleRPC/internel/client"
	"fmt"
)

func main() {
	fmt.Println("client start")
	var client = internelclient.Client{Endpoint: "http://localhost:8080/"}
	output, err := client.Call("UserService", "SayHello", "golang", 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("output:%v\n", output)
}
