package dto

import "github.com/google/uuid"

type PatientAddRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RecordAddRequest struct {
	Id        uuid.UUID `json:"id"`
	Systolic  int       `json:"systolic"`
	Diastolic int       `json:"diastolic"`
	Pulse     int       `json:"pulse"`
	Glucose   int       `json:"glucose"`
}
