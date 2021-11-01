package user

import (
	"context"
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	us *UserService
}

func NewAuthMiddleware(us *UserService) *AuthMiddleware {
	return &AuthMiddleware{us: us}
}

func (am *AuthMiddleware) CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u *User

		// get user by id from decoded jwt cookie
		if cookie, err := r.Cookie("access_token"); err != nil {
			fmt.Println(err.Error())
		} else {
			claims, _ := ParseToken(cookie)
			u = am.us.GetOne(claims.Id)
		}

		ctx := context.WithValue(r.Context(), UserContext, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (am *AuthMiddleware) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if u := r.Context().Value(UserContext).(*User); u == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
