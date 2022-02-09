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

func (as *AuthService) SignUp(uwp *UserWithPassword) *User {
	// TODO find if user with email exists

	hp, err := bcrypt.GenerateFromPassword([]byte(uwp.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	uwp.Password = string(hp)
	u := as.us.InsertOne(uwp)
	return u
}

func (as *AuthService) SignIn(email string, password string) (*User, error) {
	uwp := as.us.GetOneByEmailWithPassword(email)
	err := bcrypt.CompareHashAndPassword([]byte(uwp.Password), []byte(password))
	fmt.Println(uwp)
	fmt.Println(password)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("wrong email and password combination")
	}
	u := &uwp.User

	return u, nil

}

func (as *AuthService) RefreshToken() {

}
