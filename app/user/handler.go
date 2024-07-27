package user

import context "context"

type UserHandler struct {
	UnimplementedHelloServiceServer
	userService userServicer
}

func NewHandler(
	userService userServicer,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (s *UserHandler) SayHello(ctx context.Context, req *UserReq) (*UserRes, error) {
	res, err := s.userService.SayHello(ctx, UserRequest{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &UserRes{
		Message: res,
	}, nil
}
