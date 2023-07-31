package symbol_test

import (
	"testing"

	"slot/machine/services/symbol"
	"slot/machine/config"
	"slot/machine/models"

	gomock "github.com/golang/mock/gomock"
)

type mocks struct {
	ctrl *gomock.Controller
	config *config.Config
	symbols []models.Symbol
}

func getMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)

	config := &config.Config {
		NumberOfSymbols: 3,
		Symbols: getSymbolsForConfig(),
	}

	return &mocks {
		ctrl: ctrl,
		config: config,
		symbols: getSymbols(config),
	}
}

func getSymbolService(m *mocks) *symbol.Service {
	return symbol.NewService(m.symbols, m.config)
}

func getSymbolsForConfig() []models.Symbol {
	return []models.Symbol {
		models.Symbol {
			Value: "A",
			Coefficient: 0.4,
			Probabiity: 45,
			IsAnyValue: false,
		},
		models.Symbol {
			Value: "B",
			Coefficient: 0.6,
			Probabiity: 35,
			IsAnyValue: false,
		},
		models.Symbol {
			Value: "P",
			Coefficient: 0.8,
			Probabiity: 15,
			IsAnyValue: false,
		},
		models.Symbol {
			Value: "*",
			Coefficient: 0.0,
			Probabiity: 5,
			IsAnyValue: true,
		},
	}
}

func getSymbols(config *config.Config) []models.Symbol {
	var symbols = []models.Symbol{}

	for _, symbol := range config.Symbols {
		for i :=0; i < symbol.Probabiity; i++ {
			symbols = append(symbols, symbol)
		}
	}

	return symbols
}
