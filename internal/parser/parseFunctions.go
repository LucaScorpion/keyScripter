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

func (p *parser) parseInstructionBlock() ([]runtime.Instruction, error) {
	var instr []runtime.Instruction
	for nextToken := p.peekToken(); nextToken.TokenType != lexer.TokenEOF && nextToken.TokenType != lexer.TokenBlockEnd; nextToken = p.peekToken() {
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
	case lexer.TokenNewline:
		p.readToken()
		return nil, nil
	case lexer.TokenTimestamps:
		return p.parseTimestamps()
	default:
		return nil, fmt.Errorf("unexpected token: %s", nextToken.Name())
	}
}

func (p *parser) parseAssignment() (runtime.Instruction, error) {
	// Name token, assign token.
	varName := p.readToken().Value
	p.readToken()

	// Get the assigned value.
	val, err := p.parseValue()
	if err != nil {
		return nil, err
	}

	// Store the value in the context to keep track of the kind.
	p.ctx.SetValue(varName, val.Resolve(p.ctx))

	// An assignment must be followed by a newline, comment, or EOF.
	endToken := p.readToken()
	if endToken.TokenType != lexer.TokenNewline && endToken.TokenType != lexer.TokenEOF {
		return nil, fmt.Errorf("unexpected %s token after assignment", endToken.Name())
	}

	return runtime.NewAssignment(varName, val), nil
}

func (p *parser) parseFuncCall() (runtime.Instruction, error) {
	// Get the function name, check if it exists and is a function.
	funcName := p.readToken().Value
	if !p.ctx.HasValue(funcName) {
		return nil, fmt.Errorf("unknown function: %s", funcName)
	}
	funcVal := p.ctx.GetValue(funcName)
	if funcVal.Kind() != runtime.FunctionKind {
		return nil, fmt.Errorf("cannot call a %s as a function: %s", funcVal.Kind(), funcName)
	}

	// Collect all function argument tokens.
	var argValues []*runtime.Value
	for next := p.peekToken(); next.TokenType != lexer.TokenEOF && next.TokenType != lexer.TokenNewline; next = p.peekToken() {
		if next.IsValueToken() {
			val, err := p.parseValue()
			if err != nil {
				return nil, err
			}
			argValues = append(argValues, val)
		} else {
			return nil, fmt.Errorf("unexpected token in function call: %s", next.Name())
		}
	}

	// Get all argument kinds.
	argKinds := make([]runtime.Kind, len(argValues))
	for i := 0; i < len(argValues); i++ {
		argI := argValues[i]
		resolvedKind := argI.Resolve(p.ctx).Kind()

		// For "any" kind args, change their kind in context to match the function's param kind.
		if resolvedKind == runtime.AnyKind {
			resolvedKind = funcVal.ParamKind(i)
			p.ctx.SetValue(argI.RawValue().(string), runtime.NewEmptyValue(resolvedKind))
		}

		argKinds[i] = resolvedKind
	}

	// Check if the argument count matches.
	if (funcVal.Variadic() && len(argValues) < len(funcVal.ParamKinds())-1) || (!funcVal.Variadic() && len(argValues) != len(funcVal.ParamKinds())) {
		return nil, fmt.Errorf("mismatched argument count, expected %d but got %d", len(funcVal.ParamKinds()), len(argValues))
	}

	// Check if the argument kinds match.
	for argI, argKind := range argKinds {
		// Get the param index, to account for variadic parameters.
		paramI := argI
		if argI > len(funcVal.ParamKinds())-1 {
			paramI = len(funcVal.ParamKinds()) - 1
		}

		paramKind := funcVal.ParamKind(paramI)
		if paramKind != runtime.AnyKind && paramKind != argKind {
			return nil, fmt.Errorf("mismatched argument type, expected %s but got %s", paramKind, argKind)
		}
	}

	// Store the function call instruction.
	return runtime.NewFunctionCall(funcVal, argValues), nil
}

func (p *parser) parseValue() (*runtime.Value, error) {
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
		// Check if the referenced value is defined in the context.
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

func (p *parser) parseFuncDef() (*runtime.Value, error) {
	// The opening paren is already read.

	// Start a new parser context.
	p.ctx = runtime.NewContext(p.ctx)

	// Get the parameter names.
	var paramNames []string
	var next *lexer.Token
	for next = p.peekToken(); next.TokenType == lexer.TokenIdentifier; next = p.peekToken() {
		// Consume the token, store the name.
		p.readToken()
		paramNames = append(paramNames, next.Value)

		// Store the parameter in context as any kind.
		p.ctx.SetValue(next.Value, runtime.NewEmptyValue(runtime.AnyKind))
	}

	// End of parameters, start of function body.
	if _, err := p.expectToken(lexer.TokenParenClose); err != nil {
		return nil, err
	}
	if _, err := p.expectToken(lexer.TokenBlockStart); err != nil {
		return nil, err
	}

	// Function body.
	body, err := p.parseInstructionBlock()
	if err != nil {
		return nil, err
	}

	// Function body end.
	if _, err := p.expectToken(lexer.TokenBlockEnd); err != nil {
		return nil, err
	}

	// Get the parameter kinds from context.
	paramKinds := make([]runtime.Kind, len(paramNames))
	for i, name := range paramNames {
		paramKinds[i] = p.ctx.GetValue(name).Kind()
	}

	// Restore the parser context.
	p.ctx = p.ctx.Parent()

	return runtime.NewFunctionValue(runtime.NewRuntimeFunction(paramNames, body), paramKinds), nil
}

func (p *parser) parseTimestamps() (runtime.Instruction, error) {
	// Timestamps and block start.
	p.readToken()
	p.expectToken(lexer.TokenBlockStart)

	// Timestamps body.
	body, err := p.parseTimestampsBlock()
	if err != nil {
		return nil, err
	}

	// Block end.
	p.expectToken(lexer.TokenBlockEnd)

	return runtime.NewTimestamps(body), nil
}

func (p *parser) parseTimestampsBlock() ([]runtime.TimestampInstr, error) {
	var instr []runtime.TimestampInstr
	lastTimestamp := 0
	for nextToken := p.peekToken(); nextToken.TokenType != lexer.TokenEOF && nextToken.TokenType != lexer.TokenBlockEnd; nextToken = p.peekToken() {
		// Discard newlines.
		if nextToken.TokenType == lexer.TokenNewline {
			p.readToken()
			continue
		}

		// Timestamp.
		ms, err := p.expectToken(lexer.TokenLiteralInt)
		if err != nil {
			return nil, err
		}

		// Check if the timestamp is valid.
		msInt, err := strconv.Atoi(ms.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid number value: %s", ms.Value)
		}
		if msInt < lastTimestamp {
			return nil, fmt.Errorf("timestamps must be equal or ascending: %d -> %d", lastTimestamp, msInt)
		}
		lastTimestamp = msInt

		// Parse the instruction.
		if i, err := p.parseInstruction(); err != nil {
			return nil, err
		} else if i != nil {
			instr = append(instr, runtime.NewTimestampInstr(int64(msInt), i))
		}
	}
	return instr, nil
}
