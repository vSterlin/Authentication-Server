package user

import "database/sql"

var users = []*User{
	{1, "Vladimir", "Sterlin", "vSterlin", "v@v.com", "hashedPassword"},
}

const (
	GetManySQL       = ``
	GetOneSQL        = ``
	GetOneByEmailSQL = ``
	InsertOneSQL     = `INSERT INTO users (first_name, last_name, username, email, password) 
											VALUES ($1, $2, $3, $4, $5) 
											RETURNING id, first_name, last_name, username, email;`
)

type UserRepo interface {
	GetMany() []*User
	GetOne(id int) *User
	GetOneByEmail(email string) *User
	InsertOne(u *User) *User
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetMany() []*User {
	return users
}

func (ur *userRepo) GetOne(id int) *User {
	return users[id-1]
}

func (ur *userRepo) GetOneByEmail(email string) *User {
	for _, u := range users {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func (ur *userRepo) InsertOne(u *User) *User {
	ur.db.QueryRow(InsertOneSQL, u.FirstName, u.LastName, u.Username, u.Email, u.Password).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email)
	return u
}
