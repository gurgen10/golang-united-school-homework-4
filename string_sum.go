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
	input = strings.ReplaceAll(input, " ", "")
	leadingMinus := ""

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if input[0] == '-' {
		input = input[1:]
		leadingMinus = "-"
		fmt.Println(leadingMinus)
	}
	if input[0] == '+' {
		input = input[1:]
	}

	splitInputPlus := strings.Split(input, "+")
	splitInputMinus := strings.Split(input, "-")

	if len(splitInputPlus) > 1 && len(splitInputPlus) < 3 {
		val1, err1 := strconv.Atoi(leadingMinus + splitInputPlus[0])
		val2, err2 := strconv.Atoi(splitInputPlus[1])

		if err1 != nil {
			e1 := err1.(*strconv.NumError)
			e1.Err = errorNotTwoOperands
			return "", fmt.Errorf("%w", e1)
		}
		if err2 != nil {
			e2 := err2.(*strconv.NumError)
			e2.Err = errorNotTwoOperands
			return "", fmt.Errorf("%w", e2)
		}
		output = fmt.Sprintf("%v", int64(val1)+int64(val2))
		return output, nil
	} else if len(splitInputMinus) > 1 && len(splitInputMinus) < 3 {
		val1, err1 := strconv.Atoi(leadingMinus + splitInputMinus[0])
		val2, err2 := strconv.Atoi(splitInputMinus[1])

		if err1 != nil {
			e1 := err1.(*strconv.NumError)
			e1.Err = errorNotTwoOperands
			return "", fmt.Errorf("%w", e1)
		}
		if err2 != nil {
			e2 := err2.(*strconv.NumError)
			e2.Err = errorNotTwoOperands
			return "", fmt.Errorf("%w", e2)
		}
		output = fmt.Sprintf("%v", int64(val1)-int64(val2))
		return output, nil
	} else {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
}
