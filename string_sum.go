package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
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

func StringSum(input string) (output string, err error) {
	runes := []rune(input)
	var cleanExpression string

	for _, k := range runes {
		switch k {
		case ' ':
			continue
		case '-':
			cleanExpression += "+-"
		case '+':
			fallthrough
		case '1':
			fallthrough
		case '2':
			fallthrough
		case '3':
			fallthrough
		case '4':
			fallthrough
		case '5':
			fallthrough
		case '6':
			fallthrough
		case '7':
			fallthrough
		case '8':
			fallthrough
		case '9':
			fallthrough
		case '0':
			cleanExpression += string(k)
		default:
			return "", &strconv.NumError{Func: "StringSum", Num: string(k), Err: fmt.Errorf("Invalid character: %s. Expression must contain only whitespace character, plus and minus operators and digits", string(k))}
		}
	}

	if cleanExpression == "" {
		return "", fmt.Errorf("Error: %s", errorEmptyInput)
	}

	operandsTmp := strings.Split(cleanExpression, "+")

	operands := make([]string, 0)

	for _, k := range operandsTmp {
		if k == "" {
			continue
		}

		operands = append(operands, k)
	}

	if len(operands) != 2 {
		return "", fmt.Errorf("Error: %w", errorNotTwoOperands)
	}

	var sum int

	for _, k := range operands {
		operand, err := strconv.Atoi(k)
		if err != nil {
			return "", err
		}

		sum += operand
	}

	return strconv.Itoa(sum), nil
}
