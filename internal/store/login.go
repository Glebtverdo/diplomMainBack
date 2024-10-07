package store

import (
	"context"
	"diplomMainBack/internal/models"
	plugBackGrpc "diplomMainBack/internal/store/grpc/gen"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(LoginPair models.LoginPair) (models.JwtTokenPair, error) {
	res, err := grpcClient.Login(context.Background(),
		&plugBackGrpc.LoginRequest{Login: LoginPair.Login, Password: LoginPair.Password})
	if err != nil {
		return models.JwtTokenPair{}, err
	}
	accessPayload := models.JwtClaims{
		Type: "access",
		TokenPair: models.JwtTokenPair{
			Access:  res.Access,
			Refresh: res.Refresh,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessPayload)
	access, e := token.SignedString([]byte("secret word"))
	if e != nil {
		return models.JwtTokenPair{}, e
	}
	refreshPayload := models.JwtClaims{
		Type: "refresh",
		TokenPair: models.JwtTokenPair{
			Access:  res.Access,
			Refresh: res.Refresh,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)
	refresh, e := token.SignedString([]byte("secret word"))
	if e != nil {
		return models.JwtTokenPair{}, e
	}
	var tokens = models.JwtTokenPair{
		Access:  access,
		Refresh: refresh,
	}
	return tokens, nil
}
