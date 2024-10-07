package models

import "github.com/golang-jwt/jwt"

type LoginPair struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtTokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type JwtClaims struct {
	Type      string
	TokenPair JwtTokenPair
	jwt.StandardClaims
}

type MyString string

var KeyForAuthorizationTokens MyString = "tokenPair"

var AuthorizationKey MyString = "authorization"
