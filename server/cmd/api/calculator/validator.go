package calculator

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var binaryOperators = map[rune]bool{
	'+': true,
	'-': true,
	'x': true,
	'/': true,
	'^': true,
}

var postfixUnaryOperators = map[rune]bool{
	's': true,
	'%': true,
}

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

	runes := []rune(expression)
	parenStack := make([]int, 0)

	var prev rune
	hasPrev := false

	for i, char := range runes {
		isDigit := unicode.IsDigit(char)
		isDot := char == '.'
		isBinary := binaryOperators[char]
		isPostfixUnary := postfixUnaryOperators[char]

		switch {
		case char == '(':
			parenStack = append(parenStack, i)

			if hasPrev && (unicode.IsDigit(prev) || prev == ')' || postfixUnaryOperators[prev]) {
				return fmt.Errorf("missing operator before '(' at position %d", i)
			}

		case char == ')':
			if len(parenStack) == 0 {
				return fmt.Errorf("unmatched ')' at position %d", i)
			}

			openIndex := parenStack[len(parenStack)-1]
			parenStack = parenStack[:len(parenStack)-1]

			if i == openIndex+1 {
				return fmt.Errorf("empty parentheses at position %d", openIndex)
			}
			if binaryOperators[prev] {
				return fmt.Errorf("dangling operator %q before ')' at position %d", prev, i)
			}

		case isBinary:
			if char == '-' && (!hasPrev || prev == '(' || binaryOperators[prev]) {
				break
			}
			if !hasPrev || prev == '(' {
				return fmt.Errorf("operator %q cannot follow '(' or be at start", char)
			}
			if binaryOperators[prev] {
				return fmt.Errorf("operator %q cannot immediately follow operator %q", char, prev)
			}

			if char == '/' {
				j := i + 1
				start := j
				for j < len(runes) && (unicode.IsDigit(runes[j]) || runes[j] == '.') {
					j++
				}
				if start != j && isZeroLiteral(string(runes[start:j])) {
					return fmt.Errorf("division by zero at position %d", i)
				}
			}

		case isPostfixUnary:
			if !hasPrev || !(unicode.IsDigit(prev) || prev == '.' || prev == ')' || postfixUnaryOperators[prev]) {
				return fmt.Errorf("operator %q must follow a number", char)
			}

		case isDigit || isDot:

		default:
			return fmt.Errorf("invalid character %q at position %d", char, i)
		}

		prev, hasPrev = char, true
	}

	if len(parenStack) > 0 {
		return fmt.Errorf("unbalanced parenthesis: %d unclosed '(' remaining", len(parenStack))
	}

	lastChar := runes[len(runes)-1]
	if binaryOperators[lastChar] || lastChar == '(' {
		return fmt.Errorf("expression cannot end with %q", lastChar)
	}

	return nil
}
