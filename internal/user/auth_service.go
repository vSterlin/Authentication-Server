package user

import (
	"errors"
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

func (as *AuthService) SignIn(email string, password string) (*User, error) {
	u := as.us.GetOneByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return nil, errors.New("wrong email and password combination")
	}

	return u, nil

}

func (as *AuthService) RefreshToken() {

}
