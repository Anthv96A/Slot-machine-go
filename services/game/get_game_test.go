package game_test

import (
	"testing"
	"errors"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"slot/machine/services/row"
)

func TestGetGameSuccessfullyHasWon(t *testing.T) {
	m := getMocks(t)

	gomock.InOrder(
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "a,b,c", TotalCoefficient: 1.0  }, nil),
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "d,e,f", TotalCoefficient: 1.0  }, nil),
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "g,h,i", TotalCoefficient: 1.0  }, nil),
	)
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Return(3.0)

	s := getGameService(m)

	resp, err := s.GetGame()
	
	assert.Nil(t, err)

	var expectedGameSymbols = []string {
		"a,b,c",
		"d,e,f",
		"g,h,i",
	}

	assert.Equal(t, expectedGameSymbols, resp.Lines)
	assert.Equal(t, 3.0, resp.TotalCoefficient)
	assert.True(t, resp.HasWon)
}

func TestGetGameSuccessfullyHasNotWon(t *testing.T) {
	m := getMocks(t)

	gomock.InOrder(
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "a,b,c", TotalCoefficient: 0.0  }, nil),
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "d,e,f", TotalCoefficient: 0.0  }, nil),
		m.rowService.
			EXPECT().
			GetRow().
			Return(&row.RowResponse { Symbols: "g,h,i", TotalCoefficient: 0.0  }, nil),
	)
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Return(0.0)

	s := getGameService(m)

	resp, err := s.GetGame()
	
	assert.Nil(t, err)

	var expectedGameSymbols = []string {
		"a,b,c",
		"d,e,f",
		"g,h,i",
	}

	assert.Equal(t, expectedGameSymbols, resp.Lines)
	assert.Equal(t, 0.0, resp.TotalCoefficient)
	assert.False(t, resp.HasWon)
}

func TestGetGameFailsWhenRetrievingRows(t *testing.T) {
	m := getMocks(t)

	expectedErr := errors.New("wrong")

	m.rowService.
		EXPECT().
		GetRow().
		Return(nil, expectedErr)
	
	m.coefficientService.EXPECT().Calcuate(gomock.Any()).Times(0)

	s := getGameService(m)

	resp, err := s.GetGame()
	
	assert.Nil(t, resp)
	assert.Equal(t, err.Error(), "wrong: error has occurred getting rows")
}