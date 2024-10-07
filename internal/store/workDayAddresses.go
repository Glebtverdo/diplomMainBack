package store

import (
	"context"
	"diplomMainBack/internal/models"
	plugBackGrpc "diplomMainBack/internal/store/grpc/gen"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetWorkDayAddresses(tokenPair models.JwtTokenPair) ([]models.WorkDayAddress, error) {
	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", tokenPair.Access)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := grpcClient.GetWorkDayAddresses(ctx, &plugBackGrpc.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	result := []models.WorkDayAddress{}
	for _, req := range res.Addresses {
		result = append(result, models.WorkDayAddress{
			Id:       int(req.Id),
			Name:     req.Name,
			Coords:   req.Coords,
			IsActive: req.IsActive,
		})
	}
	return result, nil
}
