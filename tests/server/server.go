package main

import (
	rpc "finger2011/simpleRPC/internel/server"
	"fmt"
	"log"
	"net/http"
)

//UserService 自定义service
type UserService struct {
}

//ServiceName base service方法，获取service名称
func (u *UserService) ServiceName() string {
	return "UserService"
}

//SayHello service自身方法
func (u *UserService) SayHello(name string, age int) (string, error) {
	return "hello, user:" + name + "; age:" + fmt.Sprintf("%d", age), nil
}

func main() {
	fmt.Println("server start")
	rpc.AddService(&UserService{})
	http.HandleFunc("/", rpc.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
