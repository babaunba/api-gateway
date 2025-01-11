package domain

import (
	"go.temporal.io/sdk/workflow"

	"github.com/babaunba/project-management/api-gateway/gen/proto/labels/v1"
)

// GetLabelsWF is not implemented
// TODO: Implement once ML service is finished
func (domain *Domain) GetLabelsWF(
	workflow.Context,
	*labels.GetLabelsRequest,
) (resp *labels.GetLabelsResponse, err error) {
	resp = &labels.GetLabelsResponse{Labels: []string{"bug", "shit"}}
	return
}
