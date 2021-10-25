package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	us *UserService
}

func NewUserController(us *UserService) *UserController {
	return &UserController{us}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	u := uc.us.GetMany()
	json.NewEncoder(w).Encode(u)
}
