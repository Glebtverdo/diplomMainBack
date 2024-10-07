package grpcHandlers

import (
	"context"
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/services"
	mainBack "diplomMainBack/internal/transport/grpc/gen"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Login(
	ctx context.Context,
	in *mainBack.LoginRequest,
) (*mainBack.LoginResponse, error) {
	loginPair := models.LoginPair{
		Login:    in.Login,
		Password: in.Password,
	}
	if loginPair.Login == "" || loginPair.Password == "" {
		return nil, fmt.Errorf("invalid data")
	}
	jwtPair, err := services.Login(loginPair)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &mainBack.LoginResponse{Access: jwtPair.Access, Refresh: jwtPair.Refresh}, nil
}
