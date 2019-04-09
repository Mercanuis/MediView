package model

import "github.com/google/uuid"

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
	Id   uuid.UUID
	Name string
	Age  int
}

//NewPatient returns a new Patient struct
func NewPatient(i uuid.UUID, n string, a int) Patient {
	return Patient{
		Id:   i,
		Name: n,
		Age:  a,
	}
}

//Record represents a medical record comprised of a Patient and their Vitals
//Patients are not expected to change once created but Vitals can change constantly
type Record struct {
	Patient Patient
	Vitals  Vitals
}

//NewRecord returns a new Record struct
func NewRecord(p Patient, v Vitals) Record {
	return Record{
		Patient: p,
		Vitals:  v,
	}
}
