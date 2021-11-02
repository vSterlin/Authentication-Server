package user

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type UserWithPassword struct {
	User
	Password string `json:"password"`
}

type ctx string

var UserContext = ctx("user")
