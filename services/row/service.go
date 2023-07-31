//go:generate mockgen --destination=../mocks/services/row/service.go slot/machine/services/row SymbolsService,CoefficientService

package row

import (
	"slot/machine/models"
)

type SymbolsService interface {
	GetSymbols() []models.Symbol
	HasWon(symbols []models.Symbol) (bool, error)
}

type CoefficientService interface {
	Calcuate(values []float64) float64
}

type RowResponse struct {
  TotalCoefficient float64
  Symbols string
}

type Service struct {
	symbolsService SymbolsService
	coefficientService CoefficientService
}

func NewService(
	symbolsService SymbolsService, 
	coefficientService CoefficientService) *Service {
	return &Service{
		symbolsService: symbolsService,
		coefficientService: coefficientService,
	}
}