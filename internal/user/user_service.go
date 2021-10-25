package user

type UserService struct {
	ur UserRepo
}

func NewUserService(ur UserRepo) *UserService {
	return &UserService{ur}
}

func (us *UserService) GetMany() []*User {
	return us.ur.GetMany()
}

func (us *UserService) GetOne(id int) *User {
	return us.ur.GetOne(id)
}

func (us *UserService) InsertOne(u *User) *User {
	return us.ur.InsertOne(u)
}
