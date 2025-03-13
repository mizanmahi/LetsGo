package fileUtils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// name of the function should start with capital letter to make it public (exported)
func SaveBalanceToFile(balance float64) {

	balanceString := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceString), 0644)
}

func ReadBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("file not found")
	}
	balanceString := string(balance)
	return strconv.ParseFloat(balanceString, 64)

}