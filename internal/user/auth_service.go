package user

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

type AuthService struct {
	us *UserService
}

func NewAuthService(us *UserService) *AuthService {
	return &AuthService{us}
}

func (as *AuthService) SignUp(u *User) *User {
	// TODO find if user with email exists

	// generate random
	salt := "asdasd"
	// do password hashing
	hp, err := scrypt.Key([]byte(u.Password), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		fmt.Println(err)
	}

	u.Password = string(hp)
	u = as.us.InsertOne(u)
	return u
}
