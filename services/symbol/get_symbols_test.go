package symbol_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolsSuccessfully(t *testing.T) {
	m := getMocks(t)
	numOfSymbols := 5
	m.config.NumberOfSymbols = numOfSymbols

	s := getSymbolService(m)

	symbols := s.GetSymbols()
	assert.Equal(t, len(symbols), numOfSymbols)
}