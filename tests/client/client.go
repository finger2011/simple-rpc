package main

import (
	internelclient "finger2011/simpleRPC/internel/client"
	"fmt"
)

func main() {
	fmt.Println("client start")
	var url = "http://localhost:8080/golang"
	var input = internelclient.Input{Name: "golang"}
	output, err := internelclient.Call("UserService", "SayHello", url, input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("output:%v\n", output)
}
