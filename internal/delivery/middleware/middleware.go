package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/delivery/rest"
	"github.com/soerjadi/wwalet/internal/usecase/user"
)

func OnlyLoggedInUser(userManagement user.Usecase, cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from header
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				unauthorizedResp(w, r)
				return
			}
			if !strings.Contains(tokenString, "Bearer") {
				unauthorizedResp(w, r)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			// Parse token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("signing method is invalid")
				} else if method != jwt.SigningMethodHS256 {
					return nil, errors.New("signing method is invalid")
				}

				return []byte(cfg.Server.SecretKey), nil
			})
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok && !token.Valid {
				unauthorizedResp(w, r)
				return
			}

			userID, ok := claims["id"].(string)
			if !ok {
				unauthorizedResp(w, r)
				return
			}

			user, err := userManagement.GetByID(context.Background(), userID)
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			reqMap := make(map[string]interface{})
			reqMap["user"] = user
			reqMap["req"] = r

			r = appendToContext(r.Context(), reqMap)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorizedResp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := rest.Response{
		Message: "Unauthenticated",
	}

	w.WriteHeader(http.StatusUnauthorized)
	x, _ := json.Marshal(resp)
	w.Write(x)
}

func appendToContext(ctx context.Context, reqMap map[string]interface{}) *http.Request {
	userID := reqMap["userID"]
	r := reqMap["req"].(*http.Request)
	ctx = context.WithValue(ctx, "user-key-respondent", userID)

	r = r.WithContext(ctx)
	return r
}
