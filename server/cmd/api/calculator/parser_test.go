package calculator

import (
	"server/cmd/api/shared"
	"testing"
)

func TestShuntYardEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected float64
	}{
		{"addition", "2+3", 5},
		{"subtraction", "10-4", 6},
		{"multiplication", "3x4", 12},
		{"division", "12/3", 4},

		{"precedence_1", "2+3x4", 14},
		{"precedence_2", "10-2x3", 4},
		{"precedence_3", "20/2+5", 15},
		{"precedence_4", "2+3x4^2", 50},

		{"parentheses_1", "(2+3)x4", 20},
		{"parentheses_2", "2x(3+4)", 14},
		{"parentheses_nested", "((2+3)x(4+5))", 45},
		{"deep_parens", "((((5))))", 5},
		{"deep_complex", "(((2+3)x(4+5)))", 45},

		{"power_simple", "2^3", 8},
		{"power_right_assoc", "2^3^2", 512},
		{"power_parenthesized", "(2^3)^2", 64},
		{"power_parenthesized_2", "2^(3^2)", 512},
		{"zero_power", "0^5", 0},
		{"one_power", "1^999", 1},
		{"power_with_neg", "2^-3", 0.125},
		{"power_neg_base", "-2^3", -8},
		{"power_neg_base_paren", "(-2)^3", -8},

		{"sqrt_simple", "s9", 3},
		{"sqrt_expression", "s(16+9)", 5},
		{"sqrt_nested", "s(s16)", 2},
		{"sqrt_with_ops", "s25+5", 10},

		{"percentage_simple", "50%", 0.5},
		{"percentage_plus", "50%+1", 1.5},
		{"percentage_mult", "50%x200", 100},
		{"percentage_div", "100%/2", 0.5},
		{"percentage_paren", "(50+50)%", 1},
		{"percentage_complex", "200%+50%", 2.5},
		{"percentage_decimal", "12.5%", 0.125},
		{"percentage_neg", "-50%", -0.5},

		{"sqrt_plus_percentage", "s16+50%", 4.5},
		{"sqrt_times_percentage", "s(4+5)x20%", 0.6},
		{"complex_2", "s((2+3)^2)%+10", 10.05},
		{"complex_3", "(2+3)x(4+5)%x10", 4.5},
		{"complex_4", "s(2^3+1)%x(5+5)", 0.3},
		{"complex_5", "-2^3+4x5-10/2", 7},
		{"complex_6", "(-2+5)^2x(10-3)", 63},
		{"complex_7", "s(2^3+4^2)%+100", 100.04898979485566},
		{"complex_8", "(2+3)%x(10-4)", 0.3},
		{"complex_9", "s(9+16)x(2+3)%", 0.25},
		{"complex_10", "-s(4+5)%x(2+3)^2", -0.75},

		{"neg_simple", "-5", -5},
		{"neg_addition", "-5+3", -2},
		{"neg_subtraction", "5+-3", 2},
		{"neg_multiplication", "-2x-3", 6},
		{"neg_division", "-10/2", -5},
		{"neg_paren", "-(2+3)", -5},
		{"neg_nested", "-(-(-5))", -5},
		{"neg_with_power", "-2^3", -8},
		{"neg_paren_power", "(-2)^3", -8},
		{"neg_complex", "-(2+3)x-(4-1)", 15},

		{"decimal_simple", "2.5+3.5", 6},
		{"decimal_mult", "2.5x4", 10},
		{"decimal_div", "5/2.5", 2},
		{"decimal_power", "2.5^2", 6.25},
		{"decimal_complex", "2.5x(3.5+1.5)", 12.5},

		{"zero", "0", 0},
		{"zero_mult", "0x5", 0},
		{"zero_div", "0/5", 0},
		{"zero_power_zero", "0^0", 1},

		{"large_numbers", "1000000+500000", 1500000},
		{"large_mult", "1000x2000", 2000000},
		{"large_power", "10^6", 1000000},

		{"long_expression_1", "1+2+3+4+5+6+7+8+9+10", 55},
		{"long_expression_2", "1x2x3x4x5x6x7x8x9x10", 3628800},
		{"long_expression_3", "10-9+8-7+6-5+4-3+2-1", 5},
		{"long_expression_4", "(1+2)x(3+4)x(5+6)x(7+8)", 3465},

		{"paren_complex_1", "((2+3)x(4+5))/(6-3)", 15},
		{"paren_complex_2", "((2^3)+(3^2))x(4+5)", 153},
		{"paren_complex_3", "s((2+3)^2+(4+5)^2)", 10.295630140987},

		{"discount", "100-20%", 99.8},
		{"tax", "100+18%", 100.18},
		{"compound", "1000x(1+5%)^10", 1628.894626777442},
		{"sqrt_complex", "s((2^3+4^2)x5+10%)", 10.959014554237985},

		{"very_deep_parens", "((((((((((2+3))))))))))", 5},
		{"very_long_power", "2^2^2^2", 65536},
		{"mixed_ops", "2+3x4^2-5/10+2x(3+4)", 63.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shuntYardEvaluate(tt.expr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !shared.AlmostEqual(got, tt.expected) {
				t.Fatalf(
					"expr=%q expected=%v got=%v",
					tt.expr,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestShuntYardEvaluateErrors(t *testing.T) {
	tests := []string{
		"",
		"(",
		")",
		"((2+3)",
		"(2+3))",

		"+",
		"-",
		"x",
		"/",
		"^",

		"2+",
		"2-",
		"2x",
		"2/",
		"2^",

		"s",
		"%",
		"s()",

		"2(3+4)",
		"(2+3)(4+5)",
		"2s9",
		"s9(2+3)",
		"(2+3)4",
		"50%(2+3)",

		"2++3",
		"2xx3",
		"2//3",
		"2^^3",

		"abc",
		"2+abc",
	}

	for _, expr := range tests {
		t.Run(expr, func(t *testing.T) {
			_, err := shuntYardEvaluate(expr)
			if err == nil {
				t.Fatalf("expected error for %q", expr)
			}
		})
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := shuntYardEvaluate("5/0")

	if err == nil {
		t.Fatal("expected division by zero error")
	}
}
