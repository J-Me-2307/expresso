package internal_test

import (
	"testing"

	"github.com/J-Me-2307/expresso/internal"
)

func TestToPostfix(t *testing.T) {
	tests := []struct {
		name   string
		tokens []*internal.Token
		want   []*internal.Token
	}{
		{
			name: "Simple Addition",
			tokens: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 3),
			},
			want: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 3),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
			},
		},
		{
			name: "Operator Precedence",
			tokens: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 3),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 4),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 5),
			},
			want: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 3),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 5),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 4),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
			},
		},
		{
			name: "With Parentheses",
			tokens: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
				internal.NewToken(internal.LPAREN_TOKEN, "(", 3),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 4),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 5),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 6),
				internal.NewToken(internal.RPAREN_TOKEN, ")", 7),
			},
			want: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "3", 1),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 4),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 6),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 5),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 2),
			},
		},
		{
			name: "Multiple Parentheses",
			tokens: []*internal.Token{
				internal.NewToken(internal.LPAREN_TOKEN, "(", 1),
				internal.NewToken(internal.NUMBER_TOKEN, "1", 2),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 3),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 4),
				internal.NewToken(internal.RPAREN_TOKEN, ")", 5),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 6),
				internal.NewToken(internal.LPAREN_TOKEN, "(", 7),
				internal.NewToken(internal.NUMBER_TOKEN, "3", 8),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 9),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 10),
				internal.NewToken(internal.RPAREN_TOKEN, ")", 11),
			},
			want: []*internal.Token{
				internal.NewToken(internal.NUMBER_TOKEN, "1", 2),
				internal.NewToken(internal.NUMBER_TOKEN, "2", 4),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 3),
				internal.NewToken(internal.NUMBER_TOKEN, "3", 8),
				internal.NewToken(internal.NUMBER_TOKEN, "4", 10),
				internal.NewToken(internal.OPERATOR_TOKEN, "+", 9),
				internal.NewToken(internal.OPERATOR_TOKEN, "*", 6),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.ToPostfix(tt.tokens)

			if len(got) != len(tt.want) {
				t.Errorf("Length mismatch: got %d tokens, want %d tokens", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i].Type != tt.want[i].Type || got[i].Value != tt.want[i].Value {
					t.Errorf("Mismatch at position %d: got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
