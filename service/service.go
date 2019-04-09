package service

import (
	"MediView/data"
)

//Service represents the service logic
type Service struct {
	data data.DAO
}

//NewService initializes a new Service
func NewService() Service {
	return Service{
		data: data.NewMemCache(),
	}
}

//GetLatestRecords returns the latest records
func (s *Service) GetLatestRecords() {
	_ = s.data.GetRecords()
	//TODO: Work here first and get the get working
}
