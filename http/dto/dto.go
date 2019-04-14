package dto

import "github.com/google/uuid"

type (
	//PatientAddRequest represents a request to add a Patient
	PatientAddRequest struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//RecordAddRequest represents a request to add a Record
	RecordAddRequest struct {
		ID        uuid.UUID `json:"id"`
		Systolic  int       `json:"systolic"`
		Diastolic int       `json:"diastolic"`
		Pulse     int       `json:"pulse"`
		Glucose   int       `json:"glucose"`
	}

	//ErrorResponse represents an error response
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		Cause    string `json:"cause"`
	}
)
