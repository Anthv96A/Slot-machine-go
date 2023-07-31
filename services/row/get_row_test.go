package row_test

import (
	"testing"
	"errors"

	"slot/machine/models"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWinnerRowSuccessfully(t *testing.T) {
	m := getMocks(t)
	symbols := getSymbols()

	m.symbolsService.EXPECT().GetSymbols().Return(symbols)
	m.symbolsService.EXPECT().HasWon(symbols).Return(true, nil)
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Return(6.0)
	
	s := getRowService(m)

	resp, err := s.GetRow()

	assert.Nil(t, err)
	assert.Equal(t, "A,B,C", resp.Symbols)
	assert.Equal(t, float64(6.0), resp.TotalCoefficient)
}

func TestGetLostRowSuccessfully(t *testing.T) {
	m := getMocks(t)
	symbols := getSymbols()

	m.symbolsService.EXPECT().GetSymbols().Return(symbols)
	m.symbolsService.EXPECT().HasWon(symbols).Return(false, nil)
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Times(0)
	
	s := getRowService(m)

	resp, err := s.GetRow()

	assert.Nil(t, err)
	assert.Equal(t, "A,B,C", resp.Symbols)
	assert.Equal(t, float64(0), resp.TotalCoefficient)
}

func TestGetHasWonFails(t *testing.T) {
	m := getMocks(t)
	symbols := getSymbols()
	expectedErr := errors.New("Some error")

	m.symbolsService.EXPECT().GetSymbols().Return(symbols)
	m.symbolsService.EXPECT().HasWon(symbols).Return(false, expectedErr)
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Times(0)
	
	s := getRowService(m)

	resp, err := s.GetRow()

	assert.Nil(t, resp)
	assert.Equal(t, err.Error(), "Some error: error has occurred")
}

func getSymbols() []models.Symbol {
	return []models.Symbol {
		models.Symbol{
			Value: "A",
			Coefficient: 2,
		},
		models.Symbol{
			Value: "B",
			Coefficient: 2,
		},
		models.Symbol{
			Value: "C",
			Coefficient: 2,
		},
	}
}
