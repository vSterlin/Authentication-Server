package user

var users = []*User{
	{1, "Vladimir", "Sterlin", "vSterlin", "hashedPassword"},
}

type UserRepo interface {
	GetMany() []*User
	GetOne(id int) *User
	InsertOne(u *User) *User
}

type userRepo struct{}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (ur *userRepo) GetMany() []*User {
	return users
}

func (ur *userRepo) GetOne(id int) *User {
	return users[id-1]
}

func (ur *userRepo) InsertOne(u *User) *User {
	u.Id = len(users) + 1
	users = append(users, u)
	return u
}
