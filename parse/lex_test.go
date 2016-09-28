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
	` ( $3 <= 4++ && $2 != 4 /= 3 || ! 6 > ($NF * 3 ^0 - 2 )) `,
	` # A comment and a newline
	`,
	` $4 ~ /a cool regex/ { print 0x12 }`,
	` $2 !~ /an awful regex/ { print 0x12 }`,
	` { sub($8, "hat", "tophat" }`,
}

var shouldNotLex = [...]string{
	`$3 | $4`,
	`$2 & 0x5`,
	` /this is an unfinished regex`,
}

func lexes(s string, t *testing.T) bool {
	l := lex(s, os.Stderr)
	n := l.nextItem()
	for ; n.yys != int(itemEOF); n = l.nextItem() {
		t.Logf("%-20s %#v\n", n, n)
		if n.yys == int(itemError) {
			return false
		}
	}
	t.Log(n)
	return true
}

func TestLexing(t *testing.T) {
	for _, s := range shouldLex {
		t.Run("ShouldLex"+s, func(t *testing.T) {
			if !lexes(s, t) {
				lexes(s, t)
			}
		})
	}

	for _, s := range shouldNotLex {
		t.Run("ShouldNotLex"+s, func(t *testing.T) {
			if lexes(s, t) {
				t.Error()
			}
		})
	}
}
