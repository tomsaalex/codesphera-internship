package middleware

import (
	"context"
	"curs1_boilerplate/util"
	"net/http"
)

type jwtDataKey string

const userEmailKey = jwtDataKey("userEmail")

func AttachUser(jwtHelper util.JwtUtil) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authCookie, err := r.Cookie("authCookie")
			newCtx := r.Context()
			if err == nil {
				userEmail, err := jwtHelper.ParseJWT(authCookie.Value)
				if err == nil {
					newCtx = SetUserEmailToContext(r.Context(), userEmail)
				}
			}
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetUserEmailFromContext(r.Context()) == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireGuest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetUserEmailFromContext(r.Context()) != "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetUserEmailToContext(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, userEmailKey, userEmail)
}

func GetUserEmailFromContext(ctx context.Context) string {
	userEmail, conversionValid := ctx.Value(userEmailKey).(string)
	if conversionValid {
		return userEmail
	}
	return ""
}
