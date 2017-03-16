package parse

import "testing"

var shouldParse = [...]string{
	`{print}`,
	` /Err\/die/ `,
	//` { x += $2; $3 = lshift($1, 2) ; print } `,
	` { " Hi I'm \" georigiono" } `,
	`NF%2 == 1`,
	//`{ $2 = $3 == 4 ? "true" : 0 }`,
	//` /my pattern/ {print "Found it!", $2 $3 } `,
	//`BEGIN {OFS="\t"} {x+=$3} END{print x}`,
	//	` ( $3 <= 4++ && $2 != 4 /= 3 || ! 6 > ($NF * 3 ^0 - 2 )) `,
	//	` # A comment and a newline
	//	`,
	//	` $4 ~ /a cool regex/ { print 0x12 }`,
	//	` $2 !~ /an awful regex/ { print 0x12 }`,
	//` { sub($8, "hat", "tophat") }`,
}

func TestParsing(t *testing.T) {
	for _, s := range shouldParse {
		t.Run("ShouldParse"+s, func(t *testing.T) {
			Parse(s)
		})
	}
}
