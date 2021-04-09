package parser

import (
	"fmt"
	"github.com/LucaScorpion/keyScripter/internal/lexer"
	"github.com/LucaScorpion/keyScripter/internal/runtime"
	"strconv"
)

func (p *parser) parseInstructions() ([]runtime.Instruction, error) {
	var instr []runtime.Instruction
	for nextToken := p.peekToken(); nextToken.TokenType != lexer.TokenEOF; nextToken = p.peekToken() {
		if i, err := p.parseInstruction(); err != nil {
			return nil, err
		} else if i != nil {
			instr = append(instr, i)
		}
	}
	return instr, nil
}

func (p *parser) parseInstruction() (runtime.Instruction, error) {
	nextToken := p.peekToken()
	switch nextToken.TokenType {
	case lexer.TokenIdentifier:
		if p.peekTokenN(2).TokenType == lexer.TokenAssign {
			return p.parseAssignment()
		} else {
			return p.parseFuncCall()
		}
	case lexer.TokenComment:
		fallthrough
	case lexer.TokenNewline:
		p.readToken()
		return nil, nil
	default:
		return nil, fmt.Errorf("unexpected token: %s", nextToken.Name())
	}
}

func (p *parser) parseAssignment() (runtime.Instruction, error) {
	varName := p.readToken().Value

	// Check if the name is not a function name.
	if _, ok := runtime.Functions[varName]; ok {
		return nil, fmt.Errorf("cannot assign to \"%s\"", varName)
	}

	// Consume the assign token.
	p.readToken()

	// Get the assigned value.
	val, err := p.parseValue()
	if err != nil {
		return nil, err
	}

	// Store the value in the context to keep track of the kind.
	p.ctx.SetValue(varName, val.Resolve(p.ctx))

	// An assignment must be followed by a newline or EOF.
	endToken := p.readToken()
	if endToken.TokenType != lexer.TokenNewline && endToken.TokenType != lexer.TokenEOF {
		return nil, fmt.Errorf("unexpected %s token after assignment", endToken.Name())
	}

	return runtime.Assignment{
		Name: varName,
		Val:  val,
	}, nil
}

func (p *parser) parseFuncCall() (runtime.Instruction, error) {
	// Get the function name, look it up.
	funcName := p.readToken().Value
	fn, ok := runtime.Functions[funcName]
	if !ok {
		return nil, fmt.Errorf("unknown function \"%s\"", funcName)
	}

	// Collect all function argument tokens.
	var argValues []runtime.Value
	for next := p.peekToken(); next.TokenType != lexer.TokenEOF && next.TokenType != lexer.TokenNewline; next = p.peekToken() {
		if next.IsValueToken() {
			val, err := p.parseValue()
			if err != nil {
				return nil, err
			}
			argValues = append(argValues, val)
		} else if next.TokenType == lexer.TokenComment {
			p.readToken()
		} else {
			return nil, fmt.Errorf("unexpected token in function call: %s", next.Name())
		}
	}

	// Validate the function.
	argKinds := make([]runtime.Kind, len(argValues))
	for i := 0; i < len(argValues); i++ {
		argKinds[i] = argValues[i].Resolve(p.ctx).Kind()
	}
	if err := fn.Validate(argKinds); err != nil {
		return nil, err
	}

	// Store the function call instruction.
	return runtime.FunctionCall{
		Fn:   fn,
		Args: argValues,
	}, nil
}

func (p *parser) parseValue() (runtime.Value, error) {
	valueToken := p.readToken()
	switch valueToken.TokenType {
	case lexer.TokenLiteralString:
		str, err := processString(valueToken.Value)
		if err != nil {
			return nil, err
		}
		return runtime.NewStringValue(str), nil
	case lexer.TokenLiteralInt:
		intVal, err := strconv.Atoi(valueToken.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid number value: %s", valueToken.Value)
		}
		return runtime.NewNumberValue(intVal), nil
	case lexer.TokenLiteralHex:
		intVal, err := strconv.ParseInt(valueToken.Value, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid number value: %s", valueToken.Value)
		}
		return runtime.NewNumberValue(int(intVal)), nil
	case lexer.TokenIdentifier:
		// Check if the referenced value is defined in the current context.
		if !p.ctx.HasValue(valueToken.Value) {
			return nil, fmt.Errorf("undefined value: %s", valueToken.Value)
		}
		return runtime.NewVariableValue(valueToken.Value), nil
	case lexer.TokenParenOpen:
		return p.parseFuncDef()
	default:
		return nil, fmt.Errorf("unexpected token as value: %s", valueToken.Name())
	}
}

func (p *parser) parseFuncDef() (runtime.Value, error) {
	// The opening paren is already read.
	// Get the parameter names.
	var paramNames []string
	var next *lexer.Token
	for next = p.peekToken(); next.TokenType == lexer.TokenIdentifier; next = p.peekToken() {
		// Consume the token, store the name.
		p.readToken()
		paramNames = append(paramNames, next.Value)
	}

	// End of parameters, start of function body.
	if _, err := p.expectToken(lexer.TokenParenClose); err != nil {
		return nil, err
	}
	if _, err := p.expectToken(lexer.TokenBraceOpen); err != nil {
		return nil, err
	}

	// TODO
	return nil, nil
}
