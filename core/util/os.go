package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RegisterOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, signals...)
		<-ch

		f()
	}()
}

func WaitOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch

	f()
}

func WaitOSSignalGracefulShutdown(f func(ctx context.Context), timeout time.Duration) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	f(ctx)
}
