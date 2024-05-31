package main

import (
	"github.com/usagiga/yagisan/misc/cli"
	"github.com/usagiga/yagisan/misc/scaffold"
	"log/slog"
	"os"

	// log driver
	_ "github.com/usagiga/yagisan/misc/log"
)

func main() {
	// Building modules
	server := scaffold.ScaffoldModules()
	slog.Info("server initialized. ready to serve")

	// Run server
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
