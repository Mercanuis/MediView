package main

import (
	"MediView/di"
	"MediView/http"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	exitOk = iota
	exitError

	httpPort int = 20001

	defaultResetTime  = 60 * time.Minute
	defaultDeleteTime = 24 * time.Hour

	shortResetTime  = 2 * time.Minute
	shortDeleteTime = 5 * time.Minute
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	//Initialize HTTP
	container := di.NewContainer()
	httpServer, err := container.GetHTTPServer()
	if err != nil {
		log.Fatalf("[ERROR] Failed to setup HTTP Server: %s\n", err)
		return exitError
	}

	receiver, err := container.GetReceiver()
	if err != nil {
		log.Fatalf("[ERROR] Failed to setup queue: %s\n", err)
		return exitError
	}

	httpLn, err := setupListener(httpPort)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to listen HTTP port: %s\n", err)
		return exitError
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resTime := time.NewTicker(defaultResetTime)
	delTime := time.NewTicker(defaultDeleteTime)
	if len(args) != 1 && args[1] == "short" {
		resTime = time.NewTicker(shortResetTime)
		delTime = time.NewTicker(shortDeleteTime)
	}

	//Go routines to run the following
	//- Call the HTTP Server to initialize the HTTP handlers
	//- Initialize the receiver to consume messages from the handler
	//- Call the reset function every hour
	//- Call the delete function every 24 hours
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return httpServer.Serve(httpLn) })
	wg.Go(func() error { return receiver.ConsumeFromQueue() })
	wg.Go(func() error { return resetTimer(resTime, httpServer) })
	wg.Go(func() error { return deleteTimer(delTime, httpServer) })

	//SIGTERM for goroutines
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		log.Print("[DEBUG] Received SIGTERM signal, shutting down HTTP server\n")
		if err := httpServer.GracefulStop(ctx); err != nil {
			log.Fatalf("[ERROR] Failed to gracefully shutdown HTTP Server: %s\n", err)
		}
		log.Print("Http shut down, finished")
		break
	case <-ctx.Done():
	}

	return exitOk
}

func setupListener(port int) (net.Listener, error) {
	addr := fmt.Sprintf(":%d", port)
	return net.Listen("tcp", addr)
}

func resetTimer(t *time.Ticker, s *http.Server) error {
	for {
		select {
		case <-t.C:
			s.ResetData()
		}
	}
}

func deleteTimer(t *time.Ticker, s *http.Server) error {
	for {
		select {
		case <-t.C:
			s.DeleteData()
		}
	}
}
