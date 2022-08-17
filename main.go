package main

import (
	"fmt"
	"go/pkg/server"
	"go/pkg/db"

	"github.com/joho/godotenv"
)

//go:generate

func main() {

	if err := godotenv.Load("./configs/conf.env"); err != nil {
		fmt.Println("Can't find config file")
	}

	fmt.Println("Hello, World!")
	server.Init()
	db.Init()

}
