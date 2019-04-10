package service

import (
	"MediView/data"
	"MediView/data/model"
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
func (s *Service) GetLatestRecords() []model.Patient {
	patRecs := s.data.GetPatients()
	return patRecs
}
