package game

//go:generate mockgen --destination=../mocks/services/game/service.go slot/machine/services/game RowService,CoefficientService

import (
	"slot/machine/services/row"
	"slot/machine/config"
)

type RowService interface {
	GetRow() (*row.RowResponse, error)
}

type CoefficientService interface {
	Calcuate(values []float64) float64
}

type GameResponse struct {
	Lines []string
	HasWon bool
	TotalCoefficient float64
}

type Service struct {
 	rowService RowService
	coefficientService CoefficientService
	config *config.Config
}

func NewService(
	rowService RowService, 
	coefficientService CoefficientService,
	config *config.Config) *Service {
	return &Service{
		rowService: rowService,
		coefficientService: coefficientService,
		config: config,
	}
}