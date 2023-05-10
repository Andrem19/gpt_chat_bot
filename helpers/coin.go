package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func CountPriceAndAmounts(commands []string) (string, error) {
	sumInvest, err := strconv.ParseFloat(commands[1], 32)

	if err != nil {
		return err.Error(), err
	}
	sumCoins, err := strconv.ParseFloat(commands[2], 32)
	if err != nil {
		return err.Error(), err
	}
	var sb strings.Builder
	sb.WriteString("Step 1:\n")

	nextPriceToSell_1 := (sumInvest + 50) / sumCoins
	amount_1 := 50 / nextPriceToSell_1

	sb.WriteString(fmt.Sprintf("Price: %f ", nextPriceToSell_1))
	sb.WriteString(fmt.Sprintf("Amount: %f", amount_1))
	sb.WriteString("\n\n")
	fmt.Println(sumCoins)
	sumCoins -= amount_1
	fmt.Println(sumCoins)
	sb.WriteString("Step 2:\n")

	nextPriceToSell_2 := (sumInvest + 50) / sumCoins
	amount_2 := 50 / nextPriceToSell_2
	sb.WriteString(fmt.Sprintf("Price: %f ", nextPriceToSell_2))
	sb.WriteString(fmt.Sprintf("Amount: %f", amount_2))
	sb.WriteString("\n\n")
	sb.WriteString(fmt.Sprintf("Amount sum: %f", amount_1 + amount_2))
	return sb.String(), err
}