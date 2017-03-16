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
	errors    []error

	//prefixParseFns map[itemType]prefixParseFn
	//infixParseFns  map[itemType]infixParseFn
}

//type prefixParseFn func() Node
//type infixParseFn func(Node) Node

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

func (p *parser) peekError(t itemType) {
	msg := fmt.Errorf("expected next token to be %s, got %s instead",
		t, p.peekToken.typ)
	p.errors = append(p.errors, msg)
}

func (p *parser) expectToken(typ itemType) bool {
	if p.peekToken.typ == typ {
		p.next()
		return true
	}
	p.peekError(typ)
	return false
}

func (p *parser) parseRule() rule {
	rule := rule{}
	rule.pattern = p.parsePattern()
	rule.action = p.parseAction()

	return rule
}

func (p *parser) parsePattern() Node {
	switch p.curToken.typ {

	case itemBegin, itemEnd:
		return &simpleNode{p.curToken}
	case '{':
		return emptyNode{}

	// TODO: Remove special-casing
	case itemRegex:
		curToken := p.curToken
		p.next()
		return &simpleNode{curToken}

	default:
		exp := p.parseExpression(-1)
		p.next()
		return exp
		// return emptyNode{}
	}

	//return emptyNode{}
}

func (p *parser) parseAction() Node {
	if p.curToken.typ == itemEOF {
		return emptyNode{}
	}
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

type expression struct {
	token          item
	nextExpression Node
}

func (e *expression) String() string {
	return e.token.String() + " " + e.nextExpression.String()
}

type statementList []Node

func (sl statementList) String() string {
	var lines []string

	for _, statement := range sl {
		lines = append(lines, statement.String())
	}
	return strings.Join(lines, "\n")
}

func (p *parser) parseStatements() Node {
	var statements statementList

	for p.peekToken.typ != '}' {
		statements = append(statements, p.parseExpression(-1))
		//p.next()
	}
	//node := simpleNode{p.curToken}
	p.next()
	return statements
}

func (p *parser) peekPrecedence() int {
	return precedence(p.peekToken.typ)
}
func (p *parser) curPrecedence() int {
	return precedence(p.curToken.typ)
}

func precedence(typ itemType) int {
	switch typ {
	case '=', itemAddAssign, itemSubAssign, itemMulAssign, itemDivAssign, itemModAssign, itemPowAssign:
		return 2
	// TODO is this right?
	case '?', ':':
		return 3
	case itemOr:
		return 4
	case itemAnd:
		return 5
	case '~', itemNoMatch:
		return 7
	case '<', itemLesserEqual, itemDoubleEqual, itemNotEqual, '>', itemGreaterEqual, '|':
		return 8
	case '+', '-':
		return 10
	case '%', '*', '/':
		return 20
	case '^':
		return 30
	case itemIncrement, itemDecrement:
		return 40
	case '$':
		return 50
	case '(':
		return 60
	}
	return 0

}

type infix struct {
	token    item
	left     Node
	operator string
	right    Node
}

func (ifx *infix) String() string {
	return fmt.Sprintf("( %s ) %s ( %s )", ifx.left.String(), ifx.operator, ifx.right.String())
	//return "(" + ifx.left.String() + ifx.operator + ifx.right.String()
}

func (p *parser) parseInfixExpression(left Node) Node {
	expression := infix{
		token:    p.curToken,
		operator: p.curToken.val,
		left:     left,
	}
	precedence := p.curPrecedence()
	p.next()
	expression.right = p.parseExpression(precedence)
	fmt.Println("Returning: ", expression.String())
	return &expression
}

func (p *parser) parseExpression(precedence int) Node {

	//prefix := p.prefixParseFns[p.curToken.typ]
	//if prefix == nil {
	//return emptyNode{}
	//}
	//switch p.curToken.typ {
	//case itemIdentifier:
	var leftExp Node
	leftExp = (Node)(&simpleNode{p.curToken})

	for !p.peekToken.isExpressionTerminator() && precedence < p.peekPrecedence() {
		p.next()
		leftExp = p.parseInfixExpression(leftExp)
	}

	//}
	return leftExp
}

// Make a new parser
func New(l *lexer) *parser {
	p := parser{lexer: l}
	//p.prefixParseFns = make(map[itemType]prefixParseFn)
	//p.infixParseFns = make(map[itemType]infixParseFn)

	//p[itemIdentifier] = p.parseIdentifier
	return &p
}

// Parse converts a string into an AST. The Node that is returned is the root of the tree.
func Parse(input string) Node {

	lexer := lex(input, os.Stderr)

	fmt.Println("Going to parse:", input)

	p := New(lexer)
	p.parse()

	fmt.Println("AST:\n", p.root)

	return p.root
}
