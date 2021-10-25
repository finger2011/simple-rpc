package internelserver

import (
	"errors"
	"sync"
)

//ErrServiceNotFound service not found error
var ErrServiceNotFound = errors.New("service not found")

//Service base service
type Service interface {
	ServiceName() string
}

//存储所有注册的service
var services sync.Map

//AddService 注册service
func AddService(service Service) {
	services.Store(service.ServiceName(), service)
}

//GetService get service by service name
func GetService(serviceName string) (Service, error) {
	if service, ok := services.Load(serviceName); ok {
		return service.(Service), nil
	}
	return nil, ErrServiceNotFound
}
