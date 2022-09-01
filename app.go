package main

import (
	"fmt"

	"github.com/joho/godotenv"
	server "go-fiber.com/src"
	"go-fiber.com/src/db"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Can't find env file")
	}

	db.Init()
	// logger.Init()
	server.Init()
}
