package symbol

import (
	"math/rand"
	"time"

	"slot/machine/models"
)

func (s *Service) GetSymbols() []models.Symbol {
	ramdomisedSymbols := randomiseSymbols(s.symbols)

	return ramdomisedSymbols[:s.config.NumberOfSymbols]
}

func randomiseSymbols(symbols []models.Symbol) []models.Symbol{
	rand.Seed(time.Now().UnixNano()) 

	for i := len(symbols) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		symbols[i], symbols[j] = symbols[j], symbols[i]
	}


	return symbols
}
