package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		cfg: cfg,
	}

	log.Fatal(app.run())
}
