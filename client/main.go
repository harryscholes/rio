package main

import (
	"context"
	"os"
	"time"

	"github.com/harryscholes/rio/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("did not connect")
	}
	defer conn.Close()

	c := api.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(
		ctx,
		&api.HelloRequest{
			Name: name,
		},
	)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("could not greet")
	}

	msg := r.GetMessage()
	log.Info().
		Str("message", msg).
		Msgf("Greeting: %s", msg)
}
