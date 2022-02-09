package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	us *UserService
	as *AuthService
}

func NewUserHandler(us *UserService, as *AuthService) *UserHandler {
	return &UserHandler{us, as}
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	u := uh.us.GetMany()
	json.NewEncoder(w).Encode(u)
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	uwp := &UserWithPassword{}
	json.NewDecoder(r.Body).Decode(&uwp)
	u := uh.as.SignUp(uwp)

	SetRefreshTokenCookie(w, u)
	SetAuthTokenCookie(w, u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (uh *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	uwp := &UserWithPassword{}
	json.NewDecoder(r.Body).Decode(&uwp)
	u, err := uh.as.SignIn(uwp.Email, uwp.Password)
	je := json.NewEncoder(w)
	if err != nil {
		je.Encode(err.Error())
		return
	}

	SetRefreshTokenCookie(w, u)
	SetAuthTokenCookie(w, u)

	je.Encode(u)
}

func (uh *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(UserContext).(*User)
	json.NewEncoder(w).Encode(u)
}

// TODO
func (uh *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if cookie, _ := r.Cookie("refresh_token"); cookie != nil {
		if claims, err := ParseToken(cookie); claims != nil && err == nil {
			u := uh.us.GetOne(claims.Id)
			at := generateAccesTokenCookie(u)
			http.SetCookie(w, at)
			return
		} else {
			// if something went wrong with cookie clear it
			c := newCookie("refresh_token", "")
			http.SetCookie(w, c)
		}
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
