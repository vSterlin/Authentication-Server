package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	us *UserService
}

func NewAuthService(us *UserService) *AuthService {
	return &AuthService{us}
}

func (as *AuthService) SignUp(u *User) *User {
	// TODO find if user with email exists

	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	u.Password = string(hp)
	u = as.us.InsertOne(u)
	return u
}
