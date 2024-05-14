package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("golangstarter_s3cr3T_k3Y") // storing as ascii value.

func Newservice() *jwtService {

	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {

	cliam := jwt.MapClaims{} // This allows you to manipulate the claims as a map.
	cliam["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliam) // for generate the token

	signedToken, err := token.SignedString(SECRET_KEY) // converting token as  string

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidToken(encodedToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {

			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {

		return token, err
	}

	return token, nil
}
