package parse

import "fmt"
import "testing"

func logLex(s string) {
	fmt.Printf("Lexing:\n%s\n\n", s)
	l := lex(s)
	n := l.nextItem()
	for ; n.yys != int(itemEOF); n = l.nextItem() {
		fmt.Printf("%-20s %#v\n", n, n)
	}
	fmt.Println(n)
}

func TestIt(t *testing.T) {
	logLex(`{print}`)
	logLex(` /Err\/die/ `)
	logLex(` { x += $2; $3 = $1>>2 ; print } `)
	logLex(` { " Hi I'm \" georigiono" } `)
	fmt.Println("done")
}
