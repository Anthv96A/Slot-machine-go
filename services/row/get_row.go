package row

import (
	"fmt"
	"strings"

	"slot/machine/models"
)

func (s *Service) GetRow() (*RowResponse, error) {
	symbols := s.symbolsService.GetSymbols()
	
	isWinner, err := s.symbolsService.HasWon(symbols)
	if err != nil {
		return nil, fmt.Errorf("%w: error has occurred", err)
	}

	if !isWinner {
		return &RowResponse {
			TotalCoefficient: 0.0,
			Symbols: getSymbolsLine(symbols),
		}, nil
	}

	symbolCofficientValues := getCofficientValues(symbols)
	totalCoefficient := s.coefficientService.Calcuate(symbolCofficientValues)

	return &RowResponse {
		TotalCoefficient: totalCoefficient,
		Symbols: getSymbolsLine(symbols),
	}, nil
}

func getCofficientValues(symbols []models.Symbol) []float64 {
	var values = []float64 {}

	for _, s := range symbols {
		values = append(values, s.Coefficient)
	}

	return values
}


func getSymbolsLine(symbols []models.Symbol) string {
	var values = []string {}

	for _, s := range symbols {
		values = append(values, s.Value)
	}

	return strings.Join(values, ",")
}