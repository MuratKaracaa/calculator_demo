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

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) Run() error {
	mux := http.NewServeMux()
	shared.InitValidator()
	calculator.Register(mux)
	server := &http.Server{
		Addr:    app.Cfg.Addr,
		Handler: enableCORS(mux),
	}

	return server.ListenAndServe()
}
