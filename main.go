package main

import (
	"fmt"

	"github.com/skandragon/sqllike-parser/lexer"
)

// --- end of Ragel block --------------------------------

func main() {
	//stmt := `SELECT * FROM metrics WHERE frequency=10000 AND "_cardinalhq.name" = 'k8s.cpu.usage' AND resource.cluster.name = 'prod-cluster'`
	stmt := `45.2 + 5 * 4`
	tokens := lexer.Tokenize(stmt)
	for _, token := range tokens {
		fmt.Printf("%s\n", token.DebugString())
	}
}
