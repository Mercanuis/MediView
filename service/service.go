package service

import (
	"github.com/MediView/data"
	"github.com/MediView/data/model"

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
	err := s.data.AddPatient(name, age)
	if err != nil {
		return errors.Wrap(err, "[service] failed to add patient to system")
	}
	return nil
}

//AddRecord adds a new record for an existing Patient
func (s *Service) AddRecord(pid uuid.UUID, sys, dys, pul, glu int) (*model.Patient, error) {
	vitals := model.NewVitals(sys, dys, pul, glu)
	patient, err := s.data.AddRecord(pid, vitals)
	if err != nil {
		return nil, errors.Wrap(err, "[service] failed to add a record for patient")
	}
	return patient, nil
}

//GetHistories returns a list of Patient Histories
func (s *Service) GetHistories() model.PatientVitalHistories {
	return s.data.GetPatientHistories()
}

//ResetHistory resets the patient history
func (s *Service) ResetHistory() {
	s.data.ResetPatientHistory()
}

//DeleteHistory deletes the patient history
func (s *Service) DeleteHistory() {
	s.data.DeleteAllHistory()
}
