package http

import (
	"MediView/service"
	"context"
	"net"
	"net/http"

	"go.uber.org/zap"
)

//Server represents the HTTP main
type Server struct {
	//Application specific logic and services
	MediService service.Service

	//Third-party logic (HTTP, Logs)
	log    *zap.Logger //TODO: keep this for the package to be found for later, implement later
	mux    *http.ServeMux
	server *http.Server
}

//New returns a new Service
func New(s service.Service) (*Server, error) {
	server := &Server{
		MediService: s,
		mux:         http.NewServeMux(),
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

// GracefulStop gracefully shuts down the main without interrupting any
// active connections. If any returns error.
func (s *Server) GracefulStop(ctx context.Context) error {
	var err error
	if e := s.server.Shutdown(ctx); e != nil {
		err = e
	}
	return err
}
