package jwt

import (
	"ala-coffee-notification/configs"
	"ala-coffee-notification/utils/errs"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtKey = []byte(configs.Env.JWT.SecretKey)
)

func SignToken(id string) (string, error) {
	fmt.Println(id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(configs.Env.JWT.ExpiresTime) * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   id,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(signedToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwt.RegisteredClaims{},
		func(_ *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	fmt.Println(token)

	if err != nil {
		fmt.Println(err.Error())
		// if err.Error() == fmt.Sprintf("%s: %s", jwt.ErrTokenInvalidClaims.Error(), jwt.ErrTokenExpired.Error()) {
		return nil, errs.E(http.StatusUnauthorized, strings.TrimSpace(strings.Split(err.Error(), ":")[1]))
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errs.E(http.StatusBadRequest, "couldn't parse claims")
	}

	return claims, nil
}

func GetPayload(r *http.Request) *jwt.RegisteredClaims {
	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok {
		return nil
	}

	return claims
}
