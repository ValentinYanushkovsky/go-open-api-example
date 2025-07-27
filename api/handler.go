package api

import (
	"context"
	gen "hello-api/gen"
)

type Server struct {
	gen.UnimplementedHandler
}

func (s *Server) GetHello(ctx context.Context) (r *gen.GetHelloOK, _ error) {
	return &gen.GetHelloOK{
		Message: "Hello, world!",
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *gen.RegisterUserReq) (r gen.RegisterUserRes, _ error) {
	//err := req.Validate()
	//if err != nil {
	return &gen.RegisterUserForbidden{
		//Message: err.Error(),
		Message: "You do not have permission to do this!",
		Code:    4001,
		Details: gen.OptString{Value: "some details", Set: true},
	}, nil
	//}

	//	return &gen.RegisterUserOK{Message: "All good"}, nil
}
