package parser

import (
	"github.com/skandragon/sqllike-parser/ast"
	"github.com/skandragon/sqllike-parser/lexer"
)

func parseStatement(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentToken().Kind]
	if exists {
		return stmt_fn(p)
	}

	expression := parseExpr(p, defaultBP)
	p.expect(lexer.TokenSemicolon)

	return &ast.ExpressionStmt{
		Expr: expression,
	}
}
