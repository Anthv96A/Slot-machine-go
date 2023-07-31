package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"

	"slot/machine/models"
	"slot/machine/config"
	"slot/machine/services/game"
	"slot/machine/services/row"
	"slot/machine/services/symbol"
	"slot/machine/services/coefficient"
	"slot/machine/services/calculate_winnings"
)

func main() {
	config := getConfig()

	symbolsService := getSymbolService(config)

	coefficientService := getCofficientService()

	rowService := getRowService(coefficientService, symbolsService)
	gameService := getGameService(rowService, coefficientService, config)
	calcuateWinningsService := calculate_winnings.NewService()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your amount: ")

	amountVal, _ := reader.ReadString('\r')
	amountVal = strings.TrimSpace(amountVal)

	amount, _ := strconv.ParseFloat(amountVal, 64)

	for amount > 0 {
		fmt.Print("Enter your stake: ")
		stakedVal, _ := reader.ReadString('\r')
		stakedVal = strings.TrimSpace(stakedVal)

		staked, _ := strconv.ParseFloat(stakedVal, 64)

		resp, _ := gameService.GetGame()

		fmt.Println("Lines:")
		for _, line :=  range resp.Lines {
			fmt.Println(line)
		}

		calculatedResp := calcuateWinningsService.Calcuate(resp.TotalCoefficient, staked, amount)
		fmt.Println()
		fmt.Println("Winner?", resp.HasWon)
		fmt.Println("Total Coefficient:", resp.TotalCoefficient)
		fmt.Println("Amount Won:", calculatedResp.AmountWon)
		fmt.Println("New Balance:", calculatedResp.NewBalance)

		amount = calculatedResp.NewBalance
	}

	fmt.Println("Goodbye :)")
}

func getGameService(
	rowService *row.Service,
	coefficientService *coefficient.Service,
	config *config.Config) *game.Service {
	return game.NewService(rowService, coefficientService, config)
}

func getRowService(
	coefficientService *coefficient.Service,
	symbolsService *symbol.Service) *row.Service {
	return row.NewService(symbolsService, coefficientService)
}

func getCofficientService() *coefficient.Service {
	return coefficient.NewService()
}

func getSymbolService(config *config.Config) *symbol.Service {
	var symbols = []models.Symbol{}

	for _, symbol := range config.Symbols {
		for i :=0; i < symbol.Probabiity; i++ {
			symbols = append(symbols, symbol)
		}
	}

	return symbol.NewService(symbols, config)
}

func getConfig() *config.Config {
	return &config.Config {
		NumberOfRows: 4,
		NumberOfSymbols: 3,
		Symbols: []models.Symbol {
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
		},
	}
}
