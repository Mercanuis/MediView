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

	defaultResetTime  = 60 * time.Minute
	defaultDeleteTime = 24 * time.Hour

	shortResetTime  = 90 * time.Second
	shortDeletetime = 2 * time.Minute
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
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

	resTime := time.NewTicker(defaultResetTime)
	delTime := time.NewTicker(defaultDeleteTime)
	if args[1] == "short" {
		resTime = time.NewTicker(shortResetTime)
		delTime = time.NewTicker(shortDeletetime)
	}
	reset := resTime
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-reset.C:
				httpServer.ResetData()
				break
			case <-quit:
				reset.Stop()
				return
			}
		}
	}()

	del := delTime
	go func() {
		for {
			select {
			case <-del.C:
				httpServer.DeleteData()
				break
			case <-quit:
				del.Stop()
				return
			}
		}
	}()

	//Handle shutdown from SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		log.Print("[DEBUG] Received SIGTERM signal, shutting down HTTP server\n")
		close(quit)
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
