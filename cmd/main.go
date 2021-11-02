package main

import (
	"database/sql"
	"fmt"

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

	s := server.NewServer(8080, db)
	s.Init()
	defer s.Shutdown()

}
