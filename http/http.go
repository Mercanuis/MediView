package http

import (
	"MediView/http/dto"
	"MediView/queue"
	"MediView/queue/receiver"
	"MediView/service"
	"context"
	"log"
	"net"
	"net/http"
	"os"
)

//Server represents the HTTP main
type Server struct {
	//Application specific logic and services
	MediService *service.Service

	//Third-party logic (HTTP, Logs)
	log      *log.Logger
	mux      *http.ServeMux
	server   *http.Server
	sender   queue.Sender
	receiver receiver.Receiver
}

//New returns a new Service
func New(s *service.Service, r receiver.Receiver) (*Server, error) {
	server := &Server{
		MediService: s,
		log:         log.New(os.Stderr, "", log.LstdFlags),
		mux:         http.NewServeMux(),
		sender:      queue.NewSender(),
		receiver:    r,
	}

	server.registerHandlers()
	return server, nil
}

func (s *Server) registerHandlers() {
	s.mux.Handle("/getRecords", s.getRecordsHandler())
	s.mux.Handle("/addPatient", s.addPatientHandler())
	s.mux.Handle("/addRecord", s.addRecordHandler())
	s.mux.Handle("/getHistories", s.getHistoryHandler())
}

// Serve starts accept requests from the given listener. If any returns error.
func (s *Server) Serve(ln net.Listener) error {
	server := &http.Server{
		Handler: s.mux,
	}
	s.server = server

	// ErrServerClosed is returned by the Server's Serve
	// after a call to Shutdown or Close, we can ignore it.
	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

//StartReceiver initializes the server's receiver
func (s *Server) StartReceiver() error {
	return s.receiver.ConsumeFromQueue()
}

//ResetData adds a 'reset' message to the message queue
func (s *Server) ResetData() {
	rhr := dto.ResetHistoryRequest{}
	rhr.Type = dto.TypeRHR
	s.sender.ResetHistorySender(rhr)
}

//DeleteData adds a 'delete' message to the message queue
func (s *Server) DeleteData() {
	dhr := dto.DeleteHistoryRequest{}
	dhr.Type = dto.TypeDHR
	s.sender.DeleteHistorySender(dhr)
}

// GracefulStop gracefully shuts down the main without interrupting any
// active connections. If any returns error.
func (s *Server) GracefulStop(ctx context.Context) error {
	var err error
	if e := s.server.Shutdown(ctx); e != nil {
		err = e
	}
	return err
}
