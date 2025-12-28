package logging

import "go.uber.org/zap"
import "github.com/wantedbear007/project_skim/internal/env"

func InitLogger() {
	// to initilize zap logger for global access

	// init dev logger
	if env.IsDevelopment() {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	}

	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

}