package main

import (
	"context"
	"encoding/json"
	"github.com/PraveenKusuluri08/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type Exception struct {
	Message string
}

func endPonint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")
		err := json.NewEncoder(w).Encode(r)
		if err != nil {
			return
		}
		header = strings.TrimSpace(header)
		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			if err := json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"}); err != nil {
				return
			}
			return
		}
		tokenModel := &model.Token{}
		_, err = jwt.ParseWithClaims(header, tokenModel, func(token *jwt.Token) (interface{}, error) {

			return []byte("secret"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			err := json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			if err != nil {
				return
			}
			return
		}
		ctx := context.WithValue(r.Context(), "user", tokenModel)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
