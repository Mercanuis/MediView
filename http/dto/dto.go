package dto

import "github.com/google/uuid"

const (
	TypePAR = "PAR"
	TypeRAR = "RAR"
	TypeRHR = "RHR"
	TypeDHR = "DHR"
)

type (
	RequestType struct {
		Type string `json:"type"`
	}

	//PatientAddRequest represents a request to add a Patient
	PatientAddRequest struct {
		RequestType
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//RecordAddRequest represents a request to add a Record
	RecordAddRequest struct {
		RequestType
		ID        uuid.UUID `json:"id"`
		Systolic  int       `json:"systolic"`
		Diastolic int       `json:"diastolic"`
		Pulse     int       `json:"pulse"`
		Glucose   int       `json:"glucose"`
	}

	ResetHistoryRequest struct {
		RequestType
	}

	DeleteHistoryRequest struct {
		RequestType
	}

	//ErrorResponse represents an error response
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		Cause    string `json:"cause"`
	}
)
