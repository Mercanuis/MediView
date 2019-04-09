package data

import (
	"MediView/data/model"
	"log"
	"sync"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

type (
	//PatientCache represents a list of Patients
	PatientCache struct {
		patients map[uuid.UUID]model.Patient
	}

	//RecordCache represents a list of Records
	RecordCache struct {
		records map[uuid.UUID]model.Record
	}

	//MemCache represents a data store of Patients and Records
	MemCache struct {
		sync.RWMutex
		PatientList PatientCache
		RecordsList RecordCache
	}
)

//NewMemCache initializes a new MemCache
func NewMemCache() DAO {
	return &MemCache{
		PatientList: PatientCache{
			patients: make(map[uuid.UUID]model.Patient),
		},
		RecordsList: RecordCache{
			records: make(map[uuid.UUID]model.Record),
		},
	}
}

func (m *MemCache) AddPatient(name string, age int) (uuid.UUID, error) {
	key, err := m.createKey()
	if err != nil {
		return uuid.UUID{}, err
	}

	if _, exists := m.PatientList.patients[key]; exists {
		//key existed already
		//TODO: is this possible with UUID? Will we ever reach this point?
		log.Print("Existing key")
		return key, nil
	}

	m.Lock()
	defer m.Unlock()
	m.PatientList.patients[key] = model.NewPatient(key, name, age)
	return key, nil
}

func (m *MemCache) AddRecord(p model.Patient, v model.Vitals) (uuid.UUID, error) {
	key, err := m.createKey()
	if err != nil {
		return uuid.UUID{}, err
	}

	m.Lock()
	defer m.Unlock()
	m.RecordsList.records[key] = model.NewRecord(p, v)
	return key, nil
}

func (m *MemCache) createKey() (key uuid.UUID, err error) {
	key = uuid.New()
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unable to create key: %v", r)
		}
	}()
	return key, err
}

func (m *MemCache) GetPatient(id uuid.UUID) model.Patient {
	return m.PatientList.patients[id]
}

func (m *MemCache) GetRecord(id uuid.UUID) model.Record {
	return m.RecordsList.records[id]
}

func (m *MemCache) DeletePatient(id uuid.UUID) {
	delete(m.PatientList.patients, id)
}

func (m *MemCache) DeleteRecord(id uuid.UUID) {
	delete(m.RecordsList.records, id)
}
