package receiver

import (
	"MediView/http/dto"
	"MediView/service"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

const (
	queueName = "MEDIVIEW"
	queueURL  = "amqp://guest:guest@localhost:5672/"
)

func onError(err error, msg string) {
	if err != nil {
		log.Printf("[queue] %s: %s", msg, err)
	}
}

//Receiver is an interface that defines methods for pulling messages from the message queue
type Receiver interface {
	ConsumeFromQueue() error
	AddPatientReceiver
	AddRecordReceiver
	HistoryReceiver
}

type receiverCache struct {
	receiverConn *amqp.Connection
	service      *service.Service
}

//NewReceiver creates a new Receiver
func NewReceiver(s *service.Service) Receiver {
	r := &receiverCache{
		service: s,
	}
	return r
}

//ConsumeFromQueue initializes the receiver for the message queue
//Once alive, the receiver will listen for messages and then
//determine the message type to call the proper method in the service
func (r *receiverCache) ConsumeFromQueue() error {
	conn, err := amqp.Dial(queueURL)
	onError(err, "[receive] failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	onError(err, "[receive] failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	onError(err, "[receive] failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	onError(err, "[receive] failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[receive] received a message: %s", d.Body)
			dt := decodeType(d.Body)
			switch dt.Type {
			case dto.TypePAR:
				r.AddPatient(d.Body)
				break
			case dto.TypeRAR:
				r.AddRecord(d.Body)
				break
			case dto.TypeRHR:
				r.ResetHistory()
				break
			case dto.TypeDHR:
				r.DeleteHistory()
				break
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

func decodeType(str []byte) dto.RequestType {
	var d dto.RequestType
	err := json.Unmarshal(str, &d)
	if err != nil {
		onError(err, "[receive] failed to decode message off queue")
	}

	return d
}
