package parse

import (
	"fmt"
	"os"
)

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
		curToken := p.curToken
		p.next()
		return &simpleNode{curToken}
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

func (p *parser) atEOF() bool {
	return p.curToken.typ == itemEOF
}

func (p *parser) parsePrint() Node {
	ps := printStatement{
		printToken: p.curToken,
	}
	for !p.peekToken.isExpressionTerminator() {
		//fmt.Println("parser:", p)
		p.next()
		ps.arguments = append(ps.arguments, p.parseExpression(-1))
		if p.peekToken.typ == ',' {
			p.next()
		}
	}
	return &ps
}

func (p *parser) parseStatements() Node {
	var statements statementList

	for p.peekToken.typ != '}' {
		var n Node
		switch p.curToken.typ {
		case itemPrint:
			n = p.parsePrint()

		default:
			n = p.parseExpression(-1)
		}
		statements = append(statements, n)
		if p.atEOF() {
			panic(" Reached EOF!")
		}
	}
	//node := simpleNode{p.curToken}
	p.next()
	return statements
}

func (p *parser) peekPrecedence() int {
	return precedenceOf(p.peekToken.typ)
}
func (p *parser) curPrecedence() int {
	return precedenceOf(p.curToken.typ)
}

func precedenceOf(typ itemType) int {
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

func (p *parser) parseInfixExpression(left Node) Node {
	expression := infix{
		token:    p.curToken,
		operator: p.curToken.val,
		left:     left,
	}
	precedence := p.curPrecedence()
	p.next()
	expression.right = p.parseExpression(precedence)
	return &expression
}

func (p *parser) parseExpression(precedence int) Node {

	//prefix := p.prefixParseFns[p.curToken.typ]
	//if prefix == nil {
	//return emptyNode{}
	//}
	//switch p.curToken.typ {
	//case itemIdentifier:
	switch p.curToken.typ {
	case '$':
		token := p.curToken
		p.next()
		return &prefixExpression{
			prefix: token,
			right:  p.parseExpression(precedenceOf('$')),
		}
	}

	leftExp := (Node)(&simpleNode{p.curToken})

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
