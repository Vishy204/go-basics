package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	balanceData := string(data)
	value, err := strconv.ParseFloat(balanceData, 64)
	if err != nil {
		return 1000, errors.New("failed to parse to float")
	}
	return value, nil

}

func WriteFloatToFile(value float64, fileName string) {
	value_text := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(value_text), 0644)

}
