package user

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	ACCESS_TOKEN_SECRET  = "ACCESS_TOKEN_SECRET"
	REFRESH_TOKEN_SECRET = "REFRESH_TOKEN_SECRET"
)

type claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func newClaims(u *User, exp time.Duration) *claims {

	return &claims{
		u.Id,
		u.Username,
		u.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
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

	// TODO fix expiration
	rtSecret := []byte(os.Getenv(REFRESH_TOKEN_SECRET))
	c := newClaims(u, 7*24*time.Hour)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	rtStr, err := rt.SignedString(rtSecret)
	if err != nil {
		fmt.Println(err.Error())
	}

	return newCookie("refresh_token", rtStr)
}

func generateAccesTokenCookie(u *User) *http.Cookie {

	// TODO fix expiration
	atSecret := []byte(os.Getenv(ACCESS_TOKEN_SECRET))
	c := newClaims(u, 15*time.Second)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	atStr, err := at.SignedString(atSecret)
	if err != nil {
		fmt.Println(err.Error())
	}

	return newCookie("access_token", atStr)
}

func SetRefreshTokenCookie(w http.ResponseWriter, u *User) {
	rt := generateRefreshTokenCookie(u)

	http.SetCookie(w, rt)
}

func SetAuthTokenCookie(w http.ResponseWriter, u *User) {
	at := generateAccesTokenCookie(u)
	http.SetCookie(w, at)
}

func SetTokenCookies(w http.ResponseWriter, u *User) {
	SetRefreshTokenCookie(w, u)
	SetAuthTokenCookie(w, u)
}

func ParseToken(cookie *http.Cookie) (*claims, error) {

	t, err := jwt.ParseWithClaims(cookie.Value, &claims{}, func(at *jwt.Token) (interface{}, error) {
		s := ""
		if cookie.Name == "access_token" {
			s = os.Getenv(ACCESS_TOKEN_SECRET)
		} else {
			s = os.Getenv(REFRESH_TOKEN_SECRET)
		}
		return []byte(s), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*claims); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("could not parse token")
}
