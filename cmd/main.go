package main

import (
	"github.com/joho/godotenv"
	"github.com/vSterlin/auth/internal/server"
)

func main() {
	godotenv.Load()
	s := server.NewServer(8080)
	s.Init()
	defer s.Shutdown()

}
