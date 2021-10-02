package records

import (
	"fmt"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

type balanceRecords map[invest.Currency]float64

// GetDiffString - get price difference between portfolio positions
// and convert to string.
func GetDiffString(positions []invest.PositionBalance) string {
	diff := GetDiff(positions)

	return diffToString(diff)
}

// GetDiffString - get price difference between portfolio positions.
func GetDiff(positions []invest.PositionBalance) float64 {
	startRecords := balanceRecords{}
	actualRecords := balanceRecords{}

	for _, p := range positions {
		posValue := p.AveragePositionPrice.Value * p.Balance
		actualValue := posValue + p.ExpectedYield.Value

		currency := p.ExpectedYield.Currency

		startRecords[currency] += posValue
		actualRecords[currency] += actualValue
	}

	return countRecordsDiff(startRecords, actualRecords)
}

func countRecordsDiff(start balanceRecords, actual balanceRecords) float64 {
	const (
		hundredPercent = 100
		pivot          = 1
	)

	var diff float64

	for k := range start {
		if actual[k] == 0 || start[k] == 0 {
			continue
		}

		diff += (actual[k]/start[k] - pivot) * hundredPercent
	}

	return diff
}

func diffToString(diff float64) string {
	if diff == 0.0 {
		return "0%"
	}

	diffString := fmt.Sprintf("%.2f%%", diff)

	if diffString[0] != '-' && diff != 0.0 {
		diffString = "+" + diffString
	}

	return diffString
}
