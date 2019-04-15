package di

import (
	"MediView/queue/receiver"
)

//ReceiverProvider defines the interface of providing a Receiver
type ReceiverProvider interface {
	//GetReciever provides a Receiver
	GetReceiver() (receiver.Receiver, error)
}

//GetReceiver returns a new Receiver
func (c *containerImpl) GetReceiver() (receiver.Receiver, error) {
	if c.cache.receiver != nil {
		return *c.cache.receiver, nil
	}

	service, err := c.GetMediService()
	if err != nil {
		return nil, err
	}

	rec := receiver.NewReceiver(service)
	c.cache.receiver = &rec
	return rec, nil
}
