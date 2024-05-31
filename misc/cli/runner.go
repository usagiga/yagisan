package cli

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

type VoidRunner func() error
type CtxRunner func(ctx context.Context) error

// WithGracefulShutdown is task runner
// which implements Graceful Shutdown through catching some signals.
//
// It's suitable for `(*Server).ListenAndServe()` in `go-smtp`
func WithGracefulShutdown(task VoidRunner, onGraceful CtxRunner) error {
	errCh := make(chan error)
	ctx, _ := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,   // Keyboard Interrupt
		syscall.SIGTERM,  // Docker Daemon / systemd
		syscall.SIGHUP,   // Hung up
		syscall.SIGWINCH, // apache-like
		syscall.SIGQUIT,  // nginx-like
		// SIGKILL and SIGSTOP are caught by Kernel.
		// Applications should ignore it.
		// ref. `man 7 signal`
	)

	// prepare graceful shutdown
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		slog.ErrorContext(ctx, "unrecoverable error raised. proceed to shutdown gracefully", "err", err)

		err = gracefulShutdown(onGraceful)
		if err != nil {
			slog.ErrorContext(ctx, "unhandled error raised. quiting...", "err", err)
		}
	}()

	// run task
	go func() {
		err := task()
		if err != nil {
			errCh <- err
		}
	}()
	slog.InfoContext(ctx, fmt.Sprintf("server started. listening on %s", "localhost:10587")) // temp

	select {
	// task exited
	case err := <-errCh:
		return errors.Wrap(err, "task exited with error")

	// signals received
	// proceed to graceful shutdown
	case <-ctx.Done():
		err := errors.Wrap(ctx.Err(), "signal received")
		panic(err)
	}
}

func gracefulShutdown(onGraceful CtxRunner) error {
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		err := onGraceful(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	select {
	// graceful shutdown task exited
	case err := <-errCh:
		if err != nil {
			return errors.Wrap(err, "graceful shutdown exited with error")
		}
	// timeout
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			return errors.Wrap(err, "graceful shutdown timed out")
		}
	}

	return nil
}
