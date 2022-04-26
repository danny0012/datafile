// Package datafile reads data from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloats reads float64 value from strings of file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberString := strings.TrimSpace(scanner.Text())
		number, err := strconv.ParseFloat(numberString, 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	
	return numbers, nil
}