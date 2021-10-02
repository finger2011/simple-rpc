package api

//NameService 注册中心
type NameService interface {
	//注册服务
	RegisterService(serviceName, uri string) error
	//查找服务地址
	LookupService(serviceName string) (string, error)
}
