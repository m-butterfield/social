package main

import (
	"github.com/m-butterfield/social/app/controllers"
	"log"
)

func main() {
	log.Println("Starting server...")
	if err := controllers.Run("8000"); err != nil {
		log.Fatalln(err)
	}
}
