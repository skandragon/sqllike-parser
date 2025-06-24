package main

import (
	"fmt"

	"github.com/sanity-io/litter"

	"github.com/skandragon/sqllike-parser/lexer"
	"github.com/skandragon/sqllike-parser/parser"
)

// --- end of Ragel block --------------------------------

func main() {
	stmt := `SELECT resource.name, sum(_cardinalhq.value) FROM metrics WHERE frequency=10000 AND "_cardinalhq.name" = 'k8s.cpu.usage' AND resource.cluster.name = 'prod-cluster' GROUP BY resource.name LIMIT 100;`
	//stmt := `45.2 + (5 * 4 -5);`
	tokens := lexer.Tokenize(stmt)
	for _, token := range tokens {
		fmt.Printf("%s\n", token.DebugString())
	}

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
