package parser

import (
	"github.com/skandragon/sqllike-parser/ast"
	"github.com/skandragon/sqllike-parser/lexer"
)

type bindingPower int

const (
	defaultBP bindingPower = iota
	comma
	logical
	relational
	additive
	multiplicative
	unary
	call
	primary
)

type stmtHandler func(p *parser) (stmt ast.Stmt)
type nudHandler func(p *parser) (expr ast.Expr)
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) (expr ast.Expr)

type stmtLookup map[lexer.TokenKind]stmtHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

var bp_lu = bpLookup{}
var nud_lu = nudLookup{}
var led_lu = ledLookup{}
var stmt_lu = stmtLookup{}

func led(kind lexer.TokenKind, bp bindingPower, fn ledHandler) {
	if _, ok := led_lu[kind]; ok {
		panic("led already registered for " + kind.String())
	}
	led_lu[kind] = fn
	bp_lu[kind] = bp
}

func nud(kind lexer.TokenKind, fn nudHandler) {
	if _, ok := nud_lu[kind]; ok {
		panic("nud already registered for " + kind.String())
	}
	nud_lu[kind] = fn
}

func stmt(kind lexer.TokenKind, fn stmtHandler) {
	if _, ok := stmt_lu[kind]; ok {
		panic("stmt already registered for " + kind.String())
	}
	stmt_lu[kind] = fn
	bp_lu[kind] = defaultBP
}

func createTokenLookups() {
	nud(lexer.TokenNumber, parsePrimaryExpr)
	nud(lexer.TokenString, parsePrimaryExpr)
	nud(lexer.TokenIdentifier, parsePrimaryExpr)
	nud(lexer.TokenMinus, parseUnaryExpr)
	nud(lexer.TokenOpenParen, parseGroupingExpr)

	// Logical
	led(lexer.TokenKeywordAnd, logical, parseBinaryExpr)
	led(lexer.TokenKeywordOr, logical, parseBinaryExpr)

	// Relational
	led(lexer.TokenEqual, relational, parseBinaryExpr)
	led(lexer.TokenNotEqual, relational, parseBinaryExpr)
	led(lexer.TokenLessThan, relational, parseBinaryExpr)
	led(lexer.TokenGreaterThan, relational, parseBinaryExpr)
	led(lexer.TokenLessThanOrEqual, relational, parseBinaryExpr)
	led(lexer.TokenGreaterThanOrEqual, relational, parseBinaryExpr)
	led(lexer.TokenKeywordLike, relational, parseBinaryExpr)

	// Additive and Multiplicative
	led(lexer.TokenPlus, additive, parseBinaryExpr)
	led(lexer.TokenMinus, additive, parseBinaryExpr)
	led(lexer.TokenAsterisk, multiplicative, parseBinaryExpr)
	led(lexer.TokenSlash, multiplicative, parseBinaryExpr)
	led(lexer.TokenPercent, multiplicative, parseBinaryExpr)

}
