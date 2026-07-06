package calculator

import "server/internal/shared"

type calculationRequest struct {
	Expression string `json:"expression" validate:"required,calculation_expression"`
}

func (calReq calculationRequest) Validate() error {
	return shared.ValidateInstance.Struct(calReq)
}

type calculationResponse struct {
	Result float64 `json:"result"`
}
