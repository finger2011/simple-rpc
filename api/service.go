package api

//RPCService 对外提供服务
type RPCService interface {
	Closable
}

//Closable 该关闭
type Closable interface {
	Close() error
}
