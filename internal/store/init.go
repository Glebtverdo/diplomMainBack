package store

import (
	plugBackGrpc "diplomMainBack/internal/store/grpc/gen"
	"net/http"

	"google.golang.org/grpc"
)

var grpcClient plugBackGrpc.PlugBackGrpcClient
var httpClient *http.Client

func Init() error {
	conn, err := grpc.Dial("localhost:2000", grpc.WithInsecure())
	if err != nil {
		return err
	}
	// defer conn.Close()
	grpcClient = plugBackGrpc.NewPlugBackGrpcClient(conn)
	httpClient = &http.Client{}
	return nil
}
