package service

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetRecords(t *testing.T) {
	//Initialize service
	s := NewService()
	key1, err := s.data.AddPatient("Joe", 30)
	if err != nil {
		t.Fatalf("Failed to create patient 1: %v", err)
	}
	key2, err := s.data.AddPatient("Joan", 25)
	if err != nil {
		t.Fatalf("Failed to create patient 2: %v", err)

	}
	key3, err := s.data.AddPatient("John", 40)
	if err != nil {
		t.Fatalf("Failed to create patient 3: %v", err)
	}

	//Because array is un-ordered we will just loop through and check all entries.
	patients := s.GetLatestRecords()
	found1 := false
	found2 := false
	found3 := false

	for _, i := range patients {
		if i.Id == key1 {
			if i.Name == "Joe" && i.Age == 30 {
				found1 = true
			}
		} else if i.Id == key2 {
			if i.Name == "Joan" && i.Age == 25 {
				found2 = true
			}
		} else if i.Id == key3 {
			if i.Name == "John" && i.Age == 40 {
				found3 = true
			}
		}
	}

	assert.Assert(t, found1 && found2 && found3, "Failed to find all patients")
}
