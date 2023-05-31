package xsignal

import (
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignal = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}

type stopFunc func(graceful bool)

func WaitShutdown(stop stopFunc) {
	cc := make(chan os.Signal, 2)
	signal.Notify(cc, shutdownSignal...)
	go func() {
		sig := <-cc
		stop(sig != syscall.SIGQUIT)
	}()
}
