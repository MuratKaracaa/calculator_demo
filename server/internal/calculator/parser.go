package calculator

import (
	"errors"
	"math"
	"strconv"
)

type Token struct {
	Type  string
	Value string
}

var binaryOperators map[string]bool = map[string]bool{
	"+": true,
	"-": true,
	"x": true,
	"/": true,
	"^": true,
}

var paranthesis map[string]bool = map[string]bool{
	"(": true,
	")": true,
}

var unaryOperators map[string]bool = map[string]bool{
	"s": true,
	"%": true,
}

var operatorPriority map[string]int = map[string]int{
	"+":  1,
	"-":  1,
	"x":  2,
	"/":  2,
	"(":  0,
	")":  0,
	"s":  5,
	"^":  3,
	"%":  5,
	"u-": 4,
}

func infix(expression string) []Token {
	queue := make([]Token, 0)
	stack := make([]Token, 0)

	currentNumber := ""

	for i, char := range expression {
		if !binaryOperators[string(char)] && !unaryOperators[string(char)] && !paranthesis[string(char)] {
			currentNumber += string(char)
			continue
		}

		if currentNumber != "" {
			queue = append(queue, Token{
				Type:  "value",
				Value: currentNumber,
			})
			currentNumber = ""
		}

		switch char {

		case '(':
			stack = append(stack, Token{
				Type:  "operator",
				Value: string(char),
			})

		case ')':
			for len(stack) > 0 {
				topOperator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if topOperator.Value == "(" {
					break
				}

				queue = append(queue, topOperator)
			}

		default:

			var isUnaryMinus bool

			if i == 0 && char == '-' {
				isUnaryMinus = true
			} else if i != 0 && char == '-' {
				prev := string(expression[i-1])
				if binaryOperators[prev] || prev == "(" {
					isUnaryMinus = true
				}
			}

			if isUnaryMinus {
				stack = append(stack, Token{
					Type:  "operator",
					Value: "u-",
				})
				continue
			}

			currentPriority := operatorPriority[string(char)]

			for len(stack) > 0 {
				topOperator := stack[len(stack)-1]

				if topOperator.Value == "(" {
					break
				}

				topOperatorPriority := operatorPriority[topOperator.Value]

				if topOperatorPriority > currentPriority ||
					(topOperatorPriority == currentPriority && char != '^') {
					queue = append(queue, topOperator)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}

			stack = append(stack, Token{
				Type:  "operator",
				Value: string(char),
			})
		}
	}

	if currentNumber != "" {
		queue = append(queue, Token{
			Type:  "value",
			Value: currentNumber,
		})
	}
	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return queue
}

func postfix(queue []Token) (float64, error) {
	stack := make([]Token, 0)

	for len(queue) > 0 {

		firstFromQueue := queue[0]
		queue = queue[1:]

		if firstFromQueue.Type == "value" {
			stack = append(stack, firstFromQueue)
		} else {

			switch firstFromQueue.Value {

			case "s":
				topOfStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				parsedValue, err := strconv.ParseFloat(topOfStack.Value, 64)

				if err == nil {
					newValue := math.Sqrt(parsedValue)
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})
				}

			case "%":
				topOfStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				parsedValue, err := strconv.ParseFloat(topOfStack.Value, 64)

				if err == nil {
					newValue := parsedValue / 100
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})
				}

			case "u-":
				topOfStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				parsedValue, err := strconv.ParseFloat(topOfStack.Value, 64)

				if err == nil {
					newValue := parsedValue * -1
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})
				}

			default:

				rightElement := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				leftElement := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				rightValue, _ := strconv.ParseFloat(rightElement.Value, 64)
				leftValue, _ := strconv.ParseFloat(leftElement.Value, 64)

				switch firstFromQueue.Value {

				case "+":
					newValue := leftValue + rightValue
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})

				case "-":
					newValue := leftValue - rightValue
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})

				case "x":
					newValue := leftValue * rightValue
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})

				case "/":
					if rightValue == 0 {
						return 0, errors.New("Divider either is 0 or evalutes to 0")
					}
					newValue := leftValue / rightValue
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})

				case "^":
					newValue := math.Pow(leftValue, rightValue)
					newValueAsStr := strconv.FormatFloat(newValue, 'f', -1, 64)
					stack = append(stack, Token{Type: "value", Value: newValueAsStr})
				}
			}
		}
	}

	topOfStack := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	parsedValue, _ := strconv.ParseFloat(topOfStack.Value, 64)

	return parsedValue, nil
}

func shuntYardEvaluate(expression string) (float64, error) {
	validationError := validateExpression(expression)
	if validationError != nil {
		return 0, validationError
	}
	parsedQueue := infix(expression)
	calculatedValue, err := postfix(parsedQueue)
	if err != nil {
		return 0, err
	}

	return calculatedValue, nil
}
