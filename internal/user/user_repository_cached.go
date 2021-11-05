package user

import (
	"encoding/json"
	"fmt"

	"github.com/vSterlin/auth/internal/cache"
)

type cachedUserRepo struct {
	ur UserRepo
	c  cache.Cache
}

func NewCachedUserRepo(ur UserRepo, c cache.Cache) UserRepo {
	return &cachedUserRepo{ur: ur, c: c}
}

func (cur *cachedUserRepo) GetMany() []*User {
	// get from cache
	// if empty get by calling user repo's function
	users, err := cur.c.Get("users")

	if err != nil {
		u := cur.ur.GetMany()
		cur.c.Set("users", u)
		return u
	}

	fmt.Println("FROM CACHE")
	u := []*User{}
	json.Unmarshal([]byte(users), &u)
	return u
}

func (cur *cachedUserRepo) GetOne(id int) *User {
	return cur.ur.GetOne(id)

}

func (cur *cachedUserRepo) GetOneByEmail(email string) *UserWithPassword {
	return cur.ur.GetOneByEmail(email)

}

func (cur *cachedUserRepo) InsertOne(uwp *UserWithPassword) *User {
	return cur.ur.InsertOne(uwp)

}
