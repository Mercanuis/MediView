type bloodPressure struct {
  Systolic int
  Distolic int
}

//Vitals is a representation of a patients vital signs
type Vitals struct {
  bloodPressure
  Pulse   int    //BPM
  Glucose int    //mmol/L
}

//NewVitals returns a new Vitals struct
func NewVitals(sys, dys, pulse, glu int) Vitals {
  return Vitals{
    bloodPressure{
      Systolic: sys,
      Distolic: dys,
    },
    Pulse: pulse,
    Glucose: glu,
  }
}

//Patient represents a patient's information in the system
type Patient struct {
  Name  string
  Age   int
}

func NewPatient(i int64, n string, a int) Patient {
  return Patient {
    id: i,
    name: n,
    age: a,
  }
}
