package service

import (
	"context"

	v1 "load-book/api/load-book/v1"
	"load-book/internal/service/health/biz"
)

// Service is a health service.
type Service struct {
	v1.UnimplementedHealthServer

	uc *biz.UseCase
}

// NewService new a health service.
func NewService(uc *biz.UseCase) v1.HealthServer {
	return &Service{uc: uc}
}

func (s *Service) Get(ctx context.Context, _ *v1.GetRequest) (*v1.GetReply, error) {
	return &v1.GetReply{Status: "Success"}, nil
}
