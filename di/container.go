package di

import (
	"MediView/http"
	"MediView/service"
)

type Container interface {
	HTTPServerProvider
	MediServiceProvider
}

type instanceCache struct {
	httpServer  *http.Server
	mediService *service.Service
}
type containerImpl struct {
	cache instanceCache
}

func NewContainer() Container {
	return &containerImpl{
		cache: instanceCache{},
	}
}
