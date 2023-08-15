package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kouxi08/clean_architecture/driver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("server running...")
	driver.Serve(fmt.Sprintf("%s", os.Getenv("PORT")))
}
