package http

import (
	"MediView/data/model"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (s *Server) getRecordsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		records := s.MediService.GetLatestRecords()
		patRecords := model.PatientRecords{}
		for _, i := range records {
			_ = append(patRecords.Records, i)
		}
		if err := writeJSON(w, http.StatusOK, patRecords); err != nil {
			//TODO: Figure out what to do with error
		}
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
