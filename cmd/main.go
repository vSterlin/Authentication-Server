package main

import "github.com/vSterlin/auth/internal/server"

func main() {
	s := server.NewServer(8080)
	s.Init()
	defer s.Shutdown()

}
