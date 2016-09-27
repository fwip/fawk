package parse

import "os"

//go:generate go tool yacc -o parse.go -v parse.table parse.y
import "testing"

var shouldParse = [...]string{
	`{print}`,
	` NF%2 == 1`,
	` /my pattern/ {print "Found it!", $2 $3} `,
	`BEGIN {OFS="\t"}  {x+=$3} END{print x}`,
}

func parses(s string, t *testing.T) {
	l := lex(s, os.Stderr)
	yyErrorVerbose = true
	yyDebug = 5
	yyParse(l)
	if l.parserErrors > 0 {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	for _, s := range shouldParse {
		t.Run("ShouldParse '"+s+"'", func(t *testing.T) {
			parses(s, t)
		})
	}
}
