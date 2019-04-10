package data

import (
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
