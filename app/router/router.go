package router

import (
	"github.com/black-dragon74/rht-trn/initialize"
	"net/http"

	"github.com/black-dragon74/rht-trn/app/handlers"
	"github.com/black-dragon74/rht-trn/app/middlewares"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(lgr *zap.Logger, d *initialize.DataStore) *mux.Router {
	rtr := mux.NewRouter()
	rtr.Use(middlewares.WithContentJSON)

	// Default
	rtr.HandleFunc("/", handlers.WelcomeHandler(lgr)).Methods(http.MethodGet)

	// Read
	rtr.HandleFunc("/read", handlers.StudentRead(lgr, d)).Methods(http.MethodGet)

	// Write
	rtr.HandleFunc("/write", handlers.StudentWrite(lgr, d))

	lgr.Info("[Router] [NewRouter] Loaded the routes for the server")
	return rtr
}
