package user

var users = []*User{
	{"Vladimir", "Sterlin", "vSterlin", "hashedPassword"},
}

type UserRepo interface {
	GetMany() []*User
}

type userRepo struct{}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (ur *userRepo) GetMany() []*User {
	return users
}
