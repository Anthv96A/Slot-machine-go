package row_test

import (
	"testing"

	"slot/machine/services/row"
	mock "slot/machine/services/mocks/services/row"

	gomock "github.com/golang/mock/gomock"
)

type mocks struct {
	ctrl *gomock.Controller
	symbolsService *mock.MockSymbolsService
	coefficientService *mock.MockCoefficientService
}

func getMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	symbolsService := mock.NewMockSymbolsService(ctrl)
	coefficientService := mock.NewMockCoefficientService(ctrl)

	return &mocks {
		ctrl: ctrl,
		symbolsService: symbolsService,
		coefficientService: coefficientService,
	}
}

func getRowService(m *mocks) *row.Service {
	return row.NewService(m.symbolsService, m.coefficientService)
}

