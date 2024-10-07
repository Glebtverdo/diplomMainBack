package store

import (
	"context"
	"diplomMainBack/internal/models"
	plugBackGrpc "diplomMainBack/internal/store/grpc/gen"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetRequests(AccessToken string) ([]models.Request, error) {
	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", AccessToken)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := grpcClient.GetRequests(ctx, &plugBackGrpc.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	var requests []models.Request
	for _, req := range res.Data {
		var request models.Request
		request.Id = int(req.Id)
		request.Object.Address = req.Object.Address
		request.Object.Coords[0] = req.Object.Coords[0]
		request.Object.Coords[1] = req.Object.Coords[1]
		request.Object.Name = req.Object.Name
		requests = append(requests, request)
	}
	return requests, nil
}
