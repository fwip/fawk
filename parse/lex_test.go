package parse

import "testing"
import "os"

var shouldLex = [...]string{
	`{print}`,
	` /Err\/die/ `,
	` { x += $2; $3 = lshift($1, 2) ; print } `,
	` { " Hi I'm \" georigiono" } `,
	`NF%2 == 1`,
	`{ $2 = $3 == 4 ? "true" : 0 }`,
	` /my pattern/ {print "Found it!", $2 $3 } `,
	`BEGIN {OFS="\t"} {x+=$3} END{print x}`,
}

func lexes(s string, t *testing.T) {
	l := lex(s, os.Stderr)
	n := l.nextItem()
	for ; n.yys != int(itemEOF); n = l.nextItem() {
		t.Logf("%-20s %#v\n", n, n)
		if n.yys == int(itemError) {
			t.Error(n)
		}
	}
	t.Log(n)
}

func TestLexing(t *testing.T) {
	for _, s := range shouldLex {
		t.Run("ShouldLex"+s, func(t *testing.T) {
			lexes(s, t)
		})
	}
}
