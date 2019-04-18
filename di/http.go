package di

import (
	"github.com/MediView/http"
)

// HTTPServerProvider defines the interface of providing a HTTP server
type HTTPServerProvider interface {
	// GetHTTPServer provides an HTTP server
	GetHTTPServer() (*http.Server, error)
}

// GetHTTPServer provides an HTTP server
func (c *containerImpl) GetHTTPServer() (*http.Server, error) {
	if c.cache.httpServer != nil {
		return c.cache.httpServer, nil
	}

	service, err := c.GetMediService()
	if err != nil {
		return nil, err
	}

	receiver, err := c.GetReceiver()
	if err != nil {
		return nil, err
	}

	server, err := http.New(service, receiver)
	if err != nil {
		return nil, err
	}

	c.cache.httpServer = server
	return server, nil
}
