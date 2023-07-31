package symbol_test

import (
	"fmt"
	"testing"

	"slot/machine/models"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	symbols []models.Symbol
	hasWon bool
}

func TestHasSymbolWonSuccessfully(t *testing.T) {
	testcases := getTestCases()
	m := getMocks(t)

	s := getSymbolService(m)

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Testcase: %d has won runs successfully with result: %v",i, tc.hasWon), func(t *testing.T){
			hasWon, err := s.HasWon(tc.symbols)

			assert.Nil(t, err)
			assert.Equal(t, tc.hasWon, hasWon)
		})
	}
}

func TestHasSymbolWonFailsDueToNilSymbols(t *testing.T) {
	m := getMocks(t)

	s := getSymbolService(m)

	hasWon, err := s.HasWon(nil)

	assert.False(t, hasWon)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Symbols cannot be nil")
}

func getTestCases() []testCase {
	return []testCase {
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "*", IsAnyValue: true },
			},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "A"},
			},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "a"},
				models.Symbol{ Value: "a"},
			},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "*", IsAnyValue: true,
			}},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "*", IsAnyValue: true},
				models.Symbol{ Value: "*", IsAnyValue: true,
			}},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "*", IsAnyValue: true},
				models.Symbol{ Value: "*", IsAnyValue: true},
				models.Symbol{ Value: "*", IsAnyValue: true},
			},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A", IsAnyValue: true},
				models.Symbol{ Value: "B", IsAnyValue: true},
				models.Symbol{ Value: "C", IsAnyValue: true},
			},
			hasWon: true,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "B"},
				models.Symbol{ Value: "*", IsAnyValue: true},
			},
			hasWon: false,
		},
		testCase{
			symbols: []models.Symbol{
				models.Symbol{ Value: "A"},
				models.Symbol{ Value: "B"},
				models.Symbol{ Value: "C"},
			},
			hasWon: false,
		},
	}
}