package app

import (
	"net/http"
	"server/internal/calculator"
	"server/internal/shared"
)



type Config struct {
	Addr string
}

type Application struct {
	Cfg Config
}

func (app *Application) Run() error {
	mux := http.NewServeMux()
	shared.InitValidator()
	calculator.Register(mux)
	server := &http.Server{
		Addr:    app.Cfg.Addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}
