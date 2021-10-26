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
	u := &User{}
	json.NewDecoder(r.Body).Decode(&u)
	u = uc.as.SignUp(u)

	SetTokenCookies(w, u)

	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	u := &User{}
	json.NewDecoder(r.Body).Decode(&u)
	u, err := uc.as.SignIn(u.Email, u.Password)
	je := json.NewEncoder(w)
	if err != nil {
		je.Encode(err.Error())
		return
	}

	SetTokenCookies(w, u)

	je.Encode(u)
}
