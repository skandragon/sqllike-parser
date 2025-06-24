package ast

import "github.com/skandragon/sqllike-parser/lexer"

type NumberExpr struct {
	Value float64
}

func (n *NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (s *StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (s *SymbolExpr) expr() {}

// complex expressions

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (b *BinaryExpr) expr() {}

type UnaryExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (u *UnaryExpr) expr() {}

type CallExpr struct {
	Function string
	Args     []Expr
}

func (c *CallExpr) expr() {}
