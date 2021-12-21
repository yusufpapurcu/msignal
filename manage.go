package msignal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func DefaultCtx() context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
