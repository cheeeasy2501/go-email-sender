package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/cheeeasy2501/go-email-sender/cmd/app"
	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
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

	s := service.NewServices(cfg)

	a, err := app.NewApp(ctx, cfg, s)
	if err != nil {
		panic(err)
	}

	err = a.Run(cfg)
	if err != nil {
		panic(err)
	}

	log.Println("App is running")

	<-notifyCtx.Done()

	log.Println("App has been stopped")
}
