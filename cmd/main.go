package main

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"

	"github.com/vSterlin/auth/internal/cache"
	"github.com/vSterlin/auth/internal/config"
	"github.com/vSterlin/auth/internal/server"
)

func main() {

	if err := config.Load(); err != nil {
		fmt.Println(err.Error())
	}
	db, err := sql.Open("postgres", "user=v dbname=auth-server sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	}

	db.Ping()

	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	c := cache.NewCache(r)

	s := server.NewServer(8080, db, c)
	s.Init()
	defer s.Shutdown()

}
