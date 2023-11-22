package auth

import (
	"context"

	auth_service_v1 "github.com/markraiter/protos/gen/go/auth_service"
	"google.golang.org/grpc"
)

type serverAPI struct {
	auth_service_v1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	auth_service_v1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *auth_service_v1.LoginRequest) (*auth_service_v1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, req *auth_service_v1.RegisterRequest) (*auth_service_v1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *auth_service_v1.IsAdminRequest) (*auth_service_v1.IsAdminResponse, error) {
	panic("implement me")
}
