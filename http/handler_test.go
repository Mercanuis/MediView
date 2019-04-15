package http

import (
	"MediView/data/model"
	"MediView/http/dto"
	"MediView/http/mocks"
	"MediView/queue/receiver"
	"MediView/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"

	"gotest.tools/assert"

	"github.com/google/uuid"
)

func testHTTPServer(t *testing.T, s service.Service) *Server {
	t.Helper()
	rec := receiver.NewReceiver(&s)
	server, err := New(&s, rec)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	return server
}

func TestGetRecordsHandler(t *testing.T) {
	cases := map[string]struct {
		httpMethod     string
		errorCode      int
		GetPatientMock func() model.PatientRecords
	}{
		"GET": {
			httpMethod: http.MethodGet,
			errorCode:  200,
			GetPatientMock: func() model.PatientRecords {
				return model.PatientRecords{
					Records: []model.Patient{model.NewPatient(uuid.New(), "Joe", 33)},
				}
			},
		},
		"PUT": {
			httpMethod: http.MethodPut,
			errorCode:  405,
			GetPatientMock: func() model.PatientRecords {
				return model.PatientRecords{}
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			testService := service.NewService(&mocks.DaoMock{
				GetPatientsMock: tc.GetPatientMock,
			})
			server := testHTTPServer(t, testService)

			handlerFunc := server.getRecordsHandler()
			rr := httptest.NewRecorder()
			handlerFunc.ServeHTTP(rr, httptest.NewRequest(tc.httpMethod, "/getRecords", nil))

			resp := rr.Result()
			assert.Equal(t, resp.StatusCode, tc.errorCode, "response: %v, tc: %v", resp.StatusCode, tc.errorCode)
		})
	}
}

func TestAddPatientHandler(t *testing.T) {
	cases := map[string]struct {
		httpMethod     string
		errorCode      int
		AddRequest     dto.PatientAddRequest
		AddPatientMock func(name string, age int) error
	}{
		"GET": {
			httpMethod: http.MethodGet,
			errorCode:  405,
			AddRequest: dto.PatientAddRequest{},
			AddPatientMock: func(name string, age int) error {
				return nil
			},
		},
		"POST": {
			httpMethod: http.MethodPost,
			errorCode:  200,
			AddRequest: dto.PatientAddRequest{
				Age:  33,
				Name: "Joey",
			},
			AddPatientMock: func(name string, age int) error {
				return nil
			},
		},
		"FailedToAddPatient": {
			httpMethod: http.MethodGet,
			errorCode:  405,
			AddRequest: dto.PatientAddRequest{
				Age:  40,
				Name: "Jim",
			},
			AddPatientMock: func(name string, age int) error {
				return errors.New("Bad Request")
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			testService := service.NewService(&mocks.DaoMock{
				AddPatientMock: tc.AddPatientMock,
			})
			server := testHTTPServer(t, testService)

			handlerFunc := server.addPatientHandler()
			rr := httptest.NewRecorder()

			body, er := json.Marshal(tc.AddRequest)
			if er != nil {
				t.Fatalf("Failed to convert to JSON: %v", er)
			}
			reader := bytes.NewReader(body)
			handlerFunc.ServeHTTP(rr, httptest.NewRequest(tc.httpMethod, "/addPatient", reader))

			resp := rr.Result()
			assert.Equal(t, resp.StatusCode, tc.errorCode, "response: %v, tc: %v", resp.StatusCode, tc.errorCode)
		})
	}
}

func TestAddRecordHandler(t *testing.T) {
	cases := map[string]struct {
		httpMethod    string
		errorCode     int
		AddRequest    dto.RecordAddRequest
		AddRecordMock func(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error)
	}{
		"GET": {
			httpMethod: http.MethodGet,
			errorCode:  405,
			AddRequest: dto.RecordAddRequest{},
			AddRecordMock: func(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
				return &model.Patient{}, nil
			},
		},
		"POST": {
			httpMethod: http.MethodPost,
			errorCode:  200,
			AddRequest: dto.RecordAddRequest{
				Systolic:  128,
				Diastolic: 70,
				Pulse:     77,
				Glucose:   45,
			},
			AddRecordMock: func(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
				return &model.Patient{}, nil
			},
		},
		"FailedToAddRecord": {
			httpMethod: http.MethodGet,
			errorCode:  405,
			AddRequest: dto.RecordAddRequest{
				Systolic:  128,
				Diastolic: 70,
				Pulse:     77,
				Glucose:   45,
			},
			AddRecordMock: func(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error) {
				return &model.Patient{}, errors.New("Bad Request")
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			testService := service.NewService(&mocks.DaoMock{
				AddRecordMock: tc.AddRecordMock,
			})
			server := testHTTPServer(t, testService)

			handlerFunc := server.addRecordHandler()
			rr := httptest.NewRecorder()

			body, er := json.Marshal(tc.AddRequest)
			if er != nil {
				t.Fatalf("Failed to convert to JSON: %v", er)
			}
			reader := bytes.NewReader(body)
			handlerFunc.ServeHTTP(rr, httptest.NewRequest(tc.httpMethod, "/addRecord", reader))

			resp := rr.Result()
			assert.Equal(t, resp.StatusCode, tc.errorCode, "response: %v, tc: %v", resp.StatusCode, tc.errorCode)
		})
	}
}

func TestGetHistoryHandler(t *testing.T) {
	cases := map[string]struct {
		httpMethod     string
		errorCode      int
		GetHistoryMock func() model.PatientVitalHistories
	}{
		"GET": {
			httpMethod: http.MethodGet,
			errorCode:  200,
			GetHistoryMock: func() model.PatientVitalHistories {
				return model.PatientVitalHistories{
					Histories: []model.PatientVitalHistory{},
				}
			},
		},
		"PUT": {
			httpMethod: http.MethodPut,
			errorCode:  405,
			GetHistoryMock: func() model.PatientVitalHistories {
				return model.PatientVitalHistories{
					Histories: []model.PatientVitalHistory{},
				}
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			testService := service.NewService(&mocks.DaoMock{
				GetPatientHistoriesMock: tc.GetHistoryMock,
			})
			server := testHTTPServer(t, testService)

			handlerFunc := server.getHistoryHandler()
			rr := httptest.NewRecorder()
			handlerFunc.ServeHTTP(rr, httptest.NewRequest(tc.httpMethod, "/getHistories", nil))

			resp := rr.Result()
			assert.Equal(t, resp.StatusCode, tc.errorCode, "response: %v, tc: %v", resp.StatusCode, tc.errorCode)
		})
	}
}
