package log

import (
	cslog "github.com/PumpkinSeed/slog-context"
	"log/slog"
	"os"
)

// Usage: Just import this package
// such as `import _ "github.com/usagiga/yagisan/lib/log"`
//
// This package set default logger of slog,
// so if you want to write a log, call slog global functions (e.g. slog.Infof())
// if you want to add fields(such as Trace ID, Request ID, ...) into log in the context,
// call cslog.WithValue()

// Initialize logger for built-in slog.
// It relies on github.com/PumpkinSeed/slog-context.
func init() {
	logger := slog.New(cslog.NewHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: nil,
	})))

	slog.SetDefault(logger)
}
