package parser

import (
	"errors"
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/lexer"
	"github.com/LucaScorpion/keyScripter/internal/runtime"
	"strconv"
)

type parser struct {
	input        string
	tokens       []*lexer.Token
	tokenPos     int
	instructions []runtime.Instruction
	ctx          *runtime.Context
}

func Parse(input string) (*runtime.Script, error) {
	p := &parser{
		input: input,
		ctx:   runtime.NewContext(),
	}

	if err := p.lexTokens(); err != nil {
		return nil, err
	}

	if err := p.parseTokens(); err != nil {
		return nil, err
	}

	return runtime.NewScript(p.instructions), nil
}

func (p *parser) lexTokens() error {
	// Create and run the lexer.
	l := lexer.NewLexer(p.input)
	go l.Run()

	// Collect all the tokens.
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

func (p *parser) parseTokens() error {
	for nextToken := p.peekToken(); nextToken.TokenType != lexer.TokenEOF; nextToken = p.peekToken() {
		switch nextToken.TokenType {
		case lexer.TokenIdentifier:
			afterNextToken := p.peekTokenN(2)
			if afterNextToken.TokenType == lexer.TokenAssign {
				if err := p.parseAssignment(); err != nil {
					return err
				}
				break
			}

			if err := p.parseFunc(); err != nil {
				return err
			}
		case lexer.TokenComment:
			p.readToken()
		case lexer.TokenNewline:
			p.readToken()
		default:
			return fmt.Errorf("unexpected token: %s", lexer.TokenNames[nextToken.TokenType])
		}
	}
	return nil
}

func (p *parser) parseAssignment() error {
	varName := p.readToken().Value

	// Check if the name is not a function name.
	if _, ok := runtime.Functions[varName]; ok {
		return fmt.Errorf("cannot assign to \"%s\"", varName)
	}

	// Consume the assign token.
	p.readToken()

	// Get the assigned value.
	val, err := p.parseValue()
	if err != nil {
		return err
	}

	// Store the assignment instruction.
	p.instructions = append(p.instructions, runtime.Assignment{
		Name: varName,
		Val:  val,
	})

	// Store the value in the context.
	p.ctx.SetValue(varName, val)

	// An assignment must be followed by a newline or EOF.
	endToken := p.readToken().TokenType
	if endToken != lexer.TokenNewline && endToken != lexer.TokenEOF {
		return fmt.Errorf("unexpected %s token after assignment", lexer.TokenNames[endToken])
	}
	return nil
}

func (p *parser) parseFunc() error {
	// Get the function name, look it up.
	funcName := p.readToken().Value
	fn, ok := runtime.Functions[funcName]
	if !ok {
		return fmt.Errorf("unknown function \"%s\"", funcName)
	}

	// Collect all function argument tokens.
	var argValues []runtime.Value
	for next := p.peekToken(); next.TokenType != lexer.TokenEOF && next.TokenType != lexer.TokenNewline; next = p.peekToken() {
		if next.IsValueToken() {
			val, err := p.parseValue()
			if err != nil {
				return err
			}
			argValues = append(argValues, val)
		} else if next.TokenType == lexer.TokenComment {
			p.readToken()
		} else {
			return fmt.Errorf("unexpected token in function call: %s", lexer.TokenNames[next.TokenType])
		}
	}

	// Validate the function.
	argKinds := make([]runtime.Kind, len(argValues))
	for i := 0; i < len(argValues); i++ {
		argKinds[i] = argValues[i].Kind(p.ctx)
	}
	if err := fn.Validate(argKinds); err != nil {
		return err
	}

	// Store the function call instruction.
	p.instructions = append(p.instructions, runtime.FunctionCall{
		Fn:   fn,
		Args: argValues,
	})
	return nil
}

func (p *parser) parseValue() (runtime.Value, error) {
	valueToken := p.readToken()
	switch valueToken.TokenType {
	case lexer.TokenLiteralString:
		// TODO: Process value, remove quotes, escape sequences.
		return runtime.NewStringValue(valueToken.Value), nil
	case lexer.TokenLiteralInt:
		intVal, err := strconv.Atoi(valueToken.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid number value: %s", valueToken.Value)
		}
		return runtime.NewNumberValue(intVal), nil
	case lexer.TokenIdentifier:
		return runtime.NewVariableValue(valueToken.Value), nil
	default:
		return nil, fmt.Errorf("unexpected token as value: %s", lexer.TokenNames[valueToken.TokenType])
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
