package receiver

import (
	"MediView/http/dto"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

type AddRecordReceiver interface {
	AddRecord([]byte)
}

func (r *receiverCache) AddRecord(body []byte) {
	log.Printf("found RAR")
	rec := decodeRAR(body)
	_, err := r.service.AddRecord(rec.ID, rec.Systolic, rec.Diastolic, rec.Pulse, rec.Glucose)
	if err != nil {
		log.Print(errors.Wrap(err, "failed to add record"))
	}
}

func decodeRAR(str []byte) dto.RecordAddRequest {
	var d dto.RecordAddRequest
	err := json.Unmarshal(str, &d)
	if err != nil {
		failOnError(err, "failed to decode message off queue")
	}

	return d
}
