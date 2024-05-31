package main

import (
	"github.com/usagiga/yagisan/misc/cli"
	"github.com/usagiga/yagisan/misc/smtp"
	"log/slog"
	"os"

	// log driver
	_ "github.com/usagiga/yagisan/misc/log"
)

func main() {
	server := smtp.NewSmtpServer()
	slog.Info("server initialized. ready to serve")

	err := cli.WithGracefulShutdown(
		server.ListenAndServeTLS,
		server.Shutdown,
	)
	if err != nil {
		slog.Error("error raised in server", "err", err)
		os.Exit(1)
		return
	}

	slog.Info("server stopped. bye")
}
