package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os/user"
)

func CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("access_token")
		if err != nil {
			fmt.Println(err.Error())
		}

		// do jwt stuff with cookie
		// if cookie != nil {
		// 	jwt.Parse(cookie.Value, )
		// }

		var u *user.User
		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
