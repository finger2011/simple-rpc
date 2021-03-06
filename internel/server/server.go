package internelserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

var errMethodNotExist = errors.New("method not exist")
var errParamNotMatch = errors.New("parameters not match")
var errParamLength = errors.New("paramter length can not convset to int")

//SimpleRPCServer server
type SimpleRPCServer struct {
}

//Server singlton server
var Server *SimpleRPCServer

//ServerOption server option
type ServerOption func(server *SimpleRPCServer) error

type response struct {
	Out []interface{}
	Err error
}

//Input 输入
type Input struct {
	//数据
	In []interface{}
	//数据对应类型
	InTypes []reflect.Kind
}

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
	num, err := strconv.Atoi(r.Header.Get("simple-rpc-num"))
	var respon response
	if err != nil {
		respon = response{
			Out: nil,
			Err: errParamLength,
		}
	} else {
		// input := make([]interface{}, num)
		input := Input{In: make([]interface{}, num), InTypes: make([]reflect.Kind, num)}
		
		_ = json.Unmarshal(data, &input)
		fmt.Printf("input:%v\n", input)
		service, _ := GetService(serviceName)
		res, err := callMethod(service, methodName, input)
		respon = response{
			Out: nil,
			Err: err,
		}
		if err == nil {
			respon.Out = make([]interface{}, len(res))
			for i := 0; i < len(res); i++ {
				respon.Out[i] = res[i].Interface()
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	output, _ := json.Marshal(respon)
	//TODO http code
	_, err = w.Write(output)
	if err != nil {
		fmt.Printf("write to response error:%v\n", err)
	}
}

func callMethod(s Service, methodName string, input Input) ([]reflect.Value, error) {
	val := reflect.ValueOf(s)
	method := val.MethodByName(methodName)
	var matchable, in = inParamterMatch(method, input)
	if !matchable {
		return nil, errParamNotMatch
	}

	return method.Call(in), nil
}

func inParamterMatch(fn reflect.Value, input Input) (bool, []reflect.Value) {
	if fn.Kind() != reflect.Func {
		return false, nil
	}
	if fn.Type().NumIn() != len(input.In) {
		return false, nil
	}
	var in = make([]reflect.Value, fn.Type().NumIn())
	for i := 0; i < len(input.In); i++ {
		if fn.Type().In(i).Kind() != input.InTypes[i] {
			//TODO int类型会被json Unmarshal为float64
			return false, nil
		}
		in[i] = reflect.ValueOf(input.In[i])
		if input.InTypes[i] != reflect.TypeOf(in[i]).Kind() {
			in[i] = in[i].Convert(fn.Type().In(i))
		}
	}
	return true, in
}
