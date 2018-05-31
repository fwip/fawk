package parse

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var colors = []func(format string, a ...interface{}) string{
	color.RedString,
	color.MagentaString,
	color.BlueString,
	color.CyanString,
	color.GreenString,
}

var colorlevel = -1

func init() {
	color.NoColor = false
}

// Node represents any logical unit of the awk expression
// It must implement two methods, String() and ToGo()
type Node interface {
	String() string
	//RawValue() string
	//Position() Pos
	//Type() itemType
	//Children() []Node
	ToGo() string
}
type emptyNode struct{}

func (e emptyNode) String() string { return "" }
func (e emptyNode) ToGo() string   { return "" }

type simpleNode struct{ item item }

func (n *simpleNode) String() string {
	if n == nil {
		return ""
	}
	return n.item.val
}

func (n *simpleNode) ToGo() string { return "" }

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
func (p *program) ToGo() (out string) {
	return "package main"
}

type rule struct {
	pattern Node
	action  Node
}

func (r *rule) String() string {
	if r == nil {
		return "{ print }"
	}
	return r.pattern.String() + "{" + r.action.String() + "}"
}

type expression struct {
	token          item
	nextExpression Node
}

func (e *expression) String() string {
	return e.token.String() + " " + e.nextExpression.String()
}

type printStatement struct {
	printToken item
	arguments  []Node
}

func (ps *printStatement) String() string {
	out := "PRINT: "
	var args []string
	for _, n := range ps.arguments {
		args = append(args, n.String())
	}
	out += strings.Join(args, ", ")
	return out + " ENDPRINT"
}

func (ps *printStatement) ToGo() string {
	return fmt.Sprintf("fmt.println( %s )", ps.arguments)
}

type statementList []Node

func (sl statementList) String() string {
	var lines []string

	for _, statement := range sl {
		lines = append(lines, statement.String())
	}
	return strings.Join(lines, "\n")
}

func (sl statementList) ToGo() string {
	var lines []string
	for _, statement := range sl {
		lines = append(lines, statement.ToGo())
	}
	return strings.Join(lines, "\n")
}

type infix struct {
	token    item
	left     Node
	operator string
	right    Node
}

func (ifx *infix) String() string {
	colorlevel++
	colorf := colors[colorlevel%len(colors)]
	lb := colorf("❰")
	rb := colorf("❱")
	op := colorf(ifx.operator)
	str := lb + ifx.left.String() + rb + " " + op + " " + lb + ifx.right.String() + rb
	colorlevel--
	return str
}

func (ifx *infix) ToGo() string {
	switch ifx.operator {

	default:
		return ifx.left.ToGo() + " " + ifx.operator + " " + ifx.right.ToGo()
	}
}

type prefixExpression struct {
	prefix item
	right  Node
}

func (pe *prefixExpression) String() string {
	return pe.prefix.val + pe.right.String()
}

func (pe *prefixExpression) ToGo() string {
	switch pe.prefix {
	default:
		return pe.prefix.val + " " + pe.right.ToGo()
	}
}
