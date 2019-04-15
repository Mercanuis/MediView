package queue

import (
	"MediView/http/dto"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

const (
	queueName = "MEDIVIEW"
	queueURL  = "amqp://guest:guest@localhost:5672/"
)

//Sender is an interface to define calling the message queue to add a message
type Sender interface {
	AddPatientSender(par dto.PatientAddRequest)
	AddRecordSender(rar dto.RecordAddRequest)
	ResetHistorySender(rhr dto.ResetHistoryRequest)
	DeleteHistorySender(dhr dto.DeleteHistoryRequest)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type senderCache struct {
	senderConn *amqp.Connection
}

//NewSender creates a new Sender
func NewSender() Sender {
	return &senderCache{}
}

//AddPatientSender sends a request to add a patient
func (s *senderCache) AddPatientSender(par dto.PatientAddRequest) {
	s.send(par)
}

//AddRecordSender sends a request to add a record
func (s *senderCache) AddRecordSender(rar dto.RecordAddRequest) {
	s.send(rar)
}

//ResetHistorySender sends a request to reset history
func (s *senderCache) ResetHistorySender(rhr dto.ResetHistoryRequest) {
	s.send(rhr)
}

//DeleteHistorySender sends a request to delete history
func (s *senderCache) DeleteHistorySender(dhr dto.DeleteHistoryRequest) {
	s.send(dhr)
}

func (s *senderCache) send(data interface{}) {
	encoded, err := s.encodeMessage(data)
	if err != nil {
		failOnError(err, "failed to encode data")
	}
	s.publishToQueue(encoded)
}

func (s *senderCache) encodeMessage(data interface{}) ([]byte, error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		failOnError(err, "failed to marshall data")
	}
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err = e.Encode(data)
	if err != nil {
		return nil, err
	}

	return marshaled, nil
}

func (s *senderCache) publishToQueue(msg []byte) {
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

	body := msg
	err = ch.Publish("", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
