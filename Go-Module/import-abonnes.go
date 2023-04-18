package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hello, World!")

	if err := godotenv.Load(); err != nil {
		fmt.Println("Pas de fichier .env")
		os.Exit(1)
	}

	fmt.Println(os.Getenv("DB_HOST"))

}
