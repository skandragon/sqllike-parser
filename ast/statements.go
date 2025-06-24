package ast

type BlockStmt struct {
	Body []Stmt
}

func (b *BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expr Expr
}

func (e *ExpressionStmt) stmt() {}
