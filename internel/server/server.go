package internelserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

//SimpleRPCServer server
type SimpleRPCServer struct {
}

//Server singlton server
var Server *SimpleRPCServer

//ServerOption server option
type ServerOption func(server *SimpleRPCServer) error

//InitRPCServer init rpc server
func InitRPCServer(opts ...ServerOption) error {
	Server = &SimpleRPCServer{
		//默认实现
	}
	for _, opt := range opts {
		err := opt(Server)
		if err != nil {
			return err
		}
	}
	return nil
}

//Handler http handler
//TODO error 
func Handler(w http.ResponseWriter, r *http.Request) {
	//get service, method and data
	data, _ := ioutil.ReadAll(r.Body)
	//service name and method name sent in the http header
	serviceName := r.Header.Get("simple-rpc-service")
	methodName := r.Header.Get("simple-rpc-method")
	service, _ := GetService(serviceName)

	val := reflect.ValueOf(service)
	method := val.MethodByName(methodName)

	//filter parameters
	inType := method.Type().In(0)
	in := reflect.New(inType.Elem())
	_ = json.Unmarshal(data, in.Interface())

	//call method and get response
	res := method.Call([]reflect.Value{in})
	output, _ := json.Marshal(res[0].Interface())
	//TODO http code
	fmt.Fprintf(w, "%s", string(output))
}
