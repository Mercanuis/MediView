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

	//MemCache represents a data store of Patients and Records
	MemCache struct {
		sync.RWMutex
		PatientList PatientCache
	}
)

//NewMemCache initializes a new MemCache
func NewMemCache() DAO {
	return &MemCache{
		PatientList: PatientCache{
			patients: make(map[uuid.UUID]model.Patient),
		},
	}
}

func (m *MemCache) GetPatient(id uuid.UUID) (*model.Patient, error) {
	pat, ok := m.PatientList.patients[id]
	if !ok {
		return nil, errors.New("invalid key")
	}
	return &pat, nil
}

func (m *MemCache) AddPatient(name string, age int) (uuid.UUID, error) {
	key, err := m.createKey()
	if err != nil {
		return uuid.UUID{}, err
	}

	//TODO: is this possible with UUID? Will we ever reach this point?
	if _, exists := m.PatientList.patients[key]; exists {
		//key existed already
		log.Print("Existing key")
		return key, nil
	}

	m.Lock()
	defer m.Unlock()
	m.PatientList.patients[key] = model.NewPatient(key, name, age)
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

func (m *MemCache) GetPatients() model.PatientRecords {
	var arr []model.Patient
	for _, v := range m.PatientList.patients {
		arr = append(arr, v)
	}

	return model.PatientRecords{
		Records: arr,
	}
}

func (m *MemCache) DeletePatient(id uuid.UUID) {
	delete(m.PatientList.patients, id)
}

func (m *MemCache) AddRecord(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
	pat, ok := m.PatientList.patients[pid]
	if !ok {
		return nil, errors.New("invalid key")
	}

	pat.Vitals = vitals

	m.Lock()
	defer m.Unlock()
	m.PatientList.patients[pid] = pat
	return &pat, nil
}
