package internelserver

var _ Service = &HelloService{}

type HelloService struct {
	Service
	// SayHello func(name string) (string, error)
}

func (h *HelloService) ServiceName() string {
	return "HelloService"
}

func (h *HelloService) SayHello(name string) (string, error) {
	return "hello, " + name, nil
}
