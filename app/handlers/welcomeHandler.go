package handlers

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func WelcomeHandler(lgr *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lgr.Info("[Handler] [WelcomeHandler] Handle /")

		_, err := fmt.Fprintf(w, `{"msg": "Welcome to this barebones HTTP server!"}`)

		if err != nil {
			lgr.Error("[Handler] [WelcomeHandler] Error in handling /")
		}
	}
}
