package authentication

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userID,
			"exp":      time.Now().Add(expiresIn).Unix(),
		})

	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateJWT(tokenString, tokenSecret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
