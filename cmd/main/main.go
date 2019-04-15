package main

import (
	"MediView/di"
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
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	//Initialize HTTP
	container := di.NewContainer()
	httpServer, err := container.GetHTTPServer()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup new HTTP server: %s\n", err)
		return exitError
	}

	httpLn, err := setupListener(httpPort)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to listen HTTP port: %s\n", err)
		return exitError
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Go routines to run the following
	//- Call the HTTP Server to initialize the HTTP handlers
	//- Initialize the receiver to consume messages from the handler
	//- Call the reset function every hour
	//- Call the delete function every 24 hours
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return httpServer.Serve(httpLn) })
	wg.Go(func() error { return httpServer.StartReceiver() })

	//Resets the history every hour
	wg.Go(func() error {
		reset := make(chan os.Signal, 1)
		signal.Notify(reset, syscall.SIGTERM, os.Interrupt)
		doReset := true
		select {
		case <-reset:
			doReset = false
		}
		for doReset {
			_ = time.AfterFunc(60*time.Minute, func() {
				httpServer.ResetData()
			})
		}
		return nil
	})

	//Deletes the data every day
	wg.Go(func() error {
		del := make(chan os.Signal, 1)
		signal.Notify(del, syscall.SIGTERM, os.Interrupt)
		doReset := true
		select {
		case <-del:
			doReset = false
		}
		for doReset {
			_ = time.AfterFunc(24*time.Hour, func() {
				httpServer.DeleteData()
			})
		}
		return nil
	})

	//Handle shutdown from SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		log.Print("[DEBUG] Received SIGTERM signal, shutting down HTTP server\n")
	case <-ctx.Done():
	}

	// Gracefully shutdown HTTP main.
	if err := httpServer.GracefulStop(ctx); err != nil {
		log.Fatalf("[ERROR] Failed to gracefully shutdown HTTP Server: %s\n", err)
	}

	cancel()
	if err := wg.Wait(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] unhandled error received: %s\n", err)
		return exitError
	}

	return exitOk
}

func setupListener(port int) (net.Listener, error) {
	addr := fmt.Sprintf(":%d", port)
	return net.Listen("tcp", addr)
}
