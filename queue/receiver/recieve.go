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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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

func NewReceiver(s *service.Service) Receiver {
	r := &receiverCache{
		service: s,
	}
	return r
}

func (r *receiverCache) ConsumeFromQueue() error {
	conn, err := amqp.Dial(queueURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
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
		failOnError(err, "failed to decode message off queue")
	}

	return d
}
