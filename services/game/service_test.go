package game_test

import (
	"testing"

	"slot/machine/services/game"
	mock "slot/machine/services/mocks/services/game"
	"slot/machine/config"

	gomock "github.com/golang/mock/gomock"
)

type mocks struct {
	ctrl *gomock.Controller
	rowService *mock.MockRowService
	coefficientService *mock.MockCoefficientService
	config *config.Config
	numberOfRows int
}

func getMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	rowService := mock.NewMockRowService(ctrl)
	coefficientService := mock.NewMockCoefficientService(ctrl)
	numberOfRows := 3

	config := &config.Config {
		NumberOfRows:numberOfRows,
	}

	return &mocks {
		ctrl: ctrl,
		rowService: rowService,
		coefficientService: coefficientService,
		config: config,
		numberOfRows: numberOfRows,
	}
}

func getGameService(m *mocks) *game.Service {
	return game.NewService(m.rowService, m.coefficientService, m.config)
}
