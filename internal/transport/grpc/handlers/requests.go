package grpcHandlers

import (
	"context"
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/services"
	mainBack "diplomMainBack/internal/transport/grpc/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetRouts(
	ctx context.Context,
	in *mainBack.EmptyRequest,
) (*mainBack.RequestsCount, error) {
	tokenPair := ctx.Value(models.KeyForAuthorizationTokens).(models.JwtTokenPair)
	_, err := services.GetRouts(models.GetRoutesModel{
		Access:  tokenPair.Access,
		Refresh: tokenPair.Refresh,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &mainBack.RequestsCount{Count: int32(0)}, nil
}
