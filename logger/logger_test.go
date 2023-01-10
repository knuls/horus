package logger

import (
	"log"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogger(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatal("logger failed to create")
	}
	defer logger.GetLogger().Sync()

	// replace logger zap/core with observed zap/core to capture written logs
	core, capturedLogs := observer.New(zap.InfoLevel)
	sugar := logger.log.Sugar().WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return core
	}))
	logger.SetSugar(sugar)

	// write logs
	logger.Info("some log line", "key", "value")
	logger.Infof("some log %s", "line")
	logger.Error("some error", "error", "some-error-message")

	// assert
	entry := capturedLogs.All()[0]
	if entry.Level != zap.InfoLevel || entry.Message != "some log line" || entry.ContextMap()["key"] != "value" {
		t.Fatal("logger should have written info log with message and key/value")
	}
	if capturedLogs.Len() != 3 {
		t.Fatal("logger should have captured two log entries")
	}
	if !isLogger(logger.GetLogger()) {
		t.Fatal("logger should be of type zap/logger")
	}
	if !isStandardLogger(logger.GetStdLogger()) {
		t.Fatal("logger should be of type log/logger")
	}
}

func isLogger(o interface{}) bool {
	switch o.(type) {
	case *zap.Logger:
		return true
	default:
		return false
	}
}

func isStandardLogger(o interface{}) bool {
	switch o.(type) {
	case *log.Logger:
		return true
	default:
		return false
	}
}
