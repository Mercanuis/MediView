package di

import (
	"MediView/data"
)

//MemcacheProvider defines the interface of providing a memory cache
type MemcacheProvider interface {
	//GetReciever provides a Memcache
	GetMemcache() (data.DAO, error)
}

//GetMemcache returns a new DAO instance
func (c *containerImpl) GetMemcache() (data.DAO, error) {
	if c.cache.memcache != nil {
		return *c.cache.memcache, nil
	}

	dao := data.NewMemCache()
	c.cache.memcache = &dao
	return dao, nil
}
