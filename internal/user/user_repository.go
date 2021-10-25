package user

type UserRepo interface {
	GetMany() []*User
}

type userRepo struct{}

var users = []*User{
	{"Vladimir", "Sterlin", "vSterlin", "hashedPassword"},
}

func (ur *userRepo) GetMany() []*User {
	return users
}
