package main

import (
	"github.com/aspandyar/echo-middleware-assessment/internal/pkg/app"
	"log"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	application.Run()
}
