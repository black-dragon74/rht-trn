package initialize

import (
	"log"

	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	lgr, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		log.Fatal("Unable to init the logger")
	}

	lgr.Info("[Initialize] [Logger] Loaded the zap logger")

	return lgr
}
