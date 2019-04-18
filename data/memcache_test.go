package data

import (
	"sort"
	"testing"

	"github.com/MediView/data/model"

	"gotest.tools/assert"
)

//Comprehensive test
func TestPatientCache_AddAndGet(t *testing.T) {
	//Initialize Patient Data
	cache := NewMemCache()
	err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}
	err = cache.AddPatient("Joan", 25)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	//Verify UUID keys
	patients := cache.GetPatients()
	key1 := patients.Records[0].ID
	key2 := patients.Records[1].ID
	assert.Assert(t, key1 != key2, "Keys are matching, should never happen")

	//Sort the keys for assertion
	var keys []int
	for k := range patients.Records {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	assert.Equal(t, "Joe", patients.Records[0].Name)
	assert.Equal(t, "Joan", patients.Records[1].Name)
}

func TestGetPatients(t *testing.T) {
	cache := NewMemCache()
	err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patients := cache.GetPatients()
	assert.Equal(t, 1, len(patients.Records))
}

func TestDeletePatient(t *testing.T) {
	cache := NewMemCache()
	err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patients := cache.GetPatients()
	cache.DeletePatient(patients.Records[0].ID)

	patients = cache.GetPatients()
	assert.Equal(t, 0, len(patients.Records))
}

func TestAddRecord(t *testing.T) {
	cache := NewMemCache()
	err := cache.AddPatient("Joey", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patients := cache.GetPatients()

	patient, err := cache.AddRecord(patients.Records[0].ID, model.Vitals{
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
	err := cache.AddPatient("Joey", 33)
	if err != nil {
		t.Fatalf("Failed to set up cache: %v", err)
	}

	patients := cache.GetPatients()
	_, err = cache.AddRecord(patients.Records[0].ID, model.Vitals{
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
	patients = cache.GetPatients()
	assert.Equal(t, 0, len(patients.Records))
}
