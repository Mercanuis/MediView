package receiver

import (
	"MediView/http/dto"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

//AddPatientReceiver defines methods for calling the service
type AddPatientReceiver interface {
	AddPatient([]byte)
}

//AddPatient adds a new patient to the system data store
func (r *receiverCache) AddPatient(body []byte) {
	log.Printf("[queue] request received for new patient")
	patient := decodePAT(body)
	err := r.service.AddPatient(patient.Name, patient.Age)
	if err != nil {
		log.Print(errors.Wrap(err, "[add patient] failed to add patient"))
	}
}

func decodePAT(str []byte) dto.PatientAddRequest {
	var d dto.PatientAddRequest
	err := json.Unmarshal(str, &d)
	if err != nil {
		onError(err, "failed to decode message off queue")
	}

	return d
}
