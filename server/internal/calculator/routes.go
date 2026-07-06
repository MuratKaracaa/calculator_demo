package calculator

import "net/http"

func Register(mux *http.ServeMux) {
	service := newService()
	handler := newHandler(service)
	mux.HandleFunc("POST /calculator/calculate", handler.calculate)
}
