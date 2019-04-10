package mocks

import (
	"MediView/data/model"

	"github.com/google/uuid"
)

type DaoMock struct {
	AddPatientMock    func(name string, age int) (uuid.UUID, error)
	GetPatientMock    func(id uuid.UUID) model.Patient
	GetPatientsMock   func() model.PatientRecords
	DeletePatientMock func(id uuid.UUID)
}

func (d DaoMock) GetPatients() model.PatientRecords {
	return d.GetPatientsMock()
}

func (d DaoMock) GetPatient(id uuid.UUID) model.Patient {
	return d.GetPatientMock(id)
}

func (d DaoMock) AddPatient(name string, age int) (uuid.UUID, error) {
	return d.AddPatientMock(name, age)
}

func (d DaoMock) DeletePatient(id uuid.UUID) {
	d.DeletePatientMock(id)
}
