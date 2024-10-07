package restTransport

import (
	midlewares "diplomMainBack/internal/transport/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitServer() {
	router := mux.NewRouter()
	router.Schemes("https")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow world"))
	})
	initLoginRouter(router)
	initRoutsRouter(router)
	initWorkDayAddressesRouter(router)
	router.Use(midlewares.LoggingMiddleware)
	router.Use(midlewares.RecoveryMiddleware)
	router.Use(midlewares.GlobalHeadersMiddleware)
	router.Use(midlewares.CheckAuthorization)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins for demonstration purposes
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // Include Authorization header
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":4000", handler))
}
