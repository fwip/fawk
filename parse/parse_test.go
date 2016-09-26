package parse

//go:generate go tool yacc -o parse.go -v parse.table parse.y
import "testing"

func TestParse(t *testing.T) {
	parse(`{print}`)
	parse(` NF%2 == 1`)
}
