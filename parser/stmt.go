package parser

import (
	"strconv"

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

func parseSelectStmt(p *parser) ast.Stmt {
	p.expect(lexer.TokenKeywordSelect)
	columns := []ast.Expr{}

	if p.currentToken().Kind == lexer.TokenAsterisk {
		columns = append(columns, &ast.SymbolExpr{Value: "*"})
		p.advance()
	} else {
		columns = parseExprList(p)
	}

	p.expect(lexer.TokenKeywordFrom)
	table := p.expect(lexer.TokenIdentifier)

	var whereClause ast.Expr
	if p.currentToken().Kind == lexer.TokenKeywordWhere {
		p.advance()
		whereClause = parseExpr(p, defaultBP)
	}

	var groupBy []ast.Expr
	if p.currentToken().Kind == lexer.TokenKeywordGroup {
		p.advance()
		p.expect(lexer.TokenKeywordBy)
		groupBy = parseExprList(p)
	}

	limit := 0
	if p.currentToken().Kind == lexer.TokenKeywordLimit {
		p.advance()
		limitToken := p.expect(lexer.TokenNumber)
		var err error
		limit, err = strconv.Atoi(limitToken.Value)
		if err != nil {
			panic("invalid limit value: " + limitToken.Value)
		}
	}

	p.expect(lexer.TokenSemicolon)

	return &ast.SelectStmt{
		Columns: columns,
		Table:   table.Value,
		Where:   whereClause,
		GroupBy: groupBy,
		Limit:   limit,
	}
}

// parseExprList parses expr (',' expr)* and returns the slice.
func parseExprList(p *parser) []ast.Expr {
	var list []ast.Expr
	for {
		list = append(list, parseExpr(p, defaultBP))
		if p.currentToken().Kind != lexer.TokenComma {
			break
		}
		p.advance() // skip comma
	}
	return list
}

func parseFunctionCall(p *parser, left ast.Expr, _ bindingPower) ast.Expr {
	// Ensure the left is a valid identifier
	if _, ok := left.(*ast.SymbolExpr); !ok {
		panic("function call must start with an identifier")
	}
	funcName := left.(*ast.SymbolExpr).Value
	p.advance()

	var args []ast.Expr
	for p.currentToken().Kind != lexer.TokenCloseParen {
		switch p.currentToken().Kind {
		case lexer.TokenAsterisk:
			// count(*)
			p.advance()
			args = append(args, &ast.SymbolExpr{Value: "*"})
		default:
			// sum(x+1), max(a,b,c), etc.
			args = append(args, parseExpr(p, defaultBP))
		}

		// comma-separated
		if p.currentToken().Kind == lexer.TokenComma {
			p.advance()
			continue
		}
		break
	}

	// consume the ‘)’
	p.expect(lexer.TokenCloseParen)

	return &ast.CallExpr{
		Function: funcName,
		Args:     args,
	}
}
