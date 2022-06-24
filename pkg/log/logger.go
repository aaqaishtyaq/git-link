package log

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetUpLogs set the log output ans the log level
func SetUpLogs(out io.Writer, level string) *zap.SugaredLogger {
	zapLevel := zapcore.InfoLevel
	if level == "debug" {
		zapLevel = zapcore.DebugLevel
	}

	zapConfig := zap.NewProductionEncoderConfig()
	consoleEncoder := newConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapLevel),
	)

	log := zap.New(core)
	return log.Sugar()
}

func newConsoleEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	// if interactive terminal, make output more human-readable by default
	config.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format("2006/01/02 15:04:05.000"))
	}

	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(config)
}
