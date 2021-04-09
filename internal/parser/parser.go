package parser

import (
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
	lex := lexer.NewLexer(input)
	if err := lex.Run(); err != nil {
		return nil, err
	}

	p := &parser{
		tokens: lex.Tokens(),
		ctx:    runtime.NewContext(),
	}

	if instr, err := p.parseInstructions(); err != nil {
		return nil, err
	} else {
		return runtime.NewScript(instr), nil
	}
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
