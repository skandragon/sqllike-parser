package parser

import (
	"strconv"

	"github.com/skandragon/sqllike-parser/ast"
	"github.com/skandragon/sqllike-parser/lexer"
)

func parsePrimaryExpr(p *parser) ast.Expr {
	token := p.advance()
	switch token.Kind {
	case lexer.TokenNumber:
		v, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			panic("invalid number: " + token.Value)
		}
		return &ast.NumberExpr{Value: v}
	case lexer.TokenString:
		return &ast.StringExpr{Value: token.Value}
	case lexer.TokenIdentifier:
		return &ast.SymbolExpr{Value: token.Value}
	default:
		panic("invalid type for parsePrimaryExpr: " + token.String())
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operatorToken := p.advance()
	right := parseExpr(p, bp)

	return &ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parseUnaryExpr(p *parser) ast.Expr {
	operatorToken := p.advance()
	right := parseExpr(p, unary)

	return &ast.UnaryExpr{
		Operator: operatorToken,
		Right:    right,
	}
}

func parseGroupingExpr(p *parser) ast.Expr {
	p.expectError(lexer.TokenOpenParen, "expected '(' for grouping expression")
	expr := parseExpr(p, defaultBP)
	p.expectError(lexer.TokenCloseParen, "expected ')' to close grouping expression")
	return expr
}

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	token := p.currentToken()
	nud_fn, exists := nud_lu[token.Kind]
	if !exists {
		panic("no nud handler for " + token.String())
	}
	left := nud_fn(p)

	for bp_lu[p.currentToken().Kind] > bp {
		led_fn, exists := led_lu[p.currentToken().Kind]
		if !exists {
			panic("no led handler for " + p.currentToken().String())
		}
		bp = bp_lu[p.currentToken().Kind]
		left = led_fn(p, left, bp)
	}
	return left
}
