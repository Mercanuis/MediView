package http

import (
	"MediView/data"
	"MediView/service"
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func testListener(t *testing.T) net.Listener {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	return ln
}

func testService(t *testing.T) service.Service {
	t.Helper()
	m := data.NewMemCache()
	s := service.NewService(m)
	return s
}

func startHTTPCServer(t *testing.T) (*Server, net.Listener) {
	t.Helper()
	server, err := New(testService(t))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	ln := testListener(t)
	go func() {
		server.Serve(ln)
	}()

	return server, ln
}

func TestHTTPServer(t *testing.T) {
	cases := map[string]struct {
		path   string
		status int
	}{
		"Records": {
			"/getRecords",
			http.StatusOK,
		},
		"InvalidPath": {
			"/badPath",
			http.StatusNotFound,
		},
	}

	server, ln := startHTTPCServer(t)
	defer func() {
		if err := server.GracefulStop(context.TODO()); err != nil {
			t.Fatalf("err: %s", err)
		}
	}()

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			resp, err := http.Get("http://" + ln.Addr().String() + tc.path)
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			assert.Equal(t, resp.StatusCode, tc.status)
		})
	}
}

type failedListener struct {
	net.Listener
}

func (l *failedListener) Accept() (net.Conn, error) {
	return nil, fmt.Errorf("failed to accept")
}

func TestServerServeFailed(t *testing.T) {
	server, err := New(testService(t))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer func() {
		if err := server.GracefulStop(context.TODO()); err != nil {
			t.Fatalf("err: %s", err)
		}
	}()

	ln := testListener(t)
	err = server.Serve(&failedListener{ln})
	if err == nil {
		t.Fatalf("failure expected")
	}
	assert.Error(t, err, "failed to accept")
}
