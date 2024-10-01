package service

import (
	"errors"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"log"
	"moll-y.io/doro/internal/pkg/domain"
)

type AuthenticationService struct {
	UserRepository domain.UserRepository
}

func (as *AuthenticationService) Authenticate(email, password string) (string, error) {
	user, err := as.UserRepository.FindUserByEmail(email)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if user == nil || user.Password != password {
		return "", errors.New("Email or password incorrect.")
	}
	t := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims{
		"actor": user.ID,
	})
	s, err := t.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return s, nil
}

func (as *AuthenticationService) Parse(jwt string) (int, error) {
	var claims struct {
		Actor int `json:"actor"`
		jwtv5.RegisteredClaims
	}
	t, err := jwtv5.ParseWithClaims(jwt, &claims, func(token *jwtv5.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, nil
	}
	if !t.Valid {
		return 0, errors.New("JWT is invalid.")
	}
	return claims.Actor, nil
}
