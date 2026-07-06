package calculator

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func isZeroLiteral(numStr string) bool {
	for _, c := range numStr {
		if c != '0' && c != '.' {
			return false
		}
	}
	return true
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

	for i, char := range expression {
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
				for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
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

		default:
			return fmt.Errorf("invalid character %q at position %d", char, i)
		}

		prev, hasPrev = char, true
	}

	if len(parenthesisStack) > 0 {
		return fmt.Errorf("unbalanced parenthesis: %d unclosed '(' remaining", len(parenthesisStack))
	}

	return nil
}
