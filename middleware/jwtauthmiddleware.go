package middleware

import (
	"context"
	"curs1_boilerplate/util"
	"fmt"
	"net/http"
)

func JwtAuth(jwtHelper util.JwtUtil) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authCookie, err := r.Cookie("authCookie")
			if err != nil {
				jwtAuthFailed(w, r)
				fmt.Println("Failed big time")
				return
			}

			userEmail, err := jwtHelper.ParseJWT(authCookie.Value)
			if err != nil {
				jwtAuthFailed(w, r)
				fmt.Println(err.Error())
				return
			}

			// TODO: Use custom type for value
			newCtx := context.WithValue(r.Context(), "userEmail", userEmail)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func jwtAuthFailed(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
