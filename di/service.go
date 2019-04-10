package di

import (
	"MediView/data"
	"MediView/service"
)

//MediServiceProvider defines the interface of providing a MediService instance
type MediServiceProvider interface {
	//GetMediService provides a MediService instance
	GetMediService() (service.Service, error)
}

//GetMediService provides a MediService instance
func (c *containerImpl) GetMediService() (service.Service, error) {
	if c.cache.mediService != nil {
		return *c.cache.mediService, nil
	}

	dao := data.NewMemCache()
	s := service.NewService(dao)

	c.cache.mediService = &s
	return s, nil
}
