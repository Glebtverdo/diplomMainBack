package restTransport

import (
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getWorkDayAddresses(w http.ResponseWriter, r *http.Request) {
	tokenPair := r.Context().Value(models.KeyForAuthorizationTokens).(models.JwtTokenPair)
	workDayAddresses, err := services.GetWorkDayAddresses(models.JwtTokenPair{
		Access:  tokenPair.Access,
		Refresh: tokenPair.Refresh,
	})
	if err != nil {
		errorHandler(err, w)
		return
	}
	str, err := json.Marshal(workDayAddresses)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte(str))
}

func initWorkDayAddressesRouter(router *mux.Router) {
	router.HandleFunc("/get_work_day_addresses", getWorkDayAddresses).Methods("GET", "OPTIONS")
}
