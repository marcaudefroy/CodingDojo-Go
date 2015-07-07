package fizzbuzz

import (
	"fmt"
	"strconv"
)

// FizzBuzz return an integer or "fizz" or "buzz"
func FizzBuzz(value int) (string, error) {
	if value < 0 {
		//return "", errors.New("The parameter  must be a positive integer")
		return "", fmt.Errorf("The parameter value %d must be a positive integer", value)
	}
	restDivBy3 := value % 3
	if restDivBy3 == 0 {
		return "fizz", nil
	}

	return strconv.Itoa(value), nil
}
