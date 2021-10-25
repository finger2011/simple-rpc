package main

import (
	internelserver "finger2011/simpleRPC/internel/server"
	"fmt"
	"log"
	"net/http"
)

//UserService 自定义service
type UserService struct {
	internelserver.Service
}

//ServiceName base service方法，获取service名称
func (u *UserService) ServiceName() string {
	return "UserService"
}

//SayHello service自身方法
func (u *UserService) SayHello(name string, age int) (string, error) {
	// fmt.Printf("call name:" + name + "age:" + fmt.Sprintf("%2f", age))
	return "hello, user:" + name + "; age:" + fmt.Sprintf("%d", age), nil
}

func main() {
	fmt.Println("server start")
	internelserver.AddService(&UserService{})
	http.HandleFunc("/", internelserver.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
