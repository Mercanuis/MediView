import (
  "error"
  "sync"
)

//Patient represents a patient's information in the system
type Patient struct {
  Name  string
  Age   int
}

func NewPatient(i int64, n string, a int) Patient {
  return Patient {
    id: i,
    name: n,
    age: a,
  }
}

type PatientMap struct {
  sync.RWMutex
  patients map[int64]Patient
}

func NewPatientMap() PatientMap {
  return PatientMap {
    patients: make(map[int64]Patient)
  }
}

func(m *PatientMap) AddPatient(p Patient) error {
  if m.patients[p.Id] != nil {
    return errors.New("Patient already in system")
  }

  m.Lock()
  m.patients[p.Id] = p
  m.Unlock()
  return nil
}

func(m *PatientMap) GetPatient(pid int64) (*Patient, error) {
  m.RLock()
  p := m.patients[pid]
  m.RUnlock()

  if p == nil {
    return nil, errors.New("Patient does not exist")
  }
  return p, nil
}

func(m *PatientMap) DeletePatient(p Patient) {
  if m.patients[p.Id] == nil {
    return
  }
  m.Lock()
  delete(m.patients, p.Id)
  m.Unlock()
}
