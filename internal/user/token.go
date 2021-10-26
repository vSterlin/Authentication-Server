package user

import "net/http"

func newCookie(n string, v string) *http.Cookie {
	return &http.Cookie{
		Name:  n,
		Value: v,
		// Secure:   true,
		HttpOnly: true,
	}
}

func generateAccesToken(u *User) *http.Cookie {
	//TODO
	return newCookie("access_token", u.FirstName+u.Username+"access_token")
}

func generateRefreshToken(u *User) *http.Cookie {
	//TODO
	return newCookie("refresh_token", u.FirstName+u.Username+"refresh_token")
}

func SetTokenCookies(w http.ResponseWriter, u *User) {
	at, rt := generateAccesToken(u), generateRefreshToken(u)
	http.SetCookie(w, at)
	http.SetCookie(w, rt)
}
