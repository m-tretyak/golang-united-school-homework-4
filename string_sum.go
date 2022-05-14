package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorInvalidOperand = errors.New("invalid operand, cannot parse no more")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func getOperands(input string) ([]string, error) {
	result := make([]string, 0, 2)
	operand := []rune{'+'}
	for _, r := range input + " " {
		switch len(operand) {
		case 0:
			switch r {
			case ' ':
				continue
			case '+', '-':
				operand = append(operand, r)
				continue
			default:
				return []string{}, fmt.Errorf("%q : %w", r, errorInvalidOperand)
			}
		case 1:
			switch r {
			case ' ':
				continue
			case '+', '-':
				operand[0] = r
				continue
			default:
				operand = append(operand, r)
			}
		default:
			switch r {
			case ' ':
				result = append(result, string(operand))
				operand = []rune{}
				continue
			case '+', '-':
				result = append(result, string(operand))
				operand = []rune{r}
				continue
			default:
				operand = append(operand, r)
			}
		}
	}

	return result, nil
}

func StringSum(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	operands, err := getOperands(input)
	if len(operands) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	var result, num = 0, 0
	for _, operand := range operands {
		if num, err = idiots_written_tests(operand); err != nil {
			return "", fmt.Errorf("failed to convert %q: %w", operand, err)
		}

		result += num
	}

	return strconv.Itoa(result), nil
}

func idiots_written_tests(operand string) (int, error) {
	if operand[0] == '-' {
		return strconv.Atoi(operand)
	}

	return strconv.Atoi(operand[1:])
}
