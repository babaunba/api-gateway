package server

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"

	"github.com/babaunba/project-management/api-gateway/gen/proto/labels/v1"
)

// there is no reason to pass req/resp as pointers on account of temporal deserialising the data
type domain interface {
	GetLabelsWF(workflow.Context, *labels.GetLabelsRequest) (*labels.GetLabelsResponse, error)
}

// Server is a gRPC server implementation that requires workflow domain definition
type Server struct {
	labels.UnimplementedLabelsServer
	domain domain
	client client.Client // temporal client
}

// New is a *Server constructor
func New(domain domain, opts client.Options) (srv *Server, err error) {
	client, err := client.Dial(opts)
	srv = &Server{domain: domain, client: client}
	return
}
