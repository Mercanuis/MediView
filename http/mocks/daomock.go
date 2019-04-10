package mocks

import (
	"MediView/data/model"

	"github.com/google/uuid"
)

type DaoMock struct {
	GetPatientMock    func(id uuid.UUID) (*model.Patient, error)
	AddPatientMock    func(name string, age int) (uuid.UUID, error)
	GetPatientsMock   func() model.PatientRecords
	DeletePatientMock func(id uuid.UUID)
	AddRecordMock     func(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error)
}

func (d *DaoMock) GetPatients() model.PatientRecords {
	return d.GetPatientsMock()
}

func (d *DaoMock) GetPatient(id uuid.UUID) (*model.Patient, error) {
	return d.GetPatientMock(id)
}

func (d *DaoMock) AddPatient(name string, age int) (uuid.UUID, error) {
	return d.AddPatientMock(name, age)
}

func (d *DaoMock) DeletePatient(id uuid.UUID) {
	d.DeletePatientMock(id)
}

func (d *DaoMock) AddRecord(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
	return d.AddRecordMock(pid, vitals)
}
