package user

import context "context"

type userServicer interface {
	SayHello(ctx context.Context, userRequest UserRequest) (string, error)
}
