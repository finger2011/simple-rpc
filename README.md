# simple-rpc
simple rpc 

## v0.1 
+ 在client和server之间以http方式进行调用
+ client端
    + 设置url，调用rpc方法，传入service name, method name和参数信息
    + rpc接收service name, method name 和参数信息，并进行封装
    + 开启一个http请求，将数据发送给server端 
    + 接收server端返回，并解析数据，返回数据给方法调用者
+ server端
    + 接收http请求，解析数据
    + 参数校验：包括service name,method name以及参数类型校验
    + 调用service对应方法，返回结果
### 例子
server 端
```golang
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
```

client端
```golang
package main

import (
	rpc "finger2011/simpleRPC/internel/client"
	"fmt"
)

func main() {
	fmt.Println("client start")
    //设置URL
	var client = rpc.Client{Endpoint: "http://localhost:8080/"}
    // 调用rpc方法，传入service name, method name和参数信息
	output, err := client.Call("UserService", "SayHello", "golang", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("output:%v\n", output)
}
```

## TODO
+ 引入以pb方式生成client stub和service代码
+ 增加以gprc方式在client和server端之间的调用
+ 讲gprc和http做成可配置参数
