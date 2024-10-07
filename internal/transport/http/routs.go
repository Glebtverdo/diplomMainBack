package restTransport

import (
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getRouts(w http.ResponseWriter, r *http.Request) {
	tokenPair := r.Context().Value(models.KeyForAuthorizationTokens).(models.JwtTokenPair)
	queryParams := r.URL.Query()
	DestinationCoords := queryParams.Get("destination_coords")
	PointsCount, err := strconv.Atoi(queryParams.Get("points_count"))
	if PointsCount == 0 || err != nil {
		errorHandler(fmt.Errorf("error in points_count"), w)
		return
	}
	routes, err := services.GetRouts(models.GetRoutesModel{
		Access:            tokenPair.Access,
		Refresh:           tokenPair.Refresh,
		DestinationCoords: DestinationCoords,
		PointsCount:       PointsCount,
	})
	if err != nil {
		errorHandler(err, w)
		return
	}
	res := models.GetRoutesResponce{
		Count:  len(routes),
		Routes: routes,
	}
	str, err := json.Marshal(res)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte(str))
}

func getRequesCount(w http.ResponseWriter, r *http.Request) {
	tokenPair := r.Context().Value(models.KeyForAuthorizationTokens).(models.JwtTokenPair)
	requesCount, err := services.GetRequesCount(models.JwtTokenPair{
		Access:  tokenPair.Access,
		Refresh: tokenPair.Refresh,
	})
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte(strconv.Itoa(requesCount)))
}

func initRoutsRouter(router *mux.Router) {
	router.HandleFunc("/get_routs", getRouts).Methods("GET", "OPTIONS")
	router.HandleFunc("/get_reques_count", getRequesCount).Methods("GET", "OPTIONS")
}
