package config

import (
	"slot/machine/models"
)

type Config struct {
	Symbols []models.Symbol
	NumberOfRows int
	NumberOfSymbols int
}

