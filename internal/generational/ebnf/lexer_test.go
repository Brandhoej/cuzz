package ebnf

import "testing"

func TestLex(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		{
			name:  "Empty string",
			input: "",
			tokens: []Token{
				token(EOF, ""),
			},
		},
		{
			name:  "Left parenthesis",
			input: "(",
			tokens: []Token{
				token(LeftParenthesis, "("),
				token(EOF, ""),
			},
		},
		{
			name:  "Right parenthesis",
			input: ")",
			tokens: []Token{
				token(RightParenthesis, ")"),
				token(EOF, ""),
			},
		},
		{
			name:  "Left square bracket",
			input: "[",
			tokens: []Token{
				token(LeftSquareBracket, "["),
				token(EOF, ""),
			},
		},
		{
			name:  "Right square bracket",
			input: "]",
			tokens: []Token{
				token(RightSquareBracket, "]"),
				token(EOF, ""),
			},
		},
		{
			name:  "Dot",
			input: ".",
			tokens: []Token{
				token(Dot, "."),
				token(EOF, ""),
			},
		},
		{
			name:  "Equal",
			input: "=",
			tokens: []Token{
				token(Equal, "="),
				token(EOF, ""),
			},
		},
		{
			name:  "String",
			input: "\"Hello, World\"",
			tokens: []Token{
				token(String, "\"Hello, World\""),
				token(EOF, ""),
			},
		},
		{
			name:  "Identifier",
			input: "bar",
			tokens: []Token{
				token(Identifier, "bar"),
				token(EOF, ""),
			},
		},
		{
			name:  "Identifier with _",
			input: "foo_bar",
			tokens: []Token{
				token(Identifier, "foo_bar"),
				token(EOF, ""),
			},
		},
		{
			name:  "Golang Factor production",
			input: "Factor = production_name | token [ \"…\" token ] | Group | Option | Repetition .",
			tokens: []Token{
				token(Identifier, "Factor"),
				token(Equal, "="),
				token(Identifier, "production_name"),
				token(Pipe, "|"),
				token(Identifier, "token"),
				token(LeftSquareBracket, "["),
				token(String, "\"…\""),
				token(Identifier, "token"),
				token(RightSquareBracket, "]"),
				token(Pipe, "|"),
				token(Identifier, "Group"),
				token(Pipe, "|"),
				token(Identifier, "Option"),
				token(Pipe, "|"),
				token(Identifier, "Repetition"),
				token(Dot, "."),
				token(EOF, ""),
			},
		},
	}

	for _, test := range tests {
		lexer := LexString(test.input)

		for idx, expected := range test.tokens {
			actual, err := lexer.Next()
			if err != nil {
				t.Error(test.name, "- Error", err, "at index", idx)
			}

			if actual.Class != expected.Class {
				t.Error(test.name, "- Class", actual.Class, "expected", expected.Class, "at index", idx)
			}

			if actual.Lexeme != expected.Lexeme {
				t.Error(test.name, "- Lexeme", actual.Lexeme, "expected", expected.Lexeme, "at index", idx)
			}
		}

		token, err := lexer.Next()
		if token.Class != EOF {
			t.Error(test.name, "- Token exceeding EOF should be another EOF")
		}

		if err != nil {
			t.Error(test.name, "- Token exceeding EOF should not have any error")
		}
	}
}
