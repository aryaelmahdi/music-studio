package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type JWTInterface interface {
	GenerateJWT(userID string, role string) map[string]any
	GenerateToken(id string, role string) string
	ExtractToken(token *jwt.Token) (any, any)
}

type JWT struct {
	signKey    string
	refreshKey string
}

func NewJWT(signKey string) JWTInterface {
	return &JWT{
		signKey: signKey,
	}
}

func (j *JWT) GenerateJWT(userID string, role string) map[string]any {
	var result = map[string]any{}
	var accessToken = j.GenerateToken(userID, role)
	if accessToken == "" {
		return nil
	}
	result["access_token"] = accessToken
	return result
}

func (j *JWT) GenerateToken(id string, role string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j JWT) RefereshJWT(accessToken string, refreshToken *jwt.Token) map[string]any {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		logrus.Error("get token expiration error", err.Error())
		return nil
	}
	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = jwt.MapClaims{}

		newToken, err := jwt.ParseWithClaims(accessToken, newClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.signKey), nil
		})

		if err != nil {
			log.Error(err.Error())
			return nil
		}

		newClaim = newToken.Claims.(jwt.MapClaims)
		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newToken.Raw
		result["refresh_token"] = newSignedRefreshToken
		return result
	}

	return nil
}

func (j *JWT) generateRefreshToken(accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(j.refreshKey))

	if err != nil {
		return ""
	}

	return refreshToken
}

func (j JWT) ExtractToken(token *jwt.Token) (any, any) {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		fmt.Println(expTime.Time.Compare(time.Now()))
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			var id = mapClaim["id"]
			var role = mapClaim["role"]
			return id, role
		}

		logrus.Error("Token expired")
		return nil, nil

	}
	return nil, nil
}
