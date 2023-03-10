package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func AssignService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("BWAFUNDING")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	val, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return val, err
	}

	return val, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	val, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return val, err
	}

	return val, nil
}
