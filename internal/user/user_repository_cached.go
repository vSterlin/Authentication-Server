package user

import (
	"github.com/vSterlin/auth/internal/cache"
)

type cachedUserRepo struct {
	ur    UserRepo
	cache cache.Cache
}

func NewCachedUserRepo(ur UserRepo, c cache.Cache) UserRepo {
	return &cachedUserRepo{ur: ur, cache: c}
}

func (cur *cachedUserRepo) GetMany() []*User {
	// get from cache
	// if empty get by calling user repo's function
	return cur.ur.GetMany()
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
