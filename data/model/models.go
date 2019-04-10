package model

import "github.com/google/uuid"

type bloodPressure struct {
	Systolic  int `json:"systolic"`
	Diastolic int `json:"diastolic"`
}

//Vitals is a representation of a patients current vital signs
type Vitals struct {
	Pressure bloodPressure `json:"pressure"`
	Pulse    int           `json:"pulse"`
	Glucose  int           `json:"glucose"`
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
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Vitals Vitals    `json:"vitals"`
}

//NewPatient returns a new Patient struct
func NewPatient(i uuid.UUID, n string, a int) Patient {
	return Patient{
		Id:   i,
		Name: n,
		Age:  a,
	}
}

//PatientRecords represents a series of Patients
type PatientRecords struct {
	Records []Patient `json:"records"`
}
