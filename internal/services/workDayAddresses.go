package services

import (
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/store"
)

func GetWorkDayAddresses(tokenPair models.JwtTokenPair) ([]models.WorkDayAddress, error) {
	return store.GetWorkDayAddresses(tokenPair)
}
