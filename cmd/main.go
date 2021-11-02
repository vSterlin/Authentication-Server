package main

import (
	"fmt"

	"github.com/vSterlin/auth/internal/config"
	"github.com/vSterlin/auth/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Println(err.Error())
	}
	s := server.NewServer(8080)
	s.Init()
	defer s.Shutdown()

}
