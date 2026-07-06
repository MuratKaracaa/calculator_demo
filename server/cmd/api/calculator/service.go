package calculator

type service struct{}

func newService() *service {
	return &service{}
}

func (s *service) calculate(expression string) (float64, error) {
	return shuntYardEvaluate(expression)
}
