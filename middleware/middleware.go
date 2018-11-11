package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api/model"

	// "strings"

    jwt "github.com/dgrijalva/jwt-go"
    // "github.com/gorilla/mux"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("validate middleware")
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
				json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
				return
			}
			if token.Valid {
                // ctx := context.WithValue(context.Background(), "userInfo", token.Claims)
                ctx := context.WithValue(req.Context(), "userInfo", token.Claims)
                // context.Set(req, "decoded", token.Claims)
                // next(w, req)

				req = req.WithContext(ctx)
				next(w, req)
			} else {
				json.NewEncoder(w).Encode(model.Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: "An authorization header is required"})
		}
	})
}
