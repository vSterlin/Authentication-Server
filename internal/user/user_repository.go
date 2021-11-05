package user

import "database/sql"

const (
	GetManySQL       = `SELECT id, first_name, last_name, username, email FROM users;`
	GetOneSQL        = `SELECT id, first_name, last_name, username, email FROM users WHERE id = $1;`
	GetOneByEmailSQL = `SELECT id, first_name, last_name, username, email FROM users WHERE email = $1;`
	InsertOneSQL     = `INSERT INTO users (first_name, last_name, username, email, password) 
											VALUES ($1, $2, $3, $4, $5) 
											RETURNING id, first_name, last_name, username, email;`
)

type UserRepo interface {
	GetMany() []*User
	GetOne(id int) *User
	GetOneByEmail(email string) *UserWithPassword
	InsertOne(u *UserWithPassword) *User
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetMany() []*User {
	rows, _ := ur.db.Query(GetManySQL)
	users := []*User{}
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email)
		users = append(users, u)
	}
	return users
}

func (ur *userRepo) GetOne(id int) *User {
	u := &User{}
	ur.db.QueryRow(GetOneSQL, id).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email)
	return u
}

func (ur *userRepo) GetOneByEmail(email string) *UserWithPassword {
	uwp := &UserWithPassword{}
	ur.db.QueryRow(GetOneSQL, email).Scan(&uwp.Id, &uwp.FirstName, &uwp.LastName, &uwp.Username, &uwp.Email)
	return uwp
}

func (ur *userRepo) InsertOne(uwp *UserWithPassword) *User {
	u := &User{}
	ur.db.QueryRow(InsertOneSQL, uwp.FirstName, uwp.LastName, uwp.Username, uwp.Email, uwp.Password).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.Email)
	return u
}
