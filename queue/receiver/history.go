package receiver

import (
	"log"
)

//HistoryReceiver defines a series of methods pertaining to the
//service history
type HistoryReceiver interface {
	ResetHistory()
	DeleteHistory()
}

//ResetHistory resets the service's history
func (r *receiverCache) ResetHistory() {
	log.Printf("[queue] recieved request to reset patient history")
	r.service.ResetHistory()
	log.Printf("[queue] patient history successfully reset")
}

//DeleteHistory deletes the service's history
func (r *receiverCache) DeleteHistory() {
	log.Printf("[queue] received request to delete patient history")
	r.service.DeleteHistory()
	log.Print("[queue] patient history successfully deleted")
}
