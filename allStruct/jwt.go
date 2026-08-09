package allstruct

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("laxman") // 👈 Use env variable in real apps

type Claims struct {
	UserID string `json:"userId"` //json userId
	jwt.RegisteredClaims
}

func GenerateJWT(userID string) (string, error) {
	// Create the claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		}, // expires in 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// i am pasing name and email to create the jwt token after decoding it will retun email i need to make available for id
func ValidateJWT(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("hello1")
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("hello3")

			return "", fmt.Errorf(("invalid token signature"))
		}
	}
	if !token.Valid {
		fmt.Println("hello3")

		return "", fmt.Errorf("invalid token")
	}
	fmt.Println("in jwt.go validate fun----", claims.UserID)
	return claims.UserID, nil

}
