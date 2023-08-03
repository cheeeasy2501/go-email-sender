package v1

import (
	"context"
	"fmt"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
	"github.com/cheeeasy2501/go-email-sender/internal/service"
	"github.com/cheeeasy2501/go-email-sender/pkg/logger"
)

type GRPCSenderServer struct {
	sender.UnimplementedSenderServer
	c *config.Config
	l logger.ILogger
	s service.ISenderService
}

// TODO: подумать стоит ли прокидывать cfg сюда
func NewGRPCSenderServer(с *config.Config, l logger.ILogger, s service.ISenderService) *GRPCSenderServer {
	return &GRPCSenderServer{
		c: с,
		l: l,
		s: s,
	}
}

func (grpc *GRPCSenderServer) SendMail(ctx context.Context, r *sender.SendMailRequest) (*sender.SendMailResponse, error) {
	// TODO: изменить Err в SendResponse на что-то другое или добавить message и code
	// TODO: реализовать получение template по template_id из grpc и вставку variable в шаблон + валидацию этих значений!
	// d := dto.NewEmailDTO(grpc.c.Mail().GetAddressFrom(), r.To, r.Subject, "test html template")
	vm := map[string]interface{}{"message": "test html template"}

	d := dto.NewEmailDTO(r.To, r.Subject, vm)

	sent, err := grpc.s.Send(&d)
	if err != nil {
		grpc.l.Instance().Error("Mail isn't sent!", err)
		return &sender.SendMailResponse{
			Sent:  sent,
			Error: "Mail isn't sent!",
		}, err
	}
	// TODO: r.To[0] - изменить на работу с слайсом
	grpc.l.Instance().Info(fmt.Sprintf("Mail to %s sent!", r.To[0]))
	return &sender.SendMailResponse{
		Sent:  true,
		Error: "",
	}, nil
}
