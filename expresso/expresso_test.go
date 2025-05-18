package expresso_test

import (
	"testing"

	"github.com/J-Me-2307/expresso/expresso"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		// Basic Arithmetic
		{"Addition", "1 + 2", 3, false},
		{"Subtraction", "4 - 2", 2, false},
		{"Multiplication", "3 * 4", 12, false},
		{"Division", "8 / 2", 4, false},

		// Operator Precedence
		{"Precedence 1", "2 + 3 * 4", 14, false},
		{"Precedence 2", "2 * 3 + 4", 10, false},
		{"Precedence 3", "10 - 6 / 2", 7, false},

		// Parentheses
		{"Parentheses 1", "(2 + 3) * 4", 20, false},
		{"Parentheses 2", "2 * (3 + (4 * 5))", 46, false},
		{"Parentheses 3", "((1 + 2) * (3 + 4))", 21, false},

		// Negative Numbers & Unary Operators
		{"Unary Negative 1", "-5 + 3", -2, false},
		{"Unary Negative 2", "3 + -5", -2, false},
		{"Double Negative", "-(-3)", 3, false},

		// Decimals
		{"Decimal Addition", "3.5 + 1.2", 4.7, false},
		{"Decimal Division", "10 / 4", 2.5, false},

		// Whitespace Handling
		{"Whitespace Handling", "  3 +    4  *   2 ", 11, false},

		// Invalid Expressions
		{"Invalid Operator Sequence", "3 + * 4", 0, true},
		{"Division By Zero", "5 / (2 - 2)", 0, true},
		{"Mismatched Parentheses", "((3 + 2)", 0, true},
		{"Unknown Token", "abc + 3", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := expresso.Evaluate(tt.expression)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Evaluate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Evaluate() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
