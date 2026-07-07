package calculator

import "net/http"

type Module struct{}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) RegisterRoutes(mux *http.ServeMux) {
	service := newService()
	handler := newHandler(service)
	mux.HandleFunc("POST /calculator/calculate", handler.calculate)
}

func (m *Module) RegisterValidations() {
	registerValidations()
}
