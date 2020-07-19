package log_test

import (
	"errors"
	"testing"

	"github.com/josielsousa/challenge-accounts/providers/log"
)

func TestLogInfo(t *testing.T) {
	t.Run("Teste log info", func(t *testing.T) {
		logger := log.NewLogger()
		logger.Info("Test info...")
	})
}

func TestLogError(t *testing.T) {
	t.Run("Teste log error", func(t *testing.T) {
		logger := log.NewLogger()
		logger.Error("Test error: ", errors.New("ERROR"))
	})
}
