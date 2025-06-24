package ast

type BlockStmt struct {
	Body []Stmt
}

func (b *BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expr Expr
}

func (e *ExpressionStmt) stmt() {}

type SelectStmt struct {
	Columns []Expr
	Table   string
	Where   Expr
	GroupBy []Expr
	Limit   int
}

func (s *SelectStmt) stmt() {}
