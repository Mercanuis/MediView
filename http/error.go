package http

import (
	"encoding/json"
	"net/http"

	"github.com/MediView/http/dto"
)

const (
	//FailedToGetRecords is when the system fails to get Records
	FailedToGetRecords string = "Failed to get Records"

	//FailedToAddPatient is when the system fails to add a Patient
	FailedToAddPatient string = "Failed to add a new Patient"

	//FailedToAddRecord is when the system fails to add a Record
	FailedToAddRecord string = "Failed to add a new Record"

	//FailedToGetHistory is when the system fails to get PatientHistory
	FailedToGetHistory string = "Failed to get patient history"
)

func (s *Server) handleError(w http.ResponseWriter, err error, msg string) {
	setContentType(w, http.StatusBadRequest, "application/json; charset=UTF-8")
	body, err := json.Marshal(dto.ErrorResponse{
		ErrorMsg: msg,
		Cause:    err.Error(),
	})

	_, err = w.Write(body)
	if err != nil {
		//Do nothing for now
	}
}
