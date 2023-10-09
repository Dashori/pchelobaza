package main

import (
	server "backend/internal/server"
	// "backend/internal/services/implementation"
	// "backend/internal/pkg/hasher/implementation"
	"backend/internal/app"
	"fmt"
	"os"
)

func main() {
	var a app.App
	err := a.Init()
	fmt.Println(err)
	// if err != nil {
	// 	f
	// }
	host := os.Getenv("BACKEND_HOST")
	port := os.Getenv("BACKEND_PORT")
	fmt.Println("back ", host, port)
	server.SetupServer().Run(host + ":" + port)
}