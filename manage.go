package msignal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func DefaultCtx() context.Context {
	c := DefaultChannel()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func DefaultChannel() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	return c
}
