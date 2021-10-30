package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/vSterlin/auth/internal/user"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u := r.Context().Value(user.UserContext).(*user.User)

		if u == nil {
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
