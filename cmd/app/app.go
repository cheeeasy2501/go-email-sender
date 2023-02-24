package app

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
	v1 "github.com/cheeeasy2501/go-email-sender/internal/transport/grpc/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	ctx context.Context
	cfg *config.Config
	l   *zap.Logger
	s   *service.Services
}

func NewApp(ctx context.Context, cfg *config.Config, s *service.Services) (*App, error) {
	return &App{
		ctx: ctx,
		cfg: cfg,
		s:   s,
	}, nil
}

// App entry point method
func (a *App) Run(cfg *config.Config) error {
	err := a.InitLogger()
	if err != nil {
		return err
	}

	err = a.RunGRPC()
	if err != nil {
		return err
	}

	return nil
}

// Initial logger for app
func (a *App) InitLogger() error {
	switch a.cfg.GetAppEnv() {
	case "development":
		l, err := zap.NewDevelopment()
		if err != nil {
			return err
		}

		a.l = l
	case "production":
		l, err := zap.NewProduction()
		if err != nil {
			return err
		}

		a.l = l
	default:
		log.Println("default")
		return errors.New("Undefiend logger")
	}

	return nil
}

// Run GRPC server
func (a *App) RunGRPC() error {
	if a.cfg.GRPC().IsGRPCEnable() == false {
		a.l.Sugar().Infoln("GRPC start was interrupted. GRPC_ENABLE is " + strconv.FormatBool(a.cfg.GRPC().IsGRPCEnable()))
		return nil
	}

	grpcServer := grpc.NewServer()
	grpcSenderSevice := v1.NewGRPCSenderServer(a.l, a.s.SenderService)
	sender.RegisterSenderServer(grpcServer, grpcSenderSevice)

	l, err := net.Listen("tcp", a.cfg.GRPC().GetListenerAddr())
	if err != nil {
		return err
	}

	log.Println("Run GRPC goroutine")
	go func() {
		if err := grpcServer.Serve(l); err != nil {
			a.l.Sugar().Errorln("GRPC is not started!")
			a.l.Sugar().Error(err)
			a.ctx.Done()
		}
	}()

	log.Println("GRPS server started")

	return nil
}
