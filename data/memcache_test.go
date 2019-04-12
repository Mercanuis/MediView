package data

import (
	"MediView/data/model"
	"testing"

	"gotest.tools/assert"
)

//Comprehensive test
func TestPatientCache_AddAndGet(t *testing.T) {
	//Initialize Patient Data
	cache := NewMemCache()
	key1, err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}
	key2, err := cache.AddPatient("Joan", 25)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	//Verify UUID keys
	assert.Assert(t, key1 != key2, "Keys are matching, should never happen")

	pat1, err := cache.GetPatient(key1)
	if err != nil {
		t.Fatalf("Failed to get patient: %v", err)
	}
	pat2, err := cache.GetPatient(key2)
	if err != nil {
		t.Fatalf("Failed to get patient: %v", err)
	}
	assert.Equal(t, "Joe", pat1.Name)
	assert.Equal(t, "Joan", pat2.Name)
}

func TestGetPatients(t *testing.T) {
	cache := NewMemCache()
	_, err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patients := cache.GetPatients()
	assert.Equal(t, 1, len(patients.Records))
}

func TestDeletePatient(t *testing.T) {
	cache := NewMemCache()
	key, err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	cache.DeletePatient(key)

	patients := cache.GetPatients()
	assert.Equal(t, 0, len(patients.Records))
}

func TestAddRecord(t *testing.T) {
	cache := NewMemCache()
	key, err := cache.AddPatient("Joey", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patient, err := cache.AddRecord(key, model.Vitals{
		Pressure: model.BloodPressure{
			Systolic:  128,
			Diastolic: 78,
		},
		Pulse:   70,
		Glucose: 45,
	})
	if err != nil {
		t.Fatalf("Failed to add recored: %v", err)
	}

	assert.Equal(t, 128, patient.Vitals.Pressure.Systolic)
}

func TestPatientHistory(t *testing.T) {
	cache := NewMemCache()
	key, err := cache.AddPatient("Joey", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	_, err = cache.AddRecord(key, model.Vitals{
		Pressure: model.BloodPressure{
			Systolic:  128,
			Diastolic: 78,
		},
		Pulse:   70,
		Glucose: 45,
	})
	if err != nil {
		t.Fatalf("Failed to add record: %v", err)
	}

	cache.ResetPatientHistory()
	history := cache.GetPatientHistories()
	assert.Equal(t, 1, len(history.Histories))
	assert.Equal(t, 0, history.Histories[0].GAvg)
	cache.DeleteAllHistory()

	history = cache.GetPatientHistories()
	assert.Equal(t, 0, len(history.Histories))
}
