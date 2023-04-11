package library

import (
	"airportal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretAdmin = []byte(config.JwtSecretAdmin)

type UserClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(ID int) (string, error) {
	claims := UserClaims{
		ID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "share-server",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretAdmin)
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretAdmin), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
