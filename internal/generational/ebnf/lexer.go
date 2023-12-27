package ebnf

import (
	"errors"
	"io"
	"strings"
	"unicode"

	"github.com/brandhoej/cuzz/internal/collections"
)

var ErrInvalidToken = errors.New("invalid token")

const (
	// Single-character tokens.
	LeftParenthesis = iota
	RightParenthesis
	LeftSquareBracket
	RightSquareBracket
	Pipe
	Dot
	Equal

	// Literal tokens.
	Identifier
	String

	EOF
)

type Token struct {
	Class  int
	Lexeme string
}

type lookahead[T any] struct {
	value T
	err   error
}

func token(class int, lexeme string) Token {
	return Token{
		Class:  class,
		Lexeme: lexeme,
	}
}

type Lexer struct {
	reader     io.RuneReader
	lexeme     []rune
	lookaheads collections.Queue[lookahead[rune]]
}

func LexString(input string) Lexer {
	return Lexer{
		reader:     strings.NewReader(input),
		lexeme:     make([]rune, 0),
		lookaheads: collections.NewArrayQueue[lookahead[rune]](),
	}
}

func (lexer *Lexer) clear() {
	lexer.lexeme = lexer.lexeme[:0]
}

func (lexer *Lexer) peek() (rune, error) {
	if lexer.lookaheads.Size() > 0 {
		lookahead := lexer.lookaheads.Peek()
		return lookahead.value, lookahead.err
	}

	character, _, err := lexer.reader.ReadRune()
	lexer.lookaheads.Enqueue(lookahead[rune]{
		character, err,
	})

	return character, err
}

func (lexer *Lexer) advance() (rune, error) {
	// Peek will first check lookahead, if empty it will read the next and Enqueuee it.
	character, err := lexer.peek()
	lexer.lookaheads.Dequeue()

	if err == nil {
		lexer.lexeme = append(lexer.lexeme, character)
	}
	return character, err
}

func (lexer *Lexer) skipSpaces() (rune, error) {
	for {
		character, err := lexer.advance()
		if !unicode.IsSpace(character) && err == nil {
			return character, err
		}

		if err == nil {
			// Remove the space character from the lexeme.
			lexer.lexeme = lexer.lexeme[:len(lexer.lexeme)-1]
		}

		if err != nil {
			return character, err
		}
	}
}

func (lexer *Lexer) token(class int) Token {
	return Token{
		Class:  class,
		Lexeme: string(lexer.lexeme),
	}
}

func (lexer *Lexer) Next() (Token, error) {
	defer lexer.clear()

	character, err := lexer.skipSpaces()
	if err == io.EOF {
		return lexer.token(EOF), nil
	}

	switch character {
	case '(':
		return lexer.token(LeftParenthesis), nil
	case ')':
		return lexer.token(RightParenthesis), nil
	case '[':
		return lexer.token(LeftSquareBracket), nil
	case ']':
		return lexer.token(RightSquareBracket), nil
	case '|':
		return lexer.token(Pipe), nil
	case '=':
		return lexer.token(Equal), nil
	case '.':
		return lexer.token(Dot), nil
	}

	if lexer.identifier(character) {
		return lexer.token(Identifier), nil
	}

	if lexer.string(character) {
		return lexer.token(String), nil
	}

	return Token{}, ErrInvalidToken
}

func (lexer *Lexer) string(character rune) bool {
	if character != '"' {
		return false
	}

	// Read until next '"' rune.
	for {
		character, err := lexer.advance()
		// We most likely encountered EOF before a '"'.
		if err != nil {
			return false
		}

		if character == '"' {
			break
		}
	}

	return true
}

func (lexer *Lexer) identifier(character rune) bool {
	if !unicode.IsLetter(character) {
		return false
	}

	// Read until whitespace or none letter or none number.
	for {
		character, err := lexer.peek()

		if err == io.EOF {
			return true
		}

		if err != nil {
			return false
		}

		if unicode.IsSpace(character) ||
			!(unicode.IsLetter(character) ||
				unicode.IsNumber(character) ||
				character == '_') {
			return true
		}

		lexer.advance()
	}
}
