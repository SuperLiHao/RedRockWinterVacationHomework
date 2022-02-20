package Utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}

const TokenTime = time.Hour * 1
const RefreshTokenTime = time.Hour * 24 * 30

var Secret = []byte("RedRockBBS Author:Hao_pp")

func GetToken(username string) (string, error) {

	c := MyClaims{

		username, "AccessToken", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			Issuer:    "RedRockBBS",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(Secret)

}

func GetRefreshToken(username string) (string, error) {

	c := MyClaims{

		username, "RefreshToken", jwt.StandardClaims{
			ExpiresAt: 0,
			Issuer:    "RedRockBBS",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(Secret)

}

func PraseToken(tokenStr string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("PraseToken Error")

}

func JudgeAccessToken(msg string) (string, bool) {

	claims, err := PraseToken(msg)

	if err != nil {
		return "NULL", false
	}

	if claims.StandardClaims.ExpiresAt <= time.Now().Unix() || claims.StandardClaims.Issuer != "RedRockBBS" || claims.TokenType != "AccessToken" {
		return "NULL", false
	}

	return claims.Username, true
}
