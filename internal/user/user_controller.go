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
	json.NewEncoder(w).Encode(u)
}
