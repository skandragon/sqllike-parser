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
	return p.pos < len(p.tokens) && p.tokens[p.pos].Kind() != lexer.TokenEOF
}

func parsePrimaryExpr(p *parser) ast.Expr {
	token := p.advance()
	switch token.Kind() {
	case lexer.TokenNumber:
		v, err := strconv.ParseFloat(token.Value(), 64)
		if err != nil {
			panic("invalid number: " + token.Value())
		}
		return &ast.NumberExpr{Value: v}
	case lexer.TokenString:
		return &ast.StringExpr{Value: token.Value()}
	case lexer.TokenIdentifier:
		return &ast.SymbolExpr{Value: token.Value()}
	default:
		panic("invalid type for parsePrimaryExpr: " + token.String())
	}
}
