package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	us *UserService
	as *AuthService
}

func NewUserController(us *UserService, as *AuthService) *UserController {
	return &UserController{us, as}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	u := uc.us.GetMany()
	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	uwp := &UserWithPassword{}
	json.NewDecoder(r.Body).Decode(&uwp)
	u := uc.as.SignUp(uwp)

	SetTokenCookies(w, u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	uwp := &UserWithPassword{}
	json.NewDecoder(r.Body).Decode(&uwp)
	u, err := uc.as.SignIn(uwp.Email, uwp.Password)
	je := json.NewEncoder(w)
	if err != nil {
		je.Encode(err.Error())
		return
	}

	SetTokenCookies(w, u)

	je.Encode(u)
}

func (uc *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(UserContext).(*User)
	json.NewEncoder(w).Encode(u)
}

// TODO
func (uc *UserController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if cookie, _ := r.Cookie("refresh_token"); cookie != nil {
		if claims, err := ParseToken(cookie); claims != nil && err == nil {
			u := uc.us.GetOne(claims.Id)
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
