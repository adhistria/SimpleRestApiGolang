package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api/model"
	// "strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    fmt.Println("validate middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// arr_string_err = arr_string_err[:0]
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
            fmt.Println("header not nil")
            // bearerToken := strings.Split(authorizationHeader, " ")
            // fmt.Println(authorizationHeader)
			// if len(bearerToken) == 2 {
            fmt.Println("bearer")
            token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
                fmt.Println("jwt pares")
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return []byte("secret"), nil
            })
            if error != nil {
                // arr_string_err = append(arr_string_err, error.Error())
                // respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
                json.NewEncoder(w).Encode(model.Exception{Message: error.Error()})
                return
            }
            if token.Valid {
                ctx := context.WithValue(context.Background(), "userInfo", token.Claims)
                req = req.WithContext(ctx)
                next(w, req)
            } else {
                json.NewEncoder(w).Encode(model.Exception{Message: "Invalid authorization token"})
            }
            // }
            // fmt.Println("bearer nil")
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: "An authorization header is required"})
		}
	})
}
