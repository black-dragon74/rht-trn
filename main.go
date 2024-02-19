package main

import (
	"github.com/black-dragon74/rht-trn/app"
	"github.com/black-dragon74/rht-trn/initialize"
)

func main() {
	lgr := initialize.Logger()
	d := initialize.NewDataStore(lgr)

	app.Start(d, lgr)
}
