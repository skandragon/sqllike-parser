package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	TokenEOF TokenKind = iota
	TokenIdentifier
	TokenNumber
	TokenString
	TokenOperator

	TokenOpenParen
	TokenCloseParen
	TokenComma
	TokenSemicolon
	TokenQuestionMark
	TokenPlus
	TokenMinus
	TokenAsterisk
	TokenSlash
	TokenPercent
	TokenEqual
	TokenLessThan
	TokenGreaterThan
	TokenLessThanOrEqual
	TokenGreaterThanOrEqual
	TokenNotEqual

	TokenKeywordSelect
	TokenKeywordFrom
	TokenKeywordWhere
	TokenKeywordAnd
	TokenKeywordOr
	TokenKeywordNot
	TokenKeywordIn
	TokenKeywordLike
	TokenKeywordGroup
	TokenKeywordBy
)

var keyword_lu map[string]TokenKind = map[string]TokenKind{
	"SELECT": TokenKeywordSelect,
	"FROM":   TokenKeywordFrom,
	"WHERE":  TokenKeywordWhere,
	"AND":    TokenKeywordAnd,
	"OR":     TokenKeywordOr,
	"NOT":    TokenKeywordNot,
	"IN":     TokenKeywordIn,
	"LIKE":   TokenKeywordLike,
	"GROUP":  TokenKeywordGroup,
	"BY":     TokenKeywordBy,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func (t Token) isOneOf(kinds ...TokenKind) bool {
	return slices.Contains(kinds, t.Kind)
}

func (t TokenKind) String() string {
	switch t {
	case TokenEOF:
		return "EOF"
	case TokenIdentifier:
		return "IDENTIFIER"
	case TokenNumber:
		return "NUMBER"
	case TokenString:
		return "STRING"
	case TokenOperator:
		return "OPERATOR"
	case TokenOpenParen:
		return "OPEN_PAREN"
	case TokenCloseParen:
		return "CLOSE_PAREN"
	case TokenComma:
		return "COMMA"
	case TokenSemicolon:
		return "SEMICOLON"
	case TokenQuestionMark:
		return "QUESTION_MARK"
	case TokenPlus:
		return "PLUS"
	case TokenMinus:
		return "MINUS"
	case TokenAsterisk:
		return "ASTERISK"
	case TokenSlash:
		return "SLASH"
	case TokenPercent:
		return "PERCENT"
	case TokenEqual:
		return "EQUAL"
	case TokenLessThan:
		return "LESS_THAN"
	case TokenGreaterThan:
		return "GREATER_THAN"
	case TokenLessThanOrEqual:
		return "LESS_THAN_OR_EQUAL"
	case TokenGreaterThanOrEqual:
		return "GREATER_THAN_OR_EQUAL"
	case TokenNotEqual:
		return "NOT_EQUAL"
	case TokenKeywordSelect:
		return "KEYWORD_SELECT"
	case TokenKeywordFrom:
		return "KEYWORD_FROM"
	case TokenKeywordWhere:
		return "KEYWORD_WHERE"
	case TokenKeywordAnd:
		return "KEYWORD_AND"
	case TokenKeywordOr:
		return "KEYWORD_OR"
	case TokenKeywordNot:
		return "KEYWORD_NOT"
	case TokenKeywordIn:
		return "KEYWORD_IN"
	case TokenKeywordLike:
		return "KEYWORD_LIKE"
	case TokenKeywordGroup:
		return "KEYWORD_GROUP"
	case TokenKeywordBy:
		return "KEYWORD_BY"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", t)
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%s)", t.Kind.String(), t.Value)
}

func (t Token) DebugString() string {
	ts := t.String()
	return fmt.Sprintf("Token{kind: %s (%s)}", ts, t.Value)
}
