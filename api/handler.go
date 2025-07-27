package api

import (
	"context"
	gen "hello-api/gen"
)

type Server struct {
	gen.UnimplementedHandler
}

func (s *Server) GetHello(ctx context.Context) (*gen.GetHelloOK, error) {
	return &gen.GetHelloOK{
		Message: "Hello, world!",
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *gen.RegisterUserReq) (r *gen.RegisterUserOK, _ error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return &gen.RegisterUserOK{Message: "All good"}, nil
}
