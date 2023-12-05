package main

import (
	"backend/internal/app"
	"backend/internal/server"
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
	err = server.SetupServer(&a).Run()
	if err != nil {
		fmt.Println(err)
	}
}
