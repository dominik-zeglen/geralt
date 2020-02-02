package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

const userIDContextKey = "userID"

// UserClaims holds all token data
type UserClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (api *API) withJwt(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerContent := r.Header.Get("Authorization")
		if headerContent != "" && headerContent != "null" {
			tokenString := strings.Split(headerContent, " ")[1]
			token, err := jwt.ParseWithClaims(
				tokenString,
				&UserClaims{},
				func(token *jwt.Token) (interface{}, error) {
					if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
						return nil, errors.New("Invalid signing method")
					}
					return []byte(api.conf.secret), nil
				},
			)

			if err != nil {
				next(w, r)
				return
			}

			if claims, valid := token.Claims.(*UserClaims); valid && token.Valid {
				ctx := context.WithValue(
					r.Context(),
					userIDContextKey,
					claims.ID,
				)

				next(w, r.WithContext(ctx))
			} else {
				next(w, r)
			}
		} else {
			next(w, r)
		}
	}
}
