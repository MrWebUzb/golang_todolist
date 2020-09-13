package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/MrWebUzb/apiserver/app/server"
	"github.com/sirupsen/logrus"
)

var (
	logLevel = logrus.DebugLevel
)

func main() {
	srv := server.NewServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		srv.Log.Warnf("System call: %+v", oscall)
		cancel()
	}()

	if err := srv.Start(ctx, logLevel); err != nil {
		srv.Log.Warnf("Failed to serve: %+v", err)
	}
}
