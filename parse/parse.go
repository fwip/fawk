package parse

import (
	"fmt"
	"os"
	"strings"
)

// A Node is an element of an AST
type Node interface {
	String() string
	//RawValue() string
	//Position() Pos
	//Type() itemType
	//Children() []Node
}

type emptyNode struct{}

func (e emptyNode) String() string { return "" }

type simpleNode struct {
	item item
}

func (n *simpleNode) String() string {
	if n == nil {
		return ""
	}

	return n.item.val
}

type program struct {
	rules []rule
}

func (p *program) String() (out string) {
	var lines []string

	for _, rule := range p.rules {
		lines = append(lines, rule.String())
	}
	return strings.Join(lines, "\n")
}

type rule struct {
	pattern Node
	action  Node
}

func (r *rule) String() string {
	if r == nil {
		return "{ print }"
	}
	fmt.Println("Printing a rule...")
	return r.pattern.String() + "{" + r.action.String() + "}"
}

type parser struct {
	root      Node
	lexer     *lexer
	curToken  item
	peekToken item
}

func (p *parser) parse() error {

	p.next()
	program := program{}

	for p.next(); p.curToken.typ != itemEOF; {
		program.rules = append(program.rules, p.parseRule())
		fmt.Println("rule added")
	}
	p.root = &program
	return nil
}

func (p *parser) next() {
	if p.curToken.typ == itemEOF {
		panic("Can't advance beyond the end of input")
	}
	p.curToken = p.peekToken
	p.peekToken = p.lexer.nextItem()
	fmt.Println("Reading next...", p.curToken, p.peekToken)
}

func (p *parser) parseRule() rule {
	rule := rule{}
	rule.pattern = p.parsePattern()
	rule.action = p.parseAction()
	return rule
}

func (p *parser) parsePattern() Node {
	return emptyNode{}
}

func (p *parser) parseAction() Node {
	if p.curToken.typ != '{' {
		panic("omg")
	}
	p.next()
	statements := p.parseStatements()

	if p.curToken.typ != '}' {
		panic("omg")
	}
	p.next()

	return statements
}

func (p *parser) parseStatements() Node {
	node := simpleNode{p.curToken}
	p.next()
	return &node
}

// Parse converts a string into an AST. The Node that is returned is the root of the tree.
func Parse(input string) Node {

	lexer := lex(input, os.Stderr)

	fmt.Println("Going to parse:", input)

	p := parser{lexer: lexer}
	p.parse()

	fmt.Println("AST:\n", p.root)

	return p.root
}
