package symbol

import (
	"slot/machine/models"
	"slot/machine/config"
)

func NewService(symbols []models.Symbol, config *config.Config) *Service {
	return &Service {
		symbols: symbols,
		config: config,
	}
}

type Service struct {
	symbols []models.Symbol
	config *config.Config
}

