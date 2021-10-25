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
