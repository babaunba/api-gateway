package server

import (
	"go.temporal.io/sdk/client"

	"github.com/babaunba/project-management/api-gateway/gen/proto/labels/v1"
)

// Server is a gRPC server implementation that requires workflow domain definition
type Server struct {
	labels.UnimplementedLabelsServer
	client client.Client // temporal client
}

// New is a *Server constructor
func New(opts client.Options) (srv *Server, err error) {
	client, err := client.Dial(opts)
	srv = &Server{client: client}
	return
}
