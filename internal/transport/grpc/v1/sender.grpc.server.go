package v1

import (
	"context"
	"net/http"

	"github.com/cheeeasy2501/go-email-sender/gen/sender"
)

type GRPCSenderServer struct {
	sender.UnimplementedSenderServer
}

func (s *GRPCSenderServer) SendMail(context.Context, *sender.SendMailRequest) (*sender.SendMailResponse, error) {
	return &sender.SendMailResponse{
		Sent: true,
		Err: &sender.Error{
			Text: "mail is sent",
			Code: http.StatusOK,
		},
	}, nil
}
