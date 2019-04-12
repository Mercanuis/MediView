package service

import (
	"MediView/data"
	"MediView/data/model"

	"github.com/pkg/errors"

	"github.com/google/uuid"
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
	return s.data.GetPatients()
}

//AddPatient adds a new patient to the system
func (s *Service) AddPatient(name string, age int) error {
	_, err := s.data.AddPatient(name, age)
	if err != nil {
		return errors.Wrap(err, "failed to add patient to system")
	}
	return nil
}

//AddRecord adds a new record for an existing Patient
func (s *Service) AddRecord(pid uuid.UUID, sys, dys, pul, glu int) (*model.Patient, error) {
	vitals := model.NewVitals(sys, dys, pul, glu)
	patient, err := s.data.AddRecord(pid, vitals)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add a record for patient")
	}
	return patient, nil
}

func (s *Service) GetHistories() model.PatientVitalHistories {
	return s.data.GetPatientHistories()
}

func (s *Service) ResetHistory() {
	s.data.ResetPatientHistory()
}

func (s *Service) DeleteHistory() {
	s.data.DeleteAllHistory()
}
