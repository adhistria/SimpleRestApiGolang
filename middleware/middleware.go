package middleware

import (
	"context"
	"fmt"
	"net/http"
	"omahihrom/respond"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("omahihromjantisupersecretboy"), nil
			})
			if err != nil {
				respond.RespondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			if token.Valid {
				ctx := context.WithValue(req.Context(), "userInfo", token.Claims)
				req = req.WithContext(ctx)
				next(w, req)
			} else {
				respond.RespondWithError(w, http.StatusBadRequest, "Invalid authorization token")
			}
		} else {
			respond.RespondWithError(w, http.StatusBadRequest, "An authorization header is required")
		}
	})
}
