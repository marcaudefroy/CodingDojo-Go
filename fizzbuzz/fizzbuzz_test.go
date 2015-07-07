package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	res, _ := FizzBuzz(1)
	assert.Equal(t,
		"1",
		res,
		"Should be the same")

	res, _ = FizzBuzz(2)
	assert.Equal(t,
		"2",
		res,
		"Should be the same")

	res, _ = FizzBuzz(3)
	assert.Equal(t,
		"fizz",
		res,
		"Should be the same")
}

func TestFizzBuzzNegativeValue(t *testing.T) {
	res, err := FizzBuzz(-1)
	assert.Equal(t,
		"",
		res,
		"Should not return a value")

	assert.EqualError(t,
		err,
		"The parameter value -1 must be a positive integer",
		"must return an error")
}
