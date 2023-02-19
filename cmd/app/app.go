package app

import (
	"log"
	"net"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/gen/sender"
	"github.com/cheeeasy2501/go-email-sender/internal/transport/grpc/v1"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	grpcServer := grpc.NewServer()
	grpcSenderSevice := &v1.GRPCSenderServer{}
	sender.RegisterSenderServer(grpcServer, grpcSenderSevice)

	l, err := net.Listen("tcp", cfg.GRPC().GetListenerAddr())
	if err != nil {
		return err
	}

	log.Println("Start GRPC server")
	go func() {
		if err := grpcServer.Serve(l); err != nil {
			log.Println("GRPC IS NOT STARTED!")
			panic(err)
		}
	}()

	log.Println("GRPC IS STARTED!")

	return nil
}
