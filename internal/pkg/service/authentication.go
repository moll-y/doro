package service

import (
	"github.com/golang-jwt/jwt/v5"
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
	log.Println(user)
	if user.Password != password {
		log.Println("password incorrect")
		return "", err
	}
	// Create a new token object, specifying signing method and the claims you
	// would like it to contain.
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})
	// Sign and get the complete encoded token as a string using the secret
	ts, err := t.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return ts, nil
}
