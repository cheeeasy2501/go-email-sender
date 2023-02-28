package service

import "github.com/cheeeasy2501/go-email-sender/config"

type Services struct {
	SenderService ISenderService
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		SenderService: NewSenderService(cfg),
	}
}
