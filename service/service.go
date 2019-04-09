package service

import (
	"MediView/data"
	"MediView/data/model"

	"github.com/google/uuid"
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

func (s *Service) OnNewRecord(pid uuid.UUID, v model.Vitals) error {
	return s.AddPatientRecord(pid, v)
}

//AddPatientRecord adds a new Record to the data store or returns an error
func (s *Service) AddPatientRecord(pid uuid.UUID, v model.Vitals) error {
	patient := s.data.GetPatient(pid)
	_, err := s.data.AddRecord(patient, v)
	if err != nil {
		return err
	}
	return nil
}

//DeletePatient removes a Patient from the data store
func (s *Service) DeletePatient(pid uuid.UUID) {
	s.data.DeletePatient(pid)
}

//DeleteRecord removes a Record from the data store
func (s *Service) DeleteRecord(rid uuid.UUID) {
	s.data.DeleteRecord(rid)
}
