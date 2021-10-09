package main

import (
	"finger2011/simpleRPC/internel/server"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type UserService struct {
	internelserver.Service
}
func (u *UserService) ServiceName() string {
	return "UserService"
}

func (u *UserService) SayHello(input *internelserver.Input) (string, error) {
	return "hello, user:" + input.Name, nil
}

func main() {
	fmt.Println("server start")
	internelserver.AddService(&UserService{})
	http.HandleFunc("/", internelserver.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
