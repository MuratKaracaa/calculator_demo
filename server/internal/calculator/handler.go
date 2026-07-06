package calculator

import (
	"net/http"
	"server/internal/shared"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) calculate(w http.ResponseWriter, r *http.Request) {
	req, err := shared.DecodeAndValidate[calculationRequest](r)
	if err != nil {
		shared.WriteJSON(w, http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
		return
	}

	result, calculationError := h.service.calculate(req.Expression)
	if calculationError != nil {
		shared.WriteJSON(w, http.StatusBadRequest, shared.ErrorResponse{Error: calculationError.Error()})
		return
	}

	shared.WriteJSON(w, http.StatusOK, calculationResponse{Result: result})
}
