package parser

import (
	"errors"
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/lexer"
	"github.com/LucaScorpion/keyScripter/internal/runtime"
	"strconv"
)

type parser struct {
	input    string
	tokens   []*lexer.Token
	tokenPos int
	funcs    []*runtime.RuntimeFn
}

func Parse(input string) (*Script, error) {
	p := &parser{input: input}

	if err := p.lexTokens(); err != nil {
		return nil, err
	}

	if err := p.prepareFuncs(); err != nil {
		return nil, err
	}

	return &Script{funcs: p.funcs}, nil
}

func (p *parser) lexTokens() error {
	// Collect all the tokens.
	l := lexer.NewLexer(p.input)
	go l.Run()

	var tokens []*lexer.Token

	for t := range l.Tokens {
		tokens = append(tokens, t)
		if t.TokenType == lexer.TokenError {
			return errors.New(t.Value)
		}
	}

	p.tokens = tokens
	return nil
}

func (p *parser) prepareFuncs() error {
	for ; p.tokenPos < len(p.tokens); p.tokenPos++ {
		token := p.tokens[p.tokenPos]
		switch token.TokenType {
		case lexer.TokenFuncName:
			if err := p.prepareFunc(); err != nil {
				return err
			}
		case lexer.TokenEOF:
		case lexer.TokenComment:
		default:
			return fmt.Errorf("unexpected token: %s", lexer.TokenNames[token.TokenType])
		}
	}

	return nil
}

func (p *parser) prepareFunc() error {
	// Get the function name, look it up.
	funcName := p.tokens[p.tokenPos].Value
	fn, ok := runtime.Functions[funcName]
	if !ok {
		return fmt.Errorf("unknown function \"%s\"", funcName)
	}

	// Collect all function argument tokens.
	var argValues []interface{}
	nextToken := p.peekToken()
	for nextToken != nil {
		t := nextToken.TokenType
		if t == lexer.TokenLiteralString {
			argValues = append(argValues, nextToken.Value)
		} else if t == lexer.TokenLiteralInt {
			i, err := strconv.Atoi(nextToken.Value)
			if err != nil {
				return fmt.Errorf("%s is not a valid number", nextToken.Value)
			}
			argValues = append(argValues, i)
		} else {
			break
		}

		p.tokenPos++
		nextToken = p.peekToken()
	}

	// Prepare the function.
	prepared, err := fn(argValues)
	if err != nil {
		return err
	}

	p.funcs = append(p.funcs, &prepared)
	return nil
}

func (p *parser) peekToken() *lexer.Token {
	if p.tokenPos+1 >= len(p.tokens) {
		return nil
	}
	return p.tokens[p.tokenPos+1]
}
