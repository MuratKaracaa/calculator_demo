package main

import (
	"log"
	"server/internal/app"
)


func main() {
	cfg := app.Config{
		Addr: ":8080",
	}

	app := &app.Application{
		Cfg: cfg,
	}

	log.Fatal(app.Run())
}
