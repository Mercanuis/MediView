package data

import (
	"MediView/data/model"
	"testing"

	"gotest.tools/assert"
)

//Comprehensive test
//TODO: break into units?
func TestPatientCache(t *testing.T) {
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

	//Add Vitals
	patient1, err := cache.AddRecord(key1, model.NewVitals(128, 78, 70, 45))
	if err != nil {
		t.Fatalf("Error with patient 1: %v", err)
	}
	patient2, err := cache.AddRecord(key2, model.NewVitals(120, 70, 65, 40))
	if err != nil {
		t.Fatalf("Error with patient 2: %v", err)
	}

	//Verify Patient Data
	assert.Equal(t, "Joe", patient1.Name)
	assert.Equal(t, 128, patient1.Vitals.Pressure.Systolic)
	assert.Equal(t, "Joan", patient2.Name)
	assert.Equal(t, 120, patient2.Vitals.Pressure.Systolic)

	//Delete Patient Data
	cache.DeletePatient(key1)
	cache.DeletePatient(key2)
}
