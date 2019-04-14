package http

import (
	"MediView/http/dto"
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
		if err := writeJSON(w, http.StatusOK, records); err != nil {
			s.handleError(w, err, FailedToGetRecords)
			return
		}
		s.log.Print("Successfully performed GET of Records")
	})
}

func (s *Server) addPatientHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var patient dto.PatientAddRequest
		err := decoder.Decode(&patient)
		if err != nil {
			s.log.Printf("[Error] failed to decode JSON body: %v\n", err)
			s.handleError(w, err, FailedToAddPatient)
			return
		}

		err = s.MediService.AddPatient(patient.Name, patient.Age)
		if err != nil {
			s.log.Printf("[Error] failed to add patient: %v\n", err)
			s.handleError(w, err, FailedToAddPatient)
			return
		}

		setContentType(w, http.StatusOK, "application/json; charset=UTF-8")
		s.log.Print("Successfully performed POST of Patient")
	})
}

func (s *Server) addRecordHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var record dto.RecordAddRequest
		err := decoder.Decode(&record)
		if err != nil {
			s.log.Printf("[Error] failed to decode JSON body: %v\n", err)
			s.handleError(w, err, FailedToAddRecord)
			return
		}

		patient, err := s.MediService.AddRecord(
			record.ID,
			record.Systolic,
			record.Diastolic,
			record.Pulse,
			record.Glucose)
		if err != nil {
			s.log.Printf("[Error] failed to add patient record: %v\n", err)
			s.handleError(w, err, FailedToAddRecord)
		}

		if err := writeJSON(w, http.StatusOK, patient); err != nil {
			s.handleError(w, err, FailedToAddRecord)
			return
		}
		s.log.Print("Successfully performed POST of Record")
	})
}

func (s *Server) getHistoryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		histories := s.MediService.GetHistories()
		if err := writeJSON(w, http.StatusOK, histories); err != nil {
			s.handleError(w, err, FailedToGetHistory)
			return
		}
		s.log.Print("Successfully performed GET of History")
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
