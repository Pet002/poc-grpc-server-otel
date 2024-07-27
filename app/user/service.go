package user

import (
	"context"
	"fmt"

	"github.com/Pet002/poc-grpc-server-otel/logger"
)

type UserService struct {
}

func NewService() *UserService {
	return &UserService{}
}

func (s UserService) SayHello(ctx context.Context, userRequest UserRequest) (string, error) {

	if userRequest.Name != "petch" {
		return "", fmt.Errorf("error input: %q", userRequest.Name)
	}
	logger.Info(ctx, "test log")
	return "hello petch", nil
}
