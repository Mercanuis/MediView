package data

type bloodPressure struct {
	Systolic  int
	Diastolic int
}

//Vitals is a representation of a patients vital signs
type Vitals struct {
	Pressure bloodPressure
	Pulse    int //BPM
	Glucose  int //mmol/L
}

//NewVitals returns a new Vitals struct
func NewVitals(sys, dys, pulse, glu int) Vitals {
	return Vitals{
		Pressure: bloodPressure{
			Systolic:  sys,
			Diastolic: dys,
		},
		Pulse:   pulse,
		Glucose: glu,
	}
}

//Patient represents a patient's information in the system
type Patient struct {
	Id   int64
	Name string
	Age  int
}

//NewPatient returns a new Patient struct
func NewPatient(n string, i int64, a int) Patient {
	return Patient{
		Id:   i,
		Name: n,
		Age:  a,
	}
}
