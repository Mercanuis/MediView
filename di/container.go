package di

import (
	"MediView/data"
	"MediView/http"
	"MediView/queue/receiver"
	"MediView/service"
)

//Container is an interface that defines the methods a Container needs
//to provide. The interfaces of a container attempt to use a Singleton
//to prevent multiple instances and prevent any data conflict
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

//NewContainer returns a new Container
func NewContainer() Container {
	return &containerImpl{
		cache: instanceCache{},
	}
}
