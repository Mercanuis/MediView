package http

import (
	"MediView/data/model"
	"MediView/http/mocks"
	"MediView/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"

	"github.com/google/uuid"
)

func testHTTPServer(t *testing.T, s service.Service) *Server {
	t.Helper()
	server, err := New(s)
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
			testService := service.NewService(mocks.DaoMock{
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