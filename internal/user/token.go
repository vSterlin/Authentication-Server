package user

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func newCookie(n string, v string) *http.Cookie {
	return &http.Cookie{
		Name:  n,
		Value: v,
		// Secure:   true,
		HttpOnly: true,
	}
}

func newClaims(u *User) *jwt.MapClaims {
	return &jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"email":    u.Email,
	}
}

func generateAccesTokenCookie(u *User) *http.Cookie {
	//TODO

	atSecret := []byte("abcdefg")

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(u))
	atStr, err := at.SignedString(atSecret)
	if err != nil {
		fmt.Println(err.Error())
	}
	return newCookie("access_token", atStr)
}

func generateRefreshTokenCookie(u *User) *http.Cookie {
	//TODO
	return newCookie("refresh_token", u.FirstName+u.Username+"refresh_token")
}

func SetTokenCookies(w http.ResponseWriter, u *User) {
	at, rt := generateAccesTokenCookie(u), generateRefreshTokenCookie(u)
	http.SetCookie(w, at)
	http.SetCookie(w, rt)
}
