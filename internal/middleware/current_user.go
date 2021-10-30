package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vSterlin/auth/internal/user"
)

func CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u *user.User

		if cookie, err := r.Cookie("access_token"); err != nil {
			fmt.Println(err.Error())
		} else {
			u = user.ParseToken(cookie)
		}

		// get user by id from decoded jwt cookie
		fmt.Println(u)
		ctx := context.WithValue(r.Context(), user.UserContext, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
