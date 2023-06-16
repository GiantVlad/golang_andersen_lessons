package logger

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	sugar := NewLogger("", "debug")
	sugar.Infow("failed to fetch URL",
		"url", "/home",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "/home")
	sugar.Warnf("Failed to fetch URL: %s", "/home")
	sugar.Errorf("Failed to fetch URL: %s", "/home")
}
