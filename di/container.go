package di

import (
	"MediView/data"
	"MediView/http"
	"MediView/queue/receiver"
	"MediView/service"
)

type Container interface {
	HTTPServerProvider
	MediServiceProvider
	ReceiverProvider
	MemcacheProvider
}

type instanceCache struct {
	httpServer  *http.Server
	mediService *service.Service
	receiver    *receiver.Receiver
	memcache    *data.DAO
}
type containerImpl struct {
	cache instanceCache
}

func NewContainer() Container {
	return &containerImpl{
		cache: instanceCache{},
	}
}
