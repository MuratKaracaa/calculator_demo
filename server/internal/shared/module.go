package shared

import "net/http"

type Module interface {
	RegisterValidations()
	RegisterRoutes(mux *http.ServeMux)
}
