package data

import (
	"testing"

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
		Id:   1234567,
		Name: "Joe",
		Age:  30,
	}

	assert.Equal(t, int64(1234567), v.Id)
	assert.Equal(t, "Joe", v.Name)
	assert.Equal(t, 30, v.Age)
}
