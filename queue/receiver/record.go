package receiver

import (
	"MediView/http/dto"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

//AddRecordReceiver defines methods for adding records to the service
type AddRecordReceiver interface {
	AddRecord([]byte)
}

//AddRecord adds a new record to the system's data store
func (r *receiverCache) AddRecord(body []byte) {
	rec := decodeRAR(body)
	_, err := r.service.AddRecord(rec.ID, rec.Systolic, rec.Diastolic, rec.Pulse, rec.Glucose)
	if err != nil {
		log.Print(errors.Wrap(err, "[add record] failed to add record"))
	}
}

func decodeRAR(str []byte) dto.RecordAddRequest {
	var d dto.RecordAddRequest
	err := json.Unmarshal(str, &d)
	if err != nil {
		onError(err, "failed to decode message off queue")
	}

	return d
}
