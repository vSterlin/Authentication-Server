package user

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

const (
	ACCESS_TOKEN = iota
	REFRESH_TOKEN
)

type claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func newClaims(u *User, exp int64) *claims {
	return &claims{
		u.Id,
		u.Username,
		u.Email,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "vSterlin",
		},
	}
}

func newCookie(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
		// Secure:   true,
		HttpOnly: true,
	}
}

func generateRefreshTokenCookie(u *User) *http.Cookie {

	rtSecret := []byte("REFRESH_TOKEN")
	c := newClaims(u, 100)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	rtStr, err := rt.SignedString(rtSecret)
	if err != nil {
		fmt.Println(err.Error())
	}

	return newCookie("refresh_token", rtStr)
}

func generateAccesTokenCookie(u *User) *http.Cookie {

	atSecret := []byte("ACCESS_TOKEN")
	c := newClaims(u, 1)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	atStr, err := at.SignedString(atSecret)
	if err != nil {
		fmt.Println(err.Error())
	}

	return newCookie("access_token", atStr)
}

func SetTokenCookies(w http.ResponseWriter, u *User) {
	at, rt := generateAccesTokenCookie(u), generateRefreshTokenCookie(u)
	http.SetCookie(w, at)
	http.SetCookie(w, rt)
}