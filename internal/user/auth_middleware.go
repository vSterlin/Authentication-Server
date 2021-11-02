package user

import (
	"context"
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
		// errors don't matter since if they are present
		// will pass nil user to next handler
		if cookie, _ := r.Cookie("access_token"); cookie != nil {
			// if claims is nil user will remain nil
			if claims, err := ParseToken(cookie); claims != nil && err == nil {
				u = am.us.GetOne(claims.Id)
			}
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
