package internelclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type client interface {
}

//Input 输入
type Input struct {
	Name string
}

//Output 输入
type Output struct {
	Data map[string]interface{}
}

//Call call remote function
func Call(serviceName, methodName, url string, input Input) (Output, error) {
	var output Output
	output.Data = make(map[string]interface{})
	var inData, _ = json.Marshal(input)
	var client = &http.Client{}
	req, err := http.NewRequest(methodName, url, bytes.NewReader(inData))
	if err != nil {
		return output, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("simple-rpc-service", serviceName)
	req.Header.Set("simple-rpc-method", methodName)

	resp, err := client.Do(req)
	if err != nil {
		return output, nil
	}
	data, _ := ioutil.ReadAll(resp.Body)
	output.Data["data"] = string(data)
	// _ = json.Unmarshal(data, output.Data["data"])
	return output, nil
}
