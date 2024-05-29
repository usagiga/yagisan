package log

import (
	cslog "github.com/PumpkinSeed/slog-context"
	"log/slog"
	"os"
)

// Initialize logger for built-in slog.
// It relies on github.com/PumpkinSeed/slog-context.
func init() {
	logger := slog.New(cslog.NewHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: nil,
	})))

	slog.SetDefault(logger)
}
