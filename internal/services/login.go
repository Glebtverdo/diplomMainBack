package services

import (
	"diplomMainBack/internal/models"
	"diplomMainBack/internal/store"
)

func Login(loginPair models.LoginPair) (models.JwtTokenPair, error) {
	return store.Login(loginPair)
}
