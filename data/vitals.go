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
