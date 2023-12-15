package main

import (
	"backend/internal/app"
	"backend/internal/server"
	"fmt"
)

func main() {
	var a app.App
	err := a.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.SetupServer(&a).Run()
	if err != nil {
		fmt.Println(err)
	}
}
