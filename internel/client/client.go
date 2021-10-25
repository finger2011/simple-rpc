package internelclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

//Client 入口
type Client struct {
	Endpoint string
}

//Input 输入
type Input struct {
	//数据
	In []interface{}
	//数据对应类型
	InTypes []reflect.Kind
}

//Output 输入
type Output struct {
	Out []interface{}
	Err error
}

//Call call remote function
func (c *Client) Call(serviceName, methodName string, input ...interface{}) (Output, error) {
	var output Output
	//TODO 不同的序列化/反序列化协议
	var in = make([]interface{}, len(input))
	var inTypes = make([]reflect.Kind, len(input))
	for i := 0; i < len(input); i++ {
		in[i] = input[i]
		inTypes[i] = reflect.TypeOf(input[i]).Kind()
	}
	var inputData = Input{In: in, InTypes: inTypes}
	fmt.Printf("input:%v\n", inputData)
	var inData, _ = json.Marshal(inputData)
	var client = &http.Client{}
	req, err := http.NewRequest(methodName, c.Endpoint, bytes.NewReader(inData))
	if err != nil {
		return output, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("simple-rpc-service", serviceName)
	req.Header.Set("simple-rpc-method", methodName)
	req.Header.Set("simple-rpc-num", strconv.Itoa(len(in)))

	resp, err := client.Do(req)
	if err != nil {
		return output, nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Output{}, err
	}
	_ = json.Unmarshal(data, &output)
	if output.Err != nil {
		return Output{}, output.Err
	}

	return output, nil
}
