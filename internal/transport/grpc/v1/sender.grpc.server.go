package v1

import (
	"context"
	"net/http"

	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
	"github.com/cheeeasy2501/go-email-sender/pkg/logger"
)

type GRPCSenderServer struct {
	sender.UnimplementedSenderServer
	l logger.ILogger
	s service.ISenderService
}

func NewGRPCSenderServer(l logger.ILogger, s service.ISenderService) *GRPCSenderServer {
	return &GRPCSenderServer{
		l: l,
		s: s,
	}
}

func (grpc *GRPCSenderServer) SendMail(context.Context, *sender.SendMailRequest) (*sender.SendMailResponse, error) {
	// TODO: доделать передачу dto
	// TODO: изменить Err в SendResponse на что-то другое или добавить message и code
	d := dto.NewEmailDTO([]string{"test@mail.com"}, "Subject", "test html template")
	sent, err := grpc.s.SendMail(d)
	if err != nil {
		return &sender.SendMailResponse{
			Sent: sent,
			Err: &sender.Error{
				Text: "Mail isn't sent!",
				Code: http.StatusBadRequest,
			},
		}, err
	}

	return &sender.SendMailResponse{
		Sent: true,
		Err: &sender.Error{
			Text: "mail is sent",
			Code: http.StatusOK,
		},
	}, nil
}
