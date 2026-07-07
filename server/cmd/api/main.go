package main

import (
	"fmt"
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

	fmt.Println("Server started on port :8080")

	log.Fatal(app.Run())
}
