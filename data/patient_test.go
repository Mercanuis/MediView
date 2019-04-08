import (
  "testing",
  "assert"
)

func TestPatientMap(t *testing.T) {
  pat1 := NewPatient(123456, "John", 30)
  pat2 := NewPatient(234567, "Joe", 45)
  pat3 := NewPatient(345678, "Joan", 25)

  patMap := NewPatientMap()
  a := [3]Patient{pat1, pat2, pat3}
  for i, v := range a {
    err := patMap.AddPatient(v)
    if err != nil {
      t.Fatalf("Failed to add patient: %v", err)
    }
    p, err := patMap.GetPatient(v.id)
    if err != nil {
      t.Fatalf("Failed to get patient: %v", err)
    }
    assert.Equal(t, v, p)
  }

  err := patMap.AddPatient(pat3)
  if err == nil {
    t.Fatalf("Error Expected: Adding existing key to map")
  }

  patMap.DeletePatient(pat3)
  err := patMap.AddPatient(pat3)
  if err != nil {
    t.Fatalf("Failed to add patient: %v", err)
  }
}
