package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var Logger = New()

// New - создание нового логгера.
func New() *zap.Logger {
	//logStartTime := time.Now()
	//logName := strconv.Itoa(logStartTime.Day()) + "-" + strconv.Itoa(int(logStartTime.Month())) + "-" + strconv.Itoa(logStartTime.Year()) +".log"
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}