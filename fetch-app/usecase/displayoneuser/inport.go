package displayoneuser

import (
	"context"
)

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Claims map[string]interface{}
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	User map[string]interface{}
}
