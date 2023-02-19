package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/cheeeasy2501/go-email-sender/cmd/app"
	"github.com/cheeeasy2501/go-email-sender/config"
)

func main() {

	cfg, err := config.NewConfig(".", "env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	notifyCtx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGKILL)
	defer func() {
		cancel()
	}()

	if err != nil {
		panic(err)
	}

	app.Run(cfg)

	<-notifyCtx.Done()
}
