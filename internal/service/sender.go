package service

import (
	"net/smtp"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/internal/builder"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
)

type ISenderService interface {
	Send(d dto.IEmailDTO) (bool, error)
}

type SenderService struct {
	cfg *config.Config
	b   *builder.MailBuilder
}

func NewSenderService(cfg *config.Config) *SenderService {
	return &SenderService{
		cfg: cfg,
	}
}

func (s *SenderService) Send(d dto.IEmailDTO) (bool, error) {
	addr := s.cfg.Mail().GetAddress()
	host := s.cfg.Mail().GetHost()
	user := s.cfg.Mail().GetUsername()
	password := s.cfg.Mail().GetPassword()

	msg, err := s.b.Build(d)
	if err != nil {
		return false, err
	}

	auth := smtp.PlainAuth("", user, password, host)

	err = smtp.SendMail(addr, auth, s.cfg.Mail().GetAddressFrom(), d.To(), msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
