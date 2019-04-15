package dto

import "github.com/google/uuid"

const (
	//TypePAR represents a PatientAddRequest
	TypePAR = "PAR"
	//TypeRAR represents a RecordAddRequest
	TypeRAR = "RAR"
	//TypeRHR represents a ResetHistoryRequest
	TypeRHR = "RHR"
	//TypeDHR represents a DeleteHistoryRequest
	TypeDHR = "DHR"
)

type (
	//RequestType is a struct, meant to help the message queue determine
	//the type of request being sent to the service
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

	//ResetHistoryRequest represents a request to reset the data store
	ResetHistoryRequest struct {
		RequestType
	}

	//DeleteHistoryRequest represents a request to delete data from the data store
	DeleteHistoryRequest struct {
		RequestType
	}

	//ErrorResponse represents an error response
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		Cause    string `json:"cause"`
	}
)
