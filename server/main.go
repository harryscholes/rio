package main

import (
	"context"
	"fmt"
	"net"

	"github.com/harryscholes/rio/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type server struct{}

func main() {
	// gRPC
	log.Debug().Msg("starting tcp listener")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to listen")
	}
	log.Debug().Msg("tcp listener started")

	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &server{})
	log.Debug().Msg("starting gRPC server")
	if err := s.Serve(lis); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to serve")
	}
}

func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
