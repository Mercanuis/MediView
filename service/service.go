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
func NewService(dao data.DAO) Service {
	return Service{
		data: dao,
	}
}

//GetLatestRecords returns the latest records
func (s *Service) GetLatestRecords() model.PatientRecords {
	//TODO: keep for debugging purposes, but afterward clean this up

	//key, _ := s.data.AddPatient("Joey", 33)
	//patient := s.data.GetPatient(key)
	//patient.Vitals = model.NewVitals(128, 78, 70, 45)
	//patRecs := s.data.GetPatients()
	return s.data.GetPatients()
}

func (s *Service) AddPatient(name string, age int) error {
	_, err := s.data.AddPatient(name, age)
	if err != nil {
		return err
	}
	return nil
}
