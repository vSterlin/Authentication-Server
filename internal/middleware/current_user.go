package middleware

import (
	"fmt"
	"net/http"
)

func CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("access_token")
		if err != nil {
			fmt.Println(err.Error())
		}

		// do jwt stuff with cookie

		next.ServeHTTP(w, r)
	})
}
