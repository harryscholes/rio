package main

import (
	"context"
	"fmt"
	"net"

	"contrib.go.opencensus.io/exporter/zipkin"
	"github.com/harryscholes/rio/api"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	address = "localhost:50051"
)

type server struct{}

func main() {
	// Configure OpenConcensus exporter to export traces to Zipkin
	localEndpoint, err := openzipkin.NewEndpoint("greeter", "192.168.1.5:5454")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create the local zipkinEndpoint")
	}
	reporter := zipkinHTTP.NewReporter("http://localhost:9411/api/v2/spans")
	ze := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(ze)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

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
	_, span := trace.StartSpan(ctx, "SayHello")
	defer span.End()

	// FIXME Annotations do not show up in the trace
	span.Annotate(
		[]trace.Attribute{
			trace.StringAttribute("name", req.Name),
		},
		"Request name",
	)

	return &api.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
