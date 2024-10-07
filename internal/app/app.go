package app

import (
	"diplomMainBack/internal/store"
	grpcHandlers "diplomMainBack/internal/transport/grpc/handlers"
	restTransport "diplomMainBack/internal/transport/http"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func InitApp() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return
	}
	err := store.Init()
	if err != nil {
		log.Fatalf("can not connect to server %s", err.Error())
		return
	}
	go restTransport.InitServer()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcHandlers.InitServer()
	}()

	wg.Wait()
}
