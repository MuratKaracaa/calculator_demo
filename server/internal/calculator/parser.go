package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
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

func isZeroLiteral(numStr string) bool {
	for _, c := range numStr {
		if c != '0' && c != '.' {
			return false
		}
	}
	return true
}

func validateNumberLiteral(token string) error {
	if token == "." {
		return fmt.Errorf("invalid number format")
	}

	dotCount := 0
	for _, c := range token {
		if c == '.' {
			dotCount++
			if dotCount > 1 {
				return fmt.Errorf("invalid number format")
			}
		}
	}

	_, err := strconv.ParseFloat(token, 64)
	return err
}

func validateExpression(expression string) error {
	if expression == "" {
		return errors.New("expression cannot be empty")
	}

	if strings.ContainsAny(expression, " \t\n\r") {
		return errors.New("expression cannot contain whitespace")
	}

	lastChar := string(expression[len(expression)-1])
	if binaryOperators[lastChar] {
		return fmt.Errorf("expression cannot end with operator '%s'", lastChar)
	}

	parenthesisStack := make([]int, 0)
	var prev rune
	hasPrev := false

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		isDigit := unicode.IsDigit(char)
		isDot := char == '.'

		switch {
		case char == '(':
			if hasPrev && (unicode.IsDigit(prev) || prev == ')' || prev == '%') {
				return fmt.Errorf("missing operator before '(' at position %d", i)
			}
			parenthesisStack = append(parenthesisStack, i)

		case char == ')':
			if len(parenthesisStack) == 0 {
				return fmt.Errorf("unmatched ')' at position %d", i)
			}

			openIndex := parenthesisStack[len(parenthesisStack)-1]
			parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]

			if i == openIndex+1 {
				return fmt.Errorf("empty parentheses at position %d", openIndex)
			}
			if binaryOperators[string(prev)] {
				return fmt.Errorf("dangling operator %q before ')' at position %d", prev, i)
			}

			if i+1 < len(expression) {
				nextChar := rune(expression[i+1])
				if unicode.IsDigit(nextChar) || nextChar == 's' || nextChar == '(' {
					return fmt.Errorf("missing operator between ')' and '%c' at position %d", nextChar, i+1)
				}
			}

		case binaryOperators[string(char)]:
			if char == '-' && (!hasPrev || prev == '(') {
				if i == len(expression)-1 {
					return fmt.Errorf("dangling operator '-' at position %d", i)
				}

				nextChar := rune(expression[i+1])
				if !unicode.IsDigit(nextChar) && nextChar != '(' && nextChar != 's' {
					return fmt.Errorf("invalid character after unary minus at position %d", i+1)
				}

				prev, hasPrev = char, true
				continue
			}

			if !hasPrev || prev == '(' {
				return fmt.Errorf("operator %q cannot follow '(' or be at start", char)
			}

			if char != '-' && binaryOperators[string(prev)] {
				return fmt.Errorf("operator %q cannot immediately follow operator %q", char, prev)
			}

			if char == '/' {
				j := i + 1
				start := j

				for j < len(expression) &&
					(unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
					j++
				}

				if start != j && isZeroLiteral(expression[start:j]) {
					return fmt.Errorf("division by zero at position %d", i)
				}
			}

		case char == '%':
			if !hasPrev {
				return fmt.Errorf("'%%' cannot be at start of expression at position %d", i)
			}

			if !unicode.IsDigit(prev) && prev != ')' {
				return fmt.Errorf("'%%' must follow a number or ')' at position %d", i)
			}

			if i+1 < len(expression) {
				nextChar := rune(expression[i+1])
				if unicode.IsDigit(nextChar) || nextChar == '.' {
					return fmt.Errorf("invalid character after '%%' at position %d", i+1)
				}
			}

		case char == 's':
			if i == len(expression)-1 {
				return fmt.Errorf("incomplete sqrt function at position %d", i)
			}

			nextChar := rune(expression[i+1])
			if !unicode.IsDigit(nextChar) && nextChar != '(' {
				return fmt.Errorf("sqrt must be followed by number or '(' at position %d", i+1)
			}

			if hasPrev && (unicode.IsDigit(prev) || prev == ')' || prev == '%') {
				return fmt.Errorf("missing operator before 's' at position %d", i)
			}

		case isDigit || isDot:
			if hasPrev && prev == ')' {
				return fmt.Errorf("missing operator between ')' and number at position %d", i)
			}

			if hasPrev && prev == '%' {
				return fmt.Errorf("invalid character after '%%' at position %d", i)
			}

			if !hasPrev || (!unicode.IsDigit(prev) && prev != '.') {
				start := i
				j := i

				for j < len(expression) &&
					(unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
					j++
				}

				if err := validateNumberLiteral(expression[start:j]); err != nil {
					return fmt.Errorf("%w at position %d", err, start)
				}

				i = j - 1
				char = rune(expression[j-1])
			}

		default:
			return fmt.Errorf("invalid character %q at position %d", char, i)
		}

		prev, hasPrev = char, true
	}

	if len(parenthesisStack) > 0 {
		return fmt.Errorf(
			"unbalanced parenthesis: %d unclosed '(' remaining",
			len(parenthesisStack),
		)
	}

	return nil
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
