package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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
