package middleware

import (
	"articleproject/error"
	"articleproject/utils"
	"context"
	"net/http"
)

func VerifyAccessToken(flag int) func(handler http.Handler) http.Handler {
	//flag == 1 means check for admin, otherwise only verifies access token.
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			token = token[7:]
			id, isadmin, err := utils.VerifyToken(token)
			if err != nil {
				utils.ErrorGenerator(w, errorhandling.AccessTokenExpired)
				return
			}
			if flag == 1 && !isadmin {
				utils.ErrorGenerator(w, errorhandling.UnauthorizedError)
				return
			}
			ctx := context.WithValue(r.Context(), "id", id)
			ctx = context.WithValue(ctx, "isadmin", isadmin)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RetrieveRefreshToken() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			token = token[7:]
			id, isadmin, err := utils.VerifyToken(token)
			if err != nil {
				utils.ErrorGenerator(w, errorhandling.RefreshTokenExpired)
				return
			}
			ctx := context.WithValue(r.Context(), "token", token)
			ctx = context.WithValue(ctx, "id", id)
			ctx = context.WithValue(ctx, "isadmin", isadmin)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
