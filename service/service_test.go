package service

import (
	"sort"
	"testing"

	"github.com/MediView/data"

	"gotest.tools/assert"
)

func TestGetRecords(t *testing.T) {
	//Initialize service
	s := NewService(data.NewMemCache())
	err := s.AddPatient("Joe", 30)
	if err != nil {
		t.Fatalf("Failed to create patient 1: %v", err)
	}
	err = s.AddPatient("Joan", 25)
	if err != nil {
		t.Fatalf("Failed to create patient 2: %v", err)

	}
	err = s.AddPatient("John", 40)
	if err != nil {
		t.Fatalf("Failed to create patient 3: %v", err)
	}

	//Because array is un-ordered we will just loop through and check all entries.
	patients := s.GetLatestRecords()
	//Sort the keys for assertion
	var keys []int
	for k := range patients.Records {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	assert.Equal(t, "Joe", patients.Records[0].Name, "failed to find Joe")
	assert.Equal(t, "Joan", patients.Records[1].Name, "failed to find Joan")
	assert.Equal(t, "John", patients.Records[2].Name, "failed to find John")
}

func TestAddRecord(t *testing.T) {
	//Initialize service
	s := NewService(data.NewMemCache())
	err := s.data.AddPatient("Joe", 30)
	if err != nil {
		t.Fatalf("Failed to create patient: %v", err)
	}

	patients := s.GetLatestRecords()
	pat := patients.Records[0]
	p, err := s.AddRecord(pat.ID, 128, 78, 70, 45)
	if err != nil {
		t.Fatalf("failed to add record: %v", err)
	}

	assert.Equal(t, 128, p.Vitals.Pressure.Systolic)
	assert.Equal(t, 78, p.Vitals.Pressure.Diastolic)
	assert.Equal(t, 70, p.Vitals.Pulse)
	assert.Equal(t, 45, p.Vitals.Glucose)
}
