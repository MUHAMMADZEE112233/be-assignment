package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

// func GenerateToken(userID uint) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		UserID: userID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtKey)
// }

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
