package parser

import (
	"errors"
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/lexer"
	"github.com/LucaScorpion/keyScripter/internal/runtime"
)

type parser struct {
	tokens   []*lexer.Token
	tokenPos int
	ctx      *runtime.Context
}

func Parse(input string) (*runtime.Script, error) {
	if tokens, err := lexTokens(input); err != nil {
		return nil, err
	} else {
		p := &parser{
			tokens: tokens,
			ctx:    runtime.NewContext(),
		}

		if instr, err := p.parseInstructions(); err != nil {
			return nil, err
		} else {
			return runtime.NewScript(instr), nil
		}
	}
}

// TODO: Move this logic to the lexer.
func lexTokens(input string) ([]*lexer.Token, error) {
	// Create and run the lexer.
	l := lexer.NewLexer(input)
	go l.Run()

	// Collect all the tokens.
	var tokens []*lexer.Token
	for t := range l.Tokens {
		tokens = append(tokens, t)
		if t.TokenType == lexer.TokenError {
			return nil, errors.New(t.Value)
		}
	}

	return tokens, nil
}

// Peek the nth next token in tokens.
// If there are no more tokens, a lexer.TokenEOF is returned.
func (p *parser) peekTokenN(n int) *lexer.Token {
	if p.tokenPos+n >= len(p.tokens) {
		return &lexer.Token{
			TokenType: lexer.TokenEOF,
			Value:     "",
		}
	}
	return p.tokens[p.tokenPos+n]
}

// Peek the next token in tokens.
func (p *parser) peekToken() *lexer.Token {
	return p.peekTokenN(1)
}

// Read the next token in tokens and increment tokenPos.
func (p *parser) readToken() *lexer.Token {
	next := p.peekToken()
	p.tokenPos++
	return next
}

func (p *parser) expectToken(t lexer.TokenType) (*lexer.Token, error) {
	if next := p.readToken(); next.TokenType == t {
		return next, nil
	} else {
		return nil, fmt.Errorf("unexpected token, expected %s but got %s", lexer.TokenName(t), next.Name())
	}
}
