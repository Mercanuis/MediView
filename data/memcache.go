package data

import "MediView/data/model"

type (
	PatientCache struct {
		patients map[int64]model.Patient
	}

	RecordCache struct {
		records map[model.Patient]model.Record
	}

	MemCache struct {
		Patients PatientCache
		Records  RecordCache
	}
)

func NewMemCache() MemCache {
	return MemCache{
		Patients: PatientCache{
			patients: make(map[int64]model.Patient),
		},
		Records: RecordCache{
			records: make(map[model.Patient]model.Record),
		},
	}
}

func (m *MemCache) Add(id int64, data interface{}) error {
	//TODO: implement
	return nil
}

func (m *MemCache) Delete(id int64) error {
	//TODO: implement
	return nil
}
