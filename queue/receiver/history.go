package receiver

import (
	"log"
)

type HistoryReceiver interface {
	ResetHistory()
	DeleteHistory()
}

func (r *receiverCache) ResetHistory() {
	log.Printf("found RHR")
	r.service.ResetHistory()
	log.Printf("History for service successfully reset")
}

func (r *receiverCache) DeleteHistory() {
	log.Printf("found DHR")
	r.service.DeleteHistory()
	log.Print("History for service successfully deleted")
}
