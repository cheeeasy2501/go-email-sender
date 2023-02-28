package service

import (
	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/pkg/logger"
)

type Services struct {
	l logger.ILogger
	SenderService ISenderService
}

func NewServices(cfg *config.Config, l logger.ILogger) *Services {
	return &Services{
		l: l,
		SenderService: NewSenderService(cfg),
	}
}
