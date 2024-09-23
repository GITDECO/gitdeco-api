package pkg

import (
	"fmt"
	"gitdeco-api/internal/exception"
	"gitdeco-api/tools"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var AccessTokenExpirationTime = time.Hour * 1
var RefreshTokenExpirationTime = time.Hour * 24 * 31
var TestTokenExpirationTime = time.Hour * 24 * 365

func GenerateToken(username string, test bool) *Token {
	token := new(Token)
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expiration": time.Now().Add(tools.Ternary(test, TestTokenExpirationTime, AccessTokenExpirationTime).(time.Duration)).Unix(),
	})
	accessToken, jwtError := access.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	if jwtError != nil {
		panic(&exception.Error{Key: "TOKEN_GENERATE_ERROR", Data: "Access Token"})
	}
	token.AccessToken = accessToken

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expiration": time.Now().Add(tools.Ternary(test, TestTokenExpirationTime, RefreshTokenExpirationTime).(time.Duration)).Unix(),
	})
	refreshToken, jwtError := refresh.SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))
	if jwtError != nil {
		panic(&exception.Error{Key: "TOKEN_GENERATE_ERROR", Data: "Refresh Token"})
	}
	token.RefreshToken = refreshToken

	return token
}

func RefreshToken(refreshToken string) *Token {
	username, expiration := refreshTokenParseX(refreshToken)
	if float64(time.Now().Unix()) > expiration {
		panic(&exception.Error{Key: "TOKEN_VALID_ERROR", Data: "Refresh Token"})
	}

	return GenerateToken(username, false)
}

func ValidateToken(accessToken string) string {
	username, expiration := accessTokenParseX(accessToken)
	if float64(time.Now().Unix()) > expiration {
		panic(&exception.Error{Key: "TOKEN_VALID_ERROR", Data: "Access Token"})
	}

	return username
}

func accessTokenParseX(accessToken string) (string, float64) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})
	if err != nil {
		panic(&exception.Error{Key: "TOKEN_PARSE_ERROR", Data: "Access Token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["username"].(string), claims["expiration"].(float64)
	}

	return "", 0
}

func refreshTokenParseX(refreshToken string) (string, float64) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})
	if err != nil {
		panic(&exception.Error{Key: "TOKEN_PARSE_ERROR", Data: "Refresh Token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["username"].(string), claims["expiration"].(float64)
	}

	return "", 0
}
