package main

import (
	"net/http"
	"server/cmd/api/calculator"
	"server/cmd/api/shared"
)

type config struct {
	addr string
}

type application struct {
	cfg config
}

func (app *application) run() error {
	mux := http.NewServeMux()
	shared.InitValidator()
	calculator.Register(mux)
	server := &http.Server{
		Addr:    app.cfg.addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}
