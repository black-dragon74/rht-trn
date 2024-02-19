package handlers

import (
	"fmt"
	"github.com/black-dragon74/rht-trn/initialize"
	"github.com/black-dragon74/rht-trn/types"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func StudentWrite(lgr *zap.Logger, d *initialize.DataStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		lgr.Info("[Handlers] [StudentWrite] Handle /write")

		if request.Method != http.MethodPost {
			lgr.Error("[Handlers] [StudentWrite] Accepts only HTTP POST requests")
			http.Error(writer, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		name := request.FormValue("name")
		age := request.FormValue("age")
		class := request.FormValue("class")

		if name == "" || age == "" || class == "" {
			lgr.Error("[Handlers] [StudentWrite] Bad form data")
			return
		}

		iAge, err := strconv.Atoi(age)
		if err != nil {
			lgr.Error("[Handlers] [StudentWrite] Error while parsing age")
		}

		d.AddStudent(&types.Student{
			Name:  name,
			Age:   iAge,
			Class: class,
		})

		_, err = fmt.Fprintf(writer, `"msg": "Student added successfully"`)
	}
}
