package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/babaunba/project-management/api-gateway/gen/proto/labels/v1"
	"github.com/babaunba/project-management/api-gateway/internal/server"
)

const (
	grpcAddr = ":42043"
	httpAddr = ":42042"
)

func main() {
	protoConverter := converter.NewProtoJSONPayloadConverter()
	converter.NewProtoJSONPayloadConverter()
	converter := converter.NewCompositeDataConverter(protoConverter)

	srv := grpc.NewServer()
	{
		s, err := server.New(client.Options{
			HostPort:      "45.155.205.163:7233",
			DataConverter: converter,
		})
		if err != nil {
			log.Fatalf("failed to create server: %v", err)
		}
		labels.RegisterLabelsServer(srv, s)
	}

	reflection.Register(srv)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		err := labels.RegisterLabelsHandlerFromEndpoint(
			ctx,
			mux,
			grpcAddr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			log.Fatalf("failed to start HTTP gateway: %v", err)
		}

		log.Fatal(http.ListenAndServe(httpAddr, mux))
	}()

	log.Fatalf("failed to serve: %v", srv.Serve(lis))
}
