package parser

import (
	"strconv"

	"github.com/skandragon/sqllike-parser/ast"
	"github.com/skandragon/sqllike-parser/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	body := make([]ast.Stmt, 0)

	createTokenLookups()

	p := &parser{
		tokens: tokens,
		pos:    0,
	}

	for p.hasTokens() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.tokens[p.pos].Kind != lexer.TokenEOF
}

func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	if !p.hasTokens() {
		panic("unexpected end of input")
	}
	token := p.currentToken()
	if token.Kind != expectedKind {
		if err == nil {
			err = "expected " + expectedKind.String() + " but found " + token.String() + " at position " + strconv.Itoa(p.pos)
		}
		panic(err)
	}
	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
