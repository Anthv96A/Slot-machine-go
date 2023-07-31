package game

import (
	"fmt"

	"slot/machine/services/row"
) 

func (s *Service) GetGame() (*GameResponse, error) {
	rows, err := s.getRows()
	if err != nil {
	 	return nil,	fmt.Errorf("%w: error has occurred getting rows", err)
	}

	rowCoefficientValues := getRowCofficientValues(rows)
	gameCofficient := s.coefficientService.Calcuate(rowCoefficientValues)

	return &GameResponse {
		TotalCoefficient: gameCofficient,
		Lines: getGameLines(rows),
		HasWon: gameCofficient > 0.0,
	}, nil
}

func (s *Service) getRows() ([]row.RowResponse, error) {
	var rows = []row.RowResponse {}

	for i := 0; i < s.config.NumberOfRows; i++ {
		row, err := s.rowService.GetRow()
		if err != nil {
			return nil, err
		}

		rows = append(rows, *row)
	}

	return rows, nil
}

func getRowCofficientValues(rows []row.RowResponse) []float64 {
	var values = []float64{}

	for _, r := range rows {
		values = append(values, r.TotalCoefficient)
	}

	return values
}

func getGameLines(rows []row.RowResponse) []string {
	var values = []string{}

	for _, r := range rows {
		values = append(values, r.Symbols)
	}

	return values
}