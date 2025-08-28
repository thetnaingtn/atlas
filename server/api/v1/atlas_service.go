package v1

import (
	"context"
	"time"

	apiv1 "atlas/proto/gen/api/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Health implements the Health RPC method for APIV1Service
func (s *APIV1Service) Health(ctx context.Context, req *apiv1.HealthRequest) (*apiv1.HealthResponse, error) {
	return &apiv1.HealthResponse{
		Status:    "ok",
		Timestamp: timestamppb.New(time.Now()),
	}, nil
}
