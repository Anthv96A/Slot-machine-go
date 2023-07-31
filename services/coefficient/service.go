package coefficient

type Service struct {}

func NewService () *Service {
	return &Service{}
}

func (s *Service) Calcuate(values []float64) float64 {
	var total float64

	for _, val := range values {
		total += val
	} 

	return total
}