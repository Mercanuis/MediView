package model

import (
	"testing"

	"github.com/google/uuid"

	"gotest.tools/assert"
)

func TestVitals(t *testing.T) {
	v := NewVitals(122, 18, 78, 50)
	assert.Equal(t, 122, v.Pressure.Systolic)
	assert.Equal(t, 18, v.Pressure.Diastolic)
	assert.Equal(t, 78, v.Pulse)
	assert.Equal(t, 50, v.Glucose)
}

func TestPatient(t *testing.T) {
	v := NewPatient(uuid.New(), "Joe", 30)
	t.Logf("Created UUID: %v", v.ID)
	assert.Equal(t, "Joe", v.Name)
	assert.Equal(t, 30, v.Age)
}

func TestPatientVitalHistory(t *testing.T) {
	bpa := BloodPressure{
		Systolic:  128,
		Diastolic: 80,
	}
	v := NewPatientVitalHistory(uuid.New(), bpa, 75, 45)
	t.Logf("Created UUID: %v", v.ID)
	assert.Equal(t, 128, v.BPA.Systolic)
	assert.Equal(t, 75, v.PAvg)
	assert.Equal(t, 45, v.GAvg)
}

func TestUpdateHistory(t *testing.T) {
	bpa := BloodPressure{
		Systolic:  128,
		Diastolic: 80,
	}
	v := NewPatientVitalHistory(uuid.New(), bpa, 75, 45)
	v.UpdateHistory(145, 78, 65, 70)
	assert.Equal(t, 136, v.BPA.Systolic)
	assert.Equal(t, 79, v.BPA.Diastolic)
	assert.Equal(t, 70, v.PAvg)
	assert.Equal(t, 57, v.GAvg)
}
