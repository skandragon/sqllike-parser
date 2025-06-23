package lexer

import (
	"fmt"
	"regexp"
)

type lexer struct {
	patterns []regexPatterns
	Tokens   []Token
	source   string
	pos      int
}

type regexHandler func(l *lexer, regex *regexp.Regexp, match []int)

type regexPatterns struct {
	regex   *regexp.Regexp
	handler regexHandler
}

func Tokenize(source string) []Token {
	lex := newLexer(source)

	for !lex.atEOF() {
		matched := false
		for _, pattern := range lex.patterns {
			match := pattern.regex.FindStringIndex(lex.remainder())
			if match != nil && match[0] == 0 {
				pattern.handler(lex, pattern.regex, match)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("no regex matched at position %d: %s", lex.pos, lex.remainder()))
		}
	}

	lex.push(NewToken(TokenEOF, "EOF"))
	return lex.Tokens
}

func newLexer(source string) *lexer {
	return &lexer{
		source: source,
		patterns: []regexPatterns{
			{regexp.MustCompile(`"[_a-zA-Z]+[_a-zA-Z0-9\.]*"`), identifierHandler},
			{regexp.MustCompile(`[_a-zA-Z]+[_a-zA-Z0-9\.]*`), symbolHandler},
			{regexp.MustCompile(`\s+`), skupHandler},
			{regexp.MustCompile(`'([^']*)'`), stringHandler},
			{regexp.MustCompile(`\d+(\.\d+)?`), numberHandler},

			{regexp.MustCompile(`\(`), defaultHandler(TokenOpenParen, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(TokenCloseParen, ")")},
			{regexp.MustCompile(`,`), defaultHandler(TokenComma, ",")},
			{regexp.MustCompile(`;`), defaultHandler(TokenSemicolon, ";")},
			{regexp.MustCompile(`\?`), defaultHandler(TokenQuestionMark, "?")},
			{regexp.MustCompile(`\+`), defaultHandler(TokenPlus, "+")},
			{regexp.MustCompile(`-`), defaultHandler(TokenMinus, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(TokenAsterisk, "*")},
			{regexp.MustCompile(`/`), defaultHandler(TokenSlash, "/")},
			{regexp.MustCompile(`%`), defaultHandler(TokenPercent, "%")},
			{regexp.MustCompile(`=`), defaultHandler(TokenEqual, "=")},
			{regexp.MustCompile(`>=`), defaultHandler(TokenGreaterThanOrEqual, ">=")},
			{regexp.MustCompile(`<=`), defaultHandler(TokenLessThanOrEqual, "<=")},
			{regexp.MustCompile(`>`), defaultHandler(TokenGreaterThan, ">")},
			{regexp.MustCompile(`<`), defaultHandler(TokenLessThan, "<")},
			{regexp.MustCompile(`<>`), defaultHandler(TokenNotEqual, "<>")},
		},
	}
}

func (l *lexer) advanceN(n int) {
	if n < 0 || l.pos+n > len(l.source) {
		panic("advanceN out of bounds")
	}
	l.pos += n
}

func (l *lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *lexer) at() byte {
	if l.pos >= len(l.source) {
		return 0 // EOF
	}
	return l.source[l.pos]
}

func (l *lexer) atEOF() bool {
	return l.pos >= len(l.source)
}

func (l *lexer) remainder() string {
	if l.pos >= len(l.source) {
		return ""
	}
	return l.source[l.pos:]
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp, _ []int) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func skupHandler(l *lexer, regex *regexp.Regexp, match []int) {
	l.advanceN(match[1])
}

func numberHandler(l *lexer, regex *regexp.Regexp, match []int) {
	l.push(NewToken(TokenNumber, l.remainder()[:match[1]]))
	l.advanceN(match[1])
}

func stringHandler(l *lexer, regex *regexp.Regexp, match []int) {
	l.push(NewToken(TokenString, l.remainder()[1:match[1]-1])) // Exclude quotes
	l.advanceN(match[1])
}

func symbolHandler(l *lexer, regex *regexp.Regexp, match []int) {
	value := l.remainder()[:match[1]]
	kind := TokenIdentifier

	if keyword, exists := keyword_lu[value]; exists {
		kind = keyword
	}

	l.push(NewToken(kind, value))
	l.advanceN(match[1])
}

func identifierHandler(l *lexer, regex *regexp.Regexp, match []int) {
	l.push(NewToken(TokenString, l.remainder()[1:match[1]-1])) // Exclude quotes
	l.advanceN(match[1])
}
