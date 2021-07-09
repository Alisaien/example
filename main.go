package main

import (
	"context"
	"github.com/Alisaien/example/pkg/core"
	"github.com/Alisaien/example/pkg/routes"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// use context
	ctx, cancel := context.WithCancel(context.Background())

	core.Connect(ctx)
	defer core.DBW.Close()
	defer core.DBR.Close()

	server := http.Server{
		Addr: ":8080",
		Handler: routes.Mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	server.RegisterOnShutdown(cancel)

	// intercept SIGINT, so we can shutdown the server and close db connections gracefully
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown

	_ = server.Shutdown(context.Background())
}