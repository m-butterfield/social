package main

import (
	"github.com/m-butterfield/social/app/tasks"
	"log"
)

func main() {
	if err := tasks.Run("8001"); err != nil {
		log.Fatal(err)
	}
}
