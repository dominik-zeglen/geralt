package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
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
		ctx := r.Context()
		span, _ := opentracing.StartSpanFromContext(
			r.Context(),
			"middleware-jwt",
		)

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
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if claims, valid := token.Claims.(*UserClaims); valid && token.Valid {
				span.LogFields(
					log.String("authorized-as", claims.ID),
				)
				ctx = context.WithValue(
					ctx,
					userIDContextKey,
					claims.ID,
				)
			}
		}

		span.Finish()
		next(w, r.WithContext(ctx))
		return
	}
}
