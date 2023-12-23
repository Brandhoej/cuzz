package ebnf

import (
	"errors"
	"reflect"
	"testing"
)

func TestProduction(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output ProductionAST
	}{
		{
			name:  "Literal string",
			input: "initial = \"Hello, World!\"",
			output: ProductionAST{
				Identifier: "initial",
				Rules: []RuleAST{
					{
						Expressions: []Expression{
							StringLiteralAST{
								Value: "\"Hello, World!\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Literal identifier",
			input: "initial = World",
			output: ProductionAST{
				Identifier: "initial",
				Rules: []RuleAST{
					{
						Expressions: []Expression{
							IdentifierAST{
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:  "Literal grouped identifier",
			input: "initial = (World)",
			output: ProductionAST{
				Identifier: "initial",
				Rules: []RuleAST{
					{
						Expressions: []Expression{
							GroupingAST{
								Expression: IdentifierAST{
									Value: "World",
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Production with multiple rules",
			input: "initial = \"Hello, World!\" | World | (\"No\")",
			output: ProductionAST{
				Identifier: "initial",
				Rules: []RuleAST{
					{
						Expressions: []Expression{
							StringLiteralAST{
								Value: "\"Hello, World!\"",
							},
						},
					},
					{
						Expressions: []Expression{
							IdentifierAST{
								Value: "World",
							},
						},
					},
					{
						Expressions: []Expression{
							GroupingAST{
								Expression: StringLiteralAST{
									Value: "\"No\"",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		lexer := LexString(test.input)
		parser := NewParser(lexer)

		actual, err := parser.production()
		if err != nil {
			t.Error(test.name, "- Error", err)
		}

		if !reflect.DeepEqual(actual, test.output) {
			t.Error(test.name, "- Expression", actual, "expected", test.output)
		}

		if _, err := parser.production(); !errors.Is(err, ErrUnexpectedEOF) {
			t.Error(test.name, "- error parsing past EOF was incorrect:", err)
		}
	}
}
