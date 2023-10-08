package main

import (
	// server "backend/internal/server"
	// "backend/internal/services/implementation"
	// "backend/internal/pkg/hasher/implementation"
	"backend/internal/app"
	"fmt"
)

func test(a int) {
	a = 5;
	fmt.Println(a)
}

func main() {
	// a := 10
	// fmt.Println(a)
	// test(a)
	// fmt.Println(a)
	// passwordHasher := hasherImplementation.NewBcryptHasher()
	// UserImplementation.NewUserImplementation(passwordHasher)

	var a app.App
	err := a.Init()
	fmt.Println(err)
	// if err != nil {
	// 	f
	// }
	// server.SetupServer().Run()
}