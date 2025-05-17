package internal_test

import (
	"testing"

	"github.com/J-Me-2307/expresso/internal"
)

func TestValidateTokens(t *testing.T) {
	tests := []struct {
		name   string
		tokens []*internal.Token
		want   int // Expected number of errors
	}{
		{
			name: "Valid expression",
			tokens: []*internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 0},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 1},
				{Type: internal.NUMBER_TOKEN, Value: "2", Position: 2},
			},
			want: 0,
		},
		{
			name: "Starts with operator",
			tokens: []*internal.Token{
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 0},
				{Type: internal.NUMBER_TOKEN, Value: "2", Position: 1},
			},
			want: 1,
		},
		{
			name: "Ends with operator",
			tokens: []*internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 0},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 1},
			},
			want: 1,
		},
		{
			name: "Consecutive operators",
			tokens: []*internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 0},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 1},
				{Type: internal.OPERATOR_TOKEN, Value: "*", Position: 2},
				{Type: internal.NUMBER_TOKEN, Value: "2", Position: 3},
			},
			want: 1,
		},
		{
			name: "Invalid token detected",
			tokens: []*internal.Token{
				{Type: internal.INVALID_TOKEN, Value: "@", Position: 0},
			},
			want: 1,
		},
		{
			name: "Unmatched closing parenthesis",
			tokens: []*internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 0},
				{Type: internal.RPAREN_TOKEN, Value: ")", Position: 1},
			},
			want: 1,
		},
		{
			name: "Unmatched opening parenthesis",
			tokens: []*internal.Token{
				{Type: internal.LPAREN_TOKEN, Value: "(", Position: 0},
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 1},
			},
			want: 1,
		},
		{
			name: "Balanced parentheses",
			tokens: []*internal.Token{
				{Type: internal.LPAREN_TOKEN, Value: "(", Position: 0},
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 1},
				{Type: internal.RPAREN_TOKEN, Value: ")", Position: 2},
			},
			want: 0,
		},
		{
			name: "Multiple errors in one expression",
			tokens: []*internal.Token{
				{Type: internal.OPERATOR_TOKEN, Value: "*", Position: 0},
				{Type: internal.NUMBER_TOKEN, Value: "1", Position: 1},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 2},
				{Type: internal.OPERATOR_TOKEN, Value: "-", Position: 3}, // Consecutive operators
				{Type: internal.RPAREN_TOKEN, Value: ")", Position: 4},   // Unmatched closing parenthesis
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.ValidateTokens(tt.tokens)

			if len(got) != tt.want {
				t.Errorf("ValidateTokens() errors = %d, want %d. Errors: %v", len(got), tt.want, got)
			}
		})
	}
}
