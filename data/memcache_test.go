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

	//Verify Patient Data
	patient1 := cache.GetPatient(key1)
	assert.Equal(t, "Joe", patient1.Name)
	patient2 := cache.GetPatient(key2)
	assert.Equal(t, "Joan", patient2.Name)

	//Delete Patient Data
	cache.DeletePatient(key1)
	cache.DeletePatient(key2)
}

//Comprehensive test
//TODO: break into units?
func TestRecordCache(t *testing.T) {
	//Initialize Record Data
	cache := NewMemCache()
	pkey, err := cache.AddPatient("Joe", 33)
	if err != nil {
		t.Fatalf("Failed to set up patient: %v", err)
	}
	patient := cache.GetPatient(pkey)
	vitals1 := model.NewVitals(122, 88, 76, 45)
	vitals2 := model.NewVitals(135, 90, 80, 42)

	//Verify UUID Keys
	rkey1, err := cache.AddRecord(patient, vitals1)
	if err != nil {
		t.Fatalf("Failed to add record: %v", err)
	}
	rkey2, err := cache.AddRecord(patient, vitals2)
	if err != nil {
		t.Fatalf("Failed to add record: %v", err)
	}
	assert.Assert(t, rkey1 != rkey2, "Keys are matching, should never happen")

	//Verify Record Data
	record1 := cache.GetRecord(rkey1)
	assert.Equal(t, vitals1, record1.Vitals)
	record2 := cache.GetRecord(rkey2)
	assert.Equal(t, vitals2, record2.Vitals)

	//Delete Record Data
	cache.DeleteRecord(rkey1)
	cache.DeleteRecord(rkey2)
}
