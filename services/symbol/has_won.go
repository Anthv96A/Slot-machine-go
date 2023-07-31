package symbol

import (
	"fmt"
	"strings"

	"slot/machine/models"
)

func(s *Service) HasWon(symbols []models.Symbol) (bool, error) {
	if symbols == nil {
		return false, fmt.Errorf("Symbols cannot be nil")
	}

	if len(symbols) == 0 {
		return false, nil
	}
	
	return isWinner(symbols), nil
}

func isWinner(symbols []models.Symbol) bool {
	var excludeIsAny []models.Symbol
	for _, s := range symbols {
		if !s.IsAnyValue {
			excludeIsAny = append(excludeIsAny, s)
		}
	}

	if len(excludeIsAny) == 0 {
		return true
	}

	groupBy := make(map[string]bool)
	for _, s := range excludeIsAny {
		val := strings.ToLower(s.Value)
		
		if _, ok := groupBy[val]; !ok {
			groupBy[val] = true
		}
	}

	return len(groupBy) == 1
}