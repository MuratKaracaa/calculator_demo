package calculator

type calculationRequest struct {
	Expression string `json:"expression"`
}

func (calReq calculationRequest) Validate() error {
	return validateExpression(calReq.Expression)
}

type calculationResponse struct {
	Result float64 `json:"result"`
}
