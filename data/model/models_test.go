package model

import (
	"testing"

	"github.com/google/uuid"

	"gotest.tools/assert"
)

func TestVitals(t *testing.T) {
	v := Vitals{
		Pressure: bloodPressure{
			Systolic:  122,
			Diastolic: 78,
		},
		Pulse:   78,
		Glucose: 50,
	}

	assert.Equal(t, 122, v.Pressure.Systolic)
	assert.Equal(t, 78, v.Pressure.Diastolic)
	assert.Equal(t, 78, v.Pulse)
	assert.Equal(t, 50, v.Glucose)
}

func TestPatient(t *testing.T) {
	v := Patient{
		Id:   uuid.New(),
		Name: "Joe",
		Age:  30,
	}

	t.Logf("Created UUID: %v", v.Id)
	assert.Equal(t, "Joe", v.Name)
	assert.Equal(t, 30, v.Age)
}

func TestRecord(t *testing.T) {
	r := Record{
		Patient: Patient{
			Id:   uuid.New(),
			Age:  34,
			Name: "Joan",
		},
		Vitals: Vitals{
			Pressure: bloodPressure{
				Systolic:  123,
				Diastolic: 70,
			},
			Pulse:   80,
			Glucose: 60,
		},
	}

	t.Logf("Created UUID: %v", r.Patient.Id)
	assert.Equal(t, "Joan", r.Patient.Name)
	assert.Equal(t, 34, r.Patient.Age)
	assert.Equal(t, 123, r.Vitals.Pressure.Systolic)
	assert.Equal(t, 60, r.Vitals.Glucose)
}
