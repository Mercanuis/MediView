package model

import (
	"github.com/google/uuid"
)

type (
	//BloodPressure represents blood pressure
	BloodPressure struct {
		Systolic  int `json:"systolic"`
		Diastolic int `json:"diastolic"`
	}

	//Vitals is a representation of a patients current vital signs
	Vitals struct {
		Pressure BloodPressure `json:"pressure"`
		Pulse    int           `json:"pulse"`
		Glucose  int           `json:"glucose"`
	}

	//Patient represents a patient's information in the system
	Patient struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Age    int       `json:"age"`
		Vitals Vitals    `json:"vitals"`
	}

	//PatientRecords represents a series of Patients
	PatientRecords struct {
		Records []Patient `json:"records"`
	}

	//PatientVitalHistory represents an aggregation of a Patient's vitals
	PatientVitalHistory struct {
		ID       uuid.UUID     `json:"id"`
		BPA      BloodPressure `json:"avgBloodPressure"`
		PAvg     int           `json:"avgPulse"`
		GAvg     int           `json:"avgGlucose"`
		bpaCount int
		pCount   int
		gCount   int
	}

	//PatientVitalHistories represent a collection of history
	PatientVitalHistories struct {
		Histories []PatientVitalHistory `json:"histories"`
	}
)

//NewVitals returns a new Vitals struct
func NewVitals(sys, dys, pulse, glu int) Vitals {
	return Vitals{
		Pressure: BloodPressure{
			Systolic:  sys,
			Diastolic: dys,
		},
		Pulse:   pulse,
		Glucose: glu,
	}
}

//NewPatient returns a new Patient struct
func NewPatient(i uuid.UUID, n string, a int) Patient {
	return Patient{
		ID:   i,
		Name: n,
		Age:  a,
	}
}

//NewPatientVitalHistory returns a new instance of PatientsVitalHistory
func NewPatientVitalHistory(pid uuid.UUID, bpa BloodPressure, pul int, glu int) PatientVitalHistory {
	return PatientVitalHistory{
		ID:       pid,
		BPA:      bpa,
		bpaCount: 1,
		PAvg:     pul,
		pCount:   1,
		GAvg:     glu,
		gCount:   1,
	}
}

//UpdateHistory updates the calling PatientVitalHistory
func (h *PatientVitalHistory) UpdateHistory(sys, dys, pul, glu int) {
	h.bpaCount++
	h.BPA.Systolic = (h.BPA.Systolic + sys) / h.bpaCount
	h.BPA.Diastolic = (h.BPA.Diastolic + dys) / h.bpaCount

	h.pCount++
	h.PAvg = (h.PAvg + pul) / h.pCount

	h.gCount++
	h.GAvg = (h.GAvg + glu) / h.gCount
}
