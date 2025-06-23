package ast

type NumberExpr struct {
	Value float64
}

func (n *NumberExpr) expr() {
}

type StringExpr struct {
	Value string
}

func (s *StringExpr) expr() {
}

type SymbolExpr struct {
	Value string
}

func (s *SymbolExpr) expr() {
}

// complex expressions

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (b *BinaryExpr) expr() {
}
