package main

import (
	"MediView/di"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	exitOk = iota
	exitError

	httpPort int = 10001
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
	//- Call the HTTP main to handle any HTTP requests
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return httpServer.Serve(httpLn) })

	//Handle shutdown from SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		_, _ = fmt.Fprintf(os.Stdout, "[INFO] received SIGTERM, exiting server gracefully")
		//TODO: Implement later - logger.Info("received SIGTERM, exiting main gracefully")
	case <-ctx.Done():
	}

	// Gracefully shutdown HTTP main.
	if err := httpServer.GracefulStop(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] failed to gracefully shutdown HTTP server: %s\n", err)
		//TODO; Implement later - logger.Error("failed to gracefully shutdown HTTP main", zap.Error(err))
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
