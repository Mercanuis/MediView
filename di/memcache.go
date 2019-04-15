package di

import (
	"MediView/data"
)

//MemcacheProvider defines the interface of providing a MediService instance
type MemcacheProvider interface {
	//GetReciever provides a MediService instance
	GetMemcache() (data.DAO, error)
}

func (c *containerImpl) GetMemcache() (data.DAO, error) {
	if c.cache.memcache != nil {
		return *c.cache.memcache, nil
	}

	dao := data.NewMemCache()
	c.cache.memcache = &dao
	return dao, nil
}
