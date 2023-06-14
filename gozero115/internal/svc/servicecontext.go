package svc

import (
	"demo/gozero115/internal/config"
	"fmt"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println("NewServiceContext")
	return &ServiceContext{
		Config: c,
	}
}
