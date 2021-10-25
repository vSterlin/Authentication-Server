package user

type AuthService struct {
	us *UserService
}

func NewAuthService(us *UserService) *AuthService {
	return &AuthService{us}
}

func (as *AuthService) SignUp(u *User) *User {
	// insert
}
