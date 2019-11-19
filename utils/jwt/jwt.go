package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("VophanEngine")

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "masterportal",
		},
		Username:username,
		Password:password,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token,err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret,nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok&&tokenClaims.Valid {
			return claims, nil
		}
 	}
	return nil, err
}
