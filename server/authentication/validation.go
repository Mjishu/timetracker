package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("placeholder-key")

// func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username": userID,
// 			"exp":      time.Now().Add(expiresIn).Unix(),
// 		})

// 	tokenString, err := token.SignedString(tokenSecret)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil

// }

// func ValidateJWT(tokenString, tokenSecret string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return tokenSecret, nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}
// 	return nil
// }

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "timetracker",
		"aud": "default",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func AuthenticateMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token") //! again this might not work because the cookies will be on the frontend not on the go side of things?
	if err != nil {
		fmt.Println("Token missing in cookie")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not find the token", "success": false})
		c.Abort()
		return
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed %v\\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "token verification failed", "success": false})
		c.Abort()
		return
	}
	fmt.Printf("token verified successfully, claims %+v\\n", token.Claims)
	c.Next()
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
