package model

import (
	"testing"

	"github.com/google/uuid"

	"gotest.tools/assert"
)

func TestVitals(t *testing.T) {
	v := NewVitals(122, 18, 78, 50)
	assert.Equal(t, 122, v.Pressure.Systolic)
	assert.Equal(t, 78, v.Pressure.Diastolic)
	assert.Equal(t, 78, v.Pulse)
	assert.Equal(t, 50, v.Glucose)
}

func TestPatient(t *testing.T) {
	v := NewPatient(uuid.New(), "Joe", 30)
	t.Logf("Created UUID: %v", v.Id)
	assert.Equal(t, "Joe", v.Name)
	assert.Equal(t, 30, v.Age)
}
