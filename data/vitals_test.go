import (
  "testing",
  "assert"
)

func TestVitals(t *testing.T) {
  v := Vitals{
    bloodPressure {
      Systolic: 122,
      Distolic: 78,
    },
    Pulse: 78,
    Glucose: 50,
  }

  assert.Equal(t, 122, v.bloodPressure.Systolic)
  assert.Equal(t, 78, v.bloodPressure.Distolic)
  assert.Equal(t, 78, v.Pulse)
  assert.Equal(t, 50, v.Glucose)
}
