package main

import (
	"log"

	app "falcon/app"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error in loading .env file %s", err.Error())
	}

	app := app.AppInstance()

	log.Fatal(app.Listen(":3000"))
}
