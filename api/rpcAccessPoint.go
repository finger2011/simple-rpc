package api

//RPCAccessPoint RPC框架对外提供的服务接口
type RPCAccessPoint interface {

	//客户端获取远程服务的引用
	GetRemoteService(uri, serviceName string) (*RPCService, error)

	//服务端注册服务的实现实例
	AddServiceProvider(service *RPCService, serviceName string) (string, error)

	//获取注册中心的引用
	GetNameService() (*NameService, error)

	//服务端启动RPC框架，监听接口，开始提供远程服务
	StartServer() error
}
