package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func CurrentUser(next http.Handler) http.Handler {
	// do cookie stuff
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		fmt.Println(t1)
		next.ServeHTTP(w, r)
	})
}
