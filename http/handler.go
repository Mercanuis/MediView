package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func (s *Server) getRecordsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		records := s.MediService.GetLatestRecords()
		if err := writeJSON(w, http.StatusOK, records); err != nil {
			//TODO: Figure out what to do with error
		}
		_, _ = fmt.Fprint(os.Stdout, "Success!")
	})
}

func writeJSON(w http.ResponseWriter, code int, i interface{}) error {
	setContentType(w, code, "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return errors.Wrap(err, "unable to write JSON")
	}
	return nil
}

func setContentType(w http.ResponseWriter, code int, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
}
