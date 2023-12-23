package ebnf

import "errors"

/* Grammar → (Production `Dot`)* EOF
 * Production  → `Identifier` `=` Rules
 * Rules       → Rule (`|` Rule)*
 * Rule        → Expression+
 * Expression  → `Identifier` | `String` | `(` Rule `)`
 *
 * Here the ``-encapsulation refers to token types
 *   and not string literals. */

var (
	ErrUnexpectedEOF                   = errors.New("unexpected end of input")
	ErrFailedParsingProduction         = errors.New("failed to parse production")
	ErrFailedParsingRules              = errors.New("failed to parse rules in production")
	ErrFailedParsingExpressionInRule   = errors.New("failed to parse expression in rule")
	ErrEmptyRule                       = errors.New("rule must have at least one expression")
	ErrMissingIdentifierInProduction   = errors.New("production is missing an identifier")
	ErrMissingEqualitySignInProduction = errors.New("production is missing an equality sign following its identifier")
)

type Expression interface {
}

type GrammarAST struct {
	Productions []ProductionAST
}

type ProductionAST struct {
	Identifier string
	Rules      []RuleAST
}

type StringLiteralAST struct {
	Value string
}

type IdentifierAST struct {
	Value string
}

type GroupingAST struct {
	Expression Expression
}

type RuleAST struct {
	Expressions []Expression
}

type Parser struct {
	lexer      Lexer
	lookaheads Queue[lookahead[Token]]
}

func NewParser(lexer Lexer) Parser {
	return Parser{
		lexer:      lexer,
		lookaheads: NewArrayQueue[lookahead[Token]](),
	}
}

func (parser *Parser) peek() (Token, error) {
	if parser.lookaheads.Size() > 0 {
		lookahead := parser.lookaheads.Peek()
		return lookahead.value, lookahead.err
	}

	token, err := parser.lexer.Next()
	parser.lookaheads.Enqueue(lookahead[Token]{
		token, err,
	})

	return token, err
}

func (parser *Parser) advance() (Token, error) {
	token, err := parser.peek()
	parser.lookaheads.Dequeue()
	return token, err
}

func (parser *Parser) skip() {
	parser.advance()
}

func (parser *Parser) match(class int) (bool, Token) {
	token, err := parser.peek()
	if err != nil {
		return false, token
	}

	return token.Class == class, token
}

func (parser *Parser) consume(class int) (bool, Token) {
	matched, token := parser.match(class)
	if matched {
		parser.skip()
	}

	return matched, token
}

func (parser *Parser) Grammar() (GrammarAST, error) {
	productions := make([]ProductionAST, 0)

	for {
		production, err := parser.production()
		if err != nil {
			return GrammarAST{}, errors.Join(ErrFailedParsingProduction, err)
		}

		productions = append(productions, production)

		if consumed, _ := parser.consume(Dot); !consumed {
			break
		}
	}

	return GrammarAST{
		Productions: productions,
	}, nil
}

func (parser *Parser) production() (ProductionAST, error) {
	if matched, _ := parser.match(EOF); matched {
		return ProductionAST{}, ErrUnexpectedEOF
	}

	consumed, token := parser.consume(Identifier)
	if !consumed {
		return ProductionAST{}, ErrMissingIdentifierInProduction
	}

	if consumed, _ := parser.consume(Equal); !consumed {
		return ProductionAST{}, ErrMissingEqualitySignInProduction
	}

	rules, err := parser.rules()
	if err != nil {
		return ProductionAST{}, errors.Join(ErrFailedParsingRules, err)
	}

	return ProductionAST{
		Identifier: token.Lexeme,
		Rules:      rules,
	}, nil
}

func (parser *Parser) rules() ([]RuleAST, error) {
	rules := make([]RuleAST, 0)
	for {
		rule, err := parser.rule()
		if err != nil {
			return nil, errors.Join(ErrFailedParsingRules, err)
		}

		rules = append(rules, rule)

		if consumed, _ := parser.consume(Pipe); !consumed {
			break
		}
	}

	return rules, nil
}

func (parser *Parser) rule() (RuleAST, error) {
	var expressions []Expression = make([]Expression, 0)

	for {
		expression, err := parser.expression()
		if err != nil {
			return RuleAST{}, errors.Join(ErrFailedParsingExpressionInRule, err)
		}
		if expression == nil {
			break
		}

		expressions = append(expressions, expression)
	}

	if len(expressions) == 0 {
		return RuleAST{}, ErrEmptyRule
	}

	return RuleAST{
		Expressions: expressions,
	}, nil
}

func (parser *Parser) expression() (Expression, error) {
	if consumed, token := parser.consume(String); consumed {
		return StringLiteralAST{Value: token.Lexeme}, nil
	}

	if consumed, token := parser.consume(Identifier); consumed {
		return IdentifierAST{Value: token.Lexeme}, nil
	}

	if consumed, _ := parser.consume(LeftParenthesis); consumed {
		expression, err := parser.expression()
		if consumed, _ := parser.consume(RightParenthesis); consumed {
			return GroupingAST{
				Expression: expression,
			}, err
		}
	}

	return nil, nil
}
