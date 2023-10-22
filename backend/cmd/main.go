package main

import (
	server "backend/internal/server"
	"backend/internal/app"
	"fmt"
	"os"
)

func main() {
	var a app.App
	err := a.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	host := os.Getenv("BACKEND_HOST")
	port := os.Getenv("BACKEND_PORT")
	fmt.Println("back ", host, port)
	err = server.SetupServer().Run(host + ":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
