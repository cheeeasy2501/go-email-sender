package app

import (
	"context"
	"net"
	"strconv"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
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
