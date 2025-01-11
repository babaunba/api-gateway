package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/babaunba/project-management/api-gateway/gen/proto/labels/v1"
)

const (
	workflowName = "get-labels-workflow"
	taskQueue    = "labels-tasks"
)

// GetLabels adds a labels generation task to the queue and expects a worker to complete it
func (s *Server) GetLabels(
	ctx context.Context,
	req *labels.GetLabelsRequest,
) (resp *labels.GetLabelsResponse, err error) {
	opts := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s:%s", workflowName, uuid.New().String()),
		TaskQueue: taskQueue,
	}

	we, err := s.client.ExecuteWorkflow(ctx, opts, workflowName, req)
	if err != nil {
		err = fmt.Errorf("unable to execute workflow: %w", err) // first wrap and add a message
		err = status.Error(codes.Internal, err.Error())         // then use a gRPC status wrapper
		return
	}

	err = we.Get(ctx, &resp)
	if err != nil {
		err = fmt.Errorf("unable to get workflow result: %w", err) // analogous
		err = status.Error(codes.Internal, err.Error())
		return
	}

	return
}
