package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/black-dragon74/rht-trn/app/router"
	"github.com/black-dragon74/rht-trn/initialize"
	"go.uber.org/zap"
)

func Start(d *initialize.DataStore, lgr *zap.Logger) {

	// Load the router for HTTP server
	rtr := router.NewRouter(lgr, d)

	// Create the HTTP server with our router
	srv := &http.Server{
		Addr:    ":2974",
		Handler: rtr,
	}

	// Listen for SIGINT to shut down gracefully
	go listenForGracefulShutdown(srv, lgr)

	// Start listening
	lgr.Info("[App] [Start] Starting the HTTP server on port: 2974")
	err := srv.ListenAndServe()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		lgr.Fatal("[App] [Start] Failed to start the HTTP server")
	}
}

func listenForGracefulShutdown(srv *http.Server, lgr *zap.Logger) {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, os.Interrupt)

	<-termChan
	lgr.Info("[App] [Start] Interrupt received, attempting graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		_ = srv.Shutdown(ctx)
	}()
}
