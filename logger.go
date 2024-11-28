package ginkit

import "go.uber.org/zap"

var lg *zap.Logger

func InitLogger() {
	// Init logger
	var err error
	lg, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(lg)
}

func Logger() *zap.Logger {
	return lg
}
