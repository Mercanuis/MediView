package model

import (
	"github.com/google/uuid"
)

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

//PatientVitalHistory represents an aggregation of a Patient's vitals
type PatientVitalHistory struct {
	Id       uuid.UUID     `json:"id"`
	BPA      bloodPressure `json:"avgBloodPressure"`
	PAvg     int           `json:"avgPulse"`
	GAvg     int           `json:"avgGlucose"`
	bpaCount int
	pCount   int
	gCount   int
}

type PatientVitalHistories struct {
	Histories []PatientVitalHistory `json:"histories"`
}

//NewPatientVitalHistory returns a new instance of PatientsVitalHistory
func NewPatientVitalHistory(pid uuid.UUID, bpa bloodPressure, pul int, glu int) PatientVitalHistory {
	return PatientVitalHistory{
		Id:       pid,
		BPA:      bpa,
		bpaCount: 1,
		PAvg:     pul,
		pCount:   1,
		GAvg:     glu,
		gCount:   1,
	}
}

func (h *PatientVitalHistory) UpdateHistory(sys, dys, pul, glu int) {
	h.bpaCount++
	h.BPA.Systolic = (h.BPA.Systolic + sys) / h.bpaCount
	h.BPA.Diastolic = (h.BPA.Diastolic + dys) / h.bpaCount

	h.pCount++
	h.PAvg = (h.PAvg + pul) / h.pCount

	h.gCount++
	h.GAvg = (h.GAvg + glu) / h.gCount
}
