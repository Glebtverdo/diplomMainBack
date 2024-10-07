package restTransport

import (
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	var user models.LoginPair
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	tokenPair, err := services.Login(user)
	if err != nil {
		errorHandler(err, w)
		return
	}
	str, err := json.Marshal(tokenPair)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write(str)
}

func initLoginRouter(router *mux.Router) {
	router.HandleFunc("/login", login).Methods("POST", "OPTIONS")
}
