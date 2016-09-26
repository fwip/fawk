package parse

import "fmt"
import "testing"

var shouldLex = [...]string{
	`{print}`,
	` /Err\/die/ `,
	` { x += $2; $3 = $1>>2 ; print } `,
	` { " Hi I'm \" georigiono" } `,
	`NF%2 == 1`,
	`BEGIN {OFS="\t"} {x+=$3} END{print x}`,
	`{ $2 = $3 == 4 ? "true" : 0 }`,
}

func lexes(s string, t *testing.T) {
	l := lex(s)
	n := l.nextItem()
	for ; n.yys != int(itemEOF); n = l.nextItem() {
		t.Logf("%-20s %#v\n", n, n)
		if n.yys == int(itemError) {
			t.Error(n)
		}
	}
	t.Log(n)
}

func TestIt(t *testing.T) {
	for _, s := range shouldLex {
		t.Run("ShouldLex"+s, func(t *testing.T) {
			lexes(s, t)
		})
	}
	fmt.Println("done")
}
