package data

import (
	"MediView/data/model"
	"sync"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

type (
	//MemCache represents a data store of Patients and Records
	MemCache struct {
		sync.RWMutex
		PatientList    map[uuid.UUID]model.Patient
		PatientHistory map[uuid.UUID]model.PatientVitalHistory
	}
)

//NewMemCache initializes a new MemCache
func NewMemCache() DAO {
	return &MemCache{
		PatientList:    make(map[uuid.UUID]model.Patient),
		PatientHistory: make(map[uuid.UUID]model.PatientVitalHistory),
	}
}

func (m *MemCache) GetPatient(id uuid.UUID) (*model.Patient, error) {
	pat, ok := m.PatientList[id]
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

	//Write to table
	m.Lock()
	defer m.Unlock()
	m.PatientList[key] = model.NewPatient(key, name, age)
	return key, nil
}

func (m *MemCache) createKey() (uuid.UUID, error) {
	//err being the 'zero-value' will be nil
	var err error
	key := uuid.New()

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unable to create key: %v", r)
		}
	}()
	return key, err
}

func (m *MemCache) GetPatients() model.PatientRecords {
	var arr []model.Patient
	for _, v := range m.PatientList {
		arr = append(arr, v)
	}

	return model.PatientRecords{
		Records: arr,
	}
}

func (m *MemCache) DeletePatient(id uuid.UUID) {
	delete(m.PatientList, id)
}

func (m *MemCache) AddRecord(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
	pat, ok := m.PatientList[pid]
	if !ok {
		return nil, errors.New("invalid key")
	}

	//Create a new history for a new patient
	//If they exist, update their history
	history := m.getPatientHistory(pid, vitals)
	pat.Vitals = vitals

	//Write to both tables.
	m.Lock()
	defer m.Unlock()
	m.PatientList[pid] = pat
	m.PatientHistory[pid] = history
	return &pat, nil
}

func (m *MemCache) getPatientHistory(pid uuid.UUID, vitals model.Vitals) model.PatientVitalHistory {
	var history model.PatientVitalHistory
	if _, ok := m.PatientHistory[pid]; !ok {
		history = model.NewPatientVitalHistory(pid, vitals.Pressure, vitals.Pulse, vitals.Glucose)
	} else {
		history = m.PatientHistory[pid]
		history.UpdateHistory(vitals.Pressure.Systolic, vitals.Pressure.Diastolic, vitals.Pulse, vitals.Glucose)
	}
	return history
}

func (m *MemCache) GetPatientHistories() model.PatientVitalHistories {
	var arr []model.PatientVitalHistory
	for _, v := range m.PatientHistory {
		arr = append(arr, v)
	}

	return model.PatientVitalHistories{
		Histories: arr,
	}
}

func (m *MemCache) ResetPatientHistory() {
	for k := range m.PatientHistory {
		m.PatientHistory[k] = model.PatientVitalHistory{
			ID: k,
		}
	}
}

func (m *MemCache) DeleteAllHistory() {
	for k := range m.PatientHistory {
		delete(m.PatientHistory, k)
	}
}
