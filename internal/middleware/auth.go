package middleware

import (
	"net/http"

	"github.com/vSterlin/auth/internal/user"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if u := r.Context().Value(user.UserContext).(*user.User); u == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
