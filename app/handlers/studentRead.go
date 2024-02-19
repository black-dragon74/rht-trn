package handlers

import (
	"encoding/json"
	"github.com/black-dragon74/rht-trn/initialize"
	"net/http"

	"go.uber.org/zap"
)

func StudentRead(lgr *zap.Logger, d *initialize.DataStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handlers] [StudentRead] Handle /read")

		data, err := json.Marshal(d.Students)
		if err != nil {
			lgr.Error("[Handlers] [StudentRead] Error while marshaling to JSON")
		}

		_, err = writer.Write(data)
		if err != nil {
			lgr.Error("[Handlers] [StudentRead] Error in writing the response")
		}
	}
}
