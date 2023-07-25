package app

import (
	"bytes"
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
	amqp "github.com/cheeeasy2501/go-email-sender/internal/transport/amqp"
	v1 "github.com/cheeeasy2501/go-email-sender/internal/transport/grpc/v1"
	"github.com/cheeeasy2501/go-email-sender/pkg/logger"
	"google.golang.org/grpc"
)

type App struct {
	ctx context.Context
	cfg *config.Config
	l   logger.ILogger
	s   *service.Services
}

func NewApp(ctx context.Context, cfg *config.Config, l logger.ILogger, s *service.Services) (*App, error) {
	return &App{
		ctx: ctx,
		cfg: cfg,
		l:   l,
		s:   s,
	}, nil
}

// App entry point method
func (a *App) Run(cfg *config.Config) error {
	err := a.RunGRPC()
	if err != nil {
		return err
	}
	// TODO: тест amqp
	r := amqp.NewReceiver(cfg)

	err = r.Connect()
	if err != nil {
		panic(err)
	}
	err = r.DeclareQueue()
	if err != nil {
		panic(err)
	}
	err = r.CreateTestConsumer()
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range r.GetDeliveryChan() {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	go func() {
		for {
			r.AddTestPublish()
			log.Println("------- Test message published -------")
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

// Run GRPC server
func (a *App) RunGRPC() error {
	if a.cfg.GRPC().IsGRPCEnable() == false {
		a.l.Instance().Info("GRPC start was interrupted. GRPC_ENABLE is " + strconv.FormatBool(a.cfg.GRPC().IsGRPCEnable()))
		return nil
	}

	grpcServer := grpc.NewServer()
	grpcSenderSevice := v1.NewGRPCSenderServer(a.cfg, a.l, a.s.SenderService)
	sender.RegisterSenderServer(grpcServer, grpcSenderSevice)

	l, err := net.Listen("tcp", a.cfg.GRPC().GetListenerAddr())
	if err != nil {
		return err
	}

	a.l.Instance().Info("Run GRPC goroutine")
	go func() {
		if err := grpcServer.Serve(l); err != nil {
			a.l.Instance().Error("GRPC is not started!", err)
			a.ctx.Done()
		}
	}()

	a.l.Instance().Info("GRPS server started")

	return nil
}
