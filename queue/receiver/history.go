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
	log.Printf("found RHR")
	r.service.ResetHistory()
	log.Printf("History for service successfully reset")
}

//DeleteHistory deletes the service's history
func (r *receiverCache) DeleteHistory() {
	log.Printf("found DHR")
	r.service.DeleteHistory()
	log.Print("History for service successfully deleted")
}
