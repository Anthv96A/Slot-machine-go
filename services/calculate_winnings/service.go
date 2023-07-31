package calculate_winnings

import (
	"math"
)

type Service struct {}

type CalculatedResponse struct {
	AmountWon float64
	NewBalance float64
}

func NewService() *Service {
	return &Service{}
}

func(s *Service) Calcuate(coefficient, amountStaked, totalBalance float64) *CalculatedResponse{
	precision := math.Pow(10, float64(2))

	totalAmountWon := coefficient * amountStaked
	totalNewBalance := (totalBalance - amountStaked) + totalAmountWon

	return &CalculatedResponse {
		AmountWon: roundTo(totalAmountWon, precision),
		NewBalance: roundTo(totalNewBalance, precision),
	}
}

func roundTo(value, precision float64) float64 {
	return math.Round(value * precision) / precision
}