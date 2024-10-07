package grpcHandlers

import (
	mainBack "diplomMainBack/internal/transport/grpc/gen"
	midlewares "diplomMainBack/internal/transport/middlewares"

	"log"
	"net"

	"google.golang.org/grpc"
)

type serverAPI struct {
	mainBack.UnimplementedMainBackServer
}

func register(gRPCServer *grpc.Server) {
	mainBack.RegisterMainBackServer(gRPCServer, &serverAPI{})
}

func InitServer() {
	lis, err := net.Listen("tcp", ":400")
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(midlewares.CheckAuthorizationStreamInterceptor),
		grpc.ChainUnaryInterceptor(midlewares.CheckAuthorizationUnaryInterceptor),
	)
	register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
