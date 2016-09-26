// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pos int

// item represents a token or text string returned from the scanner.
type item struct {
	yys int    // The type of this item.
	pos Pos    // The starting position, in bytes, of this item in the input string.
	val string // The value of this item.
}

func (i item) String() string {
	switch {
	case i.yys == int(itemEOF):
		return "EOF"
	case i.yys == int(itemError):
		return i.val
	case i.yys > int(itemKeyword):
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemEOF          itemType = iota
	itemError                 // error occurred; value is text of error
	itemBool                  // boolean constant
	itemChar                  // printable ASCII character; grab bag for comma etc.
	itemCharConstant          // character constant
	itemComplex               // complex constant (1+2i); imaginary is just a number
	itemColonEquals           // colon-equals (':=') introducing a declaration
	itemField                 // alphanumeric identifier starting with '$'
	itemIdentifier            // alphanumeric identifier not starting with '$'
	itemLeftDelim             // left action delimiter
	itemLeftParen             // '(' inside action
	itemNumber                // simple number, including imaginary
	itemPipe                  // pipe symbol
	itemRawString             // raw quoted string (includes quotes)
	itemRightDelim            // right action delimiter
	itemRightParen            // ')' inside action
	itemSpace                 // run of spaces separating arguments
	itemString                // quoted string (includes quotes)
	itemText                  // plain text
	itemVariable              // variable starting with '$', such as '$' or  '$1' or '$hello'
	itemRegex                 // regex surrounded by `/`
	itemOperator              // catchall for <=, +, >>, etc
	itemQuote                 // A quoted string
	itemComment               // A comment, starting with '#'
	itemBuiltinFunc           // A built-in function
	// Keywords appear after all the rest.
	itemKeyword          // used only to delimit the keywords
	itemBlock            // block keyword
	itemDot              // the cursor, spelled '.'
	itemDefine           // define keyword
	itemElse             // else keyword
	itemIf               // if keyword
	itemNil              // the untyped nil constant, easiest to treat as a keyword
	itemRange            // range keyword
	itemTemplate         // template keyword
	itemWith             // with keyword
	itemPrint    = Print // print function  TODO: Should maybe be not keyword?
	itemBegin    = Begin // BEGIN condition
	itemEnd      = End   // END condition
)

var key = map[string]itemType{
	".":        itemDot,
	"block":    itemBlock,
	"define":   itemDefine,
	"else":     itemElse,
	"if":       itemIf,
	"range":    itemRange,
	"nil":      itemNil,
	"template": itemTemplate,
	"with":     itemWith,
	"print":    itemPrint,
	"BEGIN":    itemBegin,
	"END":      itemEnd,
}

var builtinFuncs = map[string]int{
	"atan2":   0,
	"close":   0,
	"cos":     0,
	"exp":     0,
	"gsub":    0,
	"index":   0,
	"int":     0,
	"length":  0,
	"log":     0,
	"match":   0,
	"rand":    0,
	"sin":     0,
	"split":   0,
	"sprintf": 0,
	"sqrt":    0,
	"srand":   0,
	"sub":     0,
	"substr":  0,
	"system":  0,
	"tolower": 0,
	"toupper": 0,
}

const eof = -1

// Trimming spaces.
// If the action begins "{{- " rather than "{{", then all space/tab/newlines
// preceding the action are trimmed; conversely if it ends " -}}" the
// leading spaces are trimmed. This is done entirely in the lexer; the
// parser never sees it happen. We require an ASCII space to be
// present to avoid ambiguity with things like "{{-3}}". It reads
// better with the space present anyway. For simplicity, only ASCII
// space does the job.
const (
	spaceChars      = " \t\r\n" // These are the space characters defined by Go itself.
	leftTrimMarker  = "- "      // Attached to left delimiter, trims trailing spaces from preceding text.
	rightTrimMarker = " -"      // Attached to right delimiter, trims leading spaces from following text.
	trimMarkerLen   = Pos(len(leftTrimMarker))
)

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name       string    // the name of the input; used only for error reports
	input      string    // the string being scanned
	leftDelim  string    // start of action
	rightDelim string    // end of action
	state      stateFn   // the next lexing function to enter
	pos        Pos       // current position in the input
	start      Pos       // start position of this item
	width      Pos       // width of last rune read from input
	lastPos    Pos       // position of most recent item returned by nextItem
	items      chan item // channel of scanned items
	parenDepth int       // nesting depth of ( ) exprs
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	//fmt.Println("emit:", t, l.start, l.input[l.start:l.pos])
	l.items <- item{int(t), l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *lexer) acceptAlphaNumeric() {
	for isAlphaNumeric(l.next()) {
	}
	l.backup()
}

func (l *lexer) hasPrefix(s string) bool {
	return strings.HasPrefix(l.input[l.pos:], s)
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
func (l *lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.lastPos], "\n")
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{int(itemError), l.start, fmt.Sprintf(format, args...)}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	it, ok := <-l.items
	if !ok {
		return item{0, l.pos, "EOF"}
	}
	l.lastPos = it.pos
	return it
}

type yySymType item

// Required by yacc
func (l *lexer) Lex(lval *yySymType) int {
	it := yySymType(l.nextItem())
	lval = &it
	return int(it.yys)
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}

// lex creates a new scanner for the input string.
func lex(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexPattern; l.state != nil; {
		l.state = l.state(l)
	}
	l.emit(NEWLINE)
	l.emit(itemEOF)
	close(l.items)
}

func (l *lexer) tilNilThen(state, ret stateFn) stateFn {
	for state != nil {
		state = state(l)
	}
	return ret
}

// state functions

const (
	leftDelim    = "{{"
	rightDelim   = "}}"
	leftComment  = "/*"
	rightComment = "*/"
)

func lexPattern(l *lexer) stateFn {
	l.lookForRegex()
	return lexRule
}

// Lexes a 'rule' - what to do with lines that match a pattern
func lexRule(l *lexer) stateFn {
	r := l.next()

	if isSpace(r) {
		l.ignore()
		return lexRule
	}

	if isAlphaNumeric(r) {
		return lexRuleIdentifier
	}

	switch r {
	case '\\': //Skip literal backslash, consume next item
		l.next()

	case '"':
		l.consumeUntil(`"`)
		l.emit(itemQuote)

	case '{':
		l.emit('{')
	case '}':
		l.emit('}')
		return lexPattern

	// Simple single-char tokens
	case ';', '?', ':', ',':
		l.emit(itemType(r))

	case '#':
		l.consumeUntil("\n")
		l.backup()
		l.emit(itemComment)

	case '\n':
		l.emit(NEWLINE)

	case eof:
		return nil

	case '$':
		l.acceptAlphaNumeric()
		l.emit(itemField)

	case '%':
		l.emit2Char(map[rune]itemType{
			'=': MOD_ASSIGN,
		})
	case '+':
		l.emit2Char(map[rune]itemType{
			'+': INCR,
			'=': ADD_ASSIGN,
		})
	case '-':
		l.emit2Char(map[rune]itemType{
			'-': DECR,
			'=': SUB_ASSIGN,
		})

	case '*':
		l.emit2Char(map[rune]itemType{
			'=': MUL_ASSIGN,
		})

	case '/':
		l.emit2Char(map[rune]itemType{
			'=': DIV_ASSIGN,
		})

	case '^':
		l.emit2Char(map[rune]itemType{
			'=': POW_ASSIGN,
		})

	case '>':
		l.emit2Char(map[rune]itemType{
			'>': APPEND,
			'=': GE,
		})
	case '<':
		l.emit2Char(map[rune]itemType{
			'=': LE,
		})

	case '=':
		l.emit2Char(map[rune]itemType{
			'=': EQ,
		})

	case '&':
		if l.next() == '&' {
			l.emit(AND)
		} else {
			l.errorf("& is only allowed as &&")
		}
	case '|':
		if l.next() == '|' {
			l.emit(OR)
		} else {
			l.errorf("| is only allowed as ||")
		}

	case '!':
		switch l.next() {
		case '=':
			l.emit(EQ)
		case '~':
			l.emit(NO_MATCH)
			l.lookForRegex()
		default:
			l.backup()
		}

	case '~':
		l.emit('~')
		l.lookForRegex()

	default:
		l.emit(itemChar)
	}
	return lexRule
}

// Lexes a regex
// Currently will not work for / in []
func lexRegex(l *lexer) stateFn {
	l.consumeUntil("/")
	l.emit(itemRegex)
	return nil
}

func (l *lexer) lookForRegex() {
	for {
		r := l.next()
		switch r {
		case ' ', '\t':
			l.ignore()
		case '\n':
			l.emit(NEWLINE)
		case '/':
			l.consumeUntil("/")
			l.emit(ERE)
			return

		default:
			l.backup()
			return
		}
	}
}

func (l *lexer) emit2Char(m map[rune]itemType) {
	if val, ok := m[l.next()]; ok {
		l.emit(val)
		return
	}
	l.backup()

	r, _ := utf8.DecodeRuneInString(l.input[l.start:])
	l.emit(itemType(r))
}

func (l *lexer) consumeUntil(terms string) {
	for {
		r := l.next()
		if r == '\\' {
			l.next()
			continue
		}

		if strings.ContainsRune(terms, r) {
			return
		}
		if r == eof {
			l.errorf("Unclosed element: searching for '%s'", terms)
		}
	}
}

/*
// lexText scans until an opening action delimiter, "{{".
func lexText(l *lexer) stateFn {
	l.width = 0
	if x := strings.Index(l.input[l.pos:], l.leftDelim); x >= 0 {
		ldn := Pos(len(l.leftDelim))
		l.pos += Pos(x)
		trimLength := Pos(0)
		if strings.HasPrefix(l.input[l.pos+ldn:], leftTrimMarker) {
			trimLength = rightTrimLength(l.input[l.start:l.pos])
		}
		l.pos -= trimLength
		if l.pos > l.start {
			l.emit(itemText)
		}
		l.pos += trimLength
		l.ignore()
		return lexLeftDelim
	} else {
		l.pos = Pos(len(l.input))
	}
	// Correctly reached EOF.
	if l.pos > l.start {
		l.emit(itemText)
	}
	l.emit(itemEOF)
	return nil
}

// rightTrimLength returns the length of the spaces at the end of the string.
func rightTrimLength(s string) Pos {
	return Pos(len(s) - len(strings.TrimRight(s, spaceChars)))
}

// atRightDelim reports whether the lexer is at a right delimiter, possibly preceded by a trim marker.
func (l *lexer) atRightDelim() (delim, trimSpaces bool) {
	if strings.HasPrefix(l.input[l.pos:], l.rightDelim) {
		return true, false
	}
	// The right delim might have the marker before.
	if strings.HasPrefix(l.input[l.pos:], rightTrimMarker) {
		if strings.HasPrefix(l.input[l.pos+trimMarkerLen:], l.rightDelim) {
			return true, true
		}
	}
	return false, false
}

// leftTrimLength returns the length of the spaces at the beginning of the string.
func leftTrimLength(s string) Pos {
	return Pos(len(s) - len(strings.TrimLeft(s, spaceChars)))
}

// lexLeftDelim scans the left delimiter, which is known to be present, possibly with a trim marker.
func lexLeftDelim(l *lexer) stateFn {
	l.pos += Pos(len(l.leftDelim))
	trimSpace := strings.HasPrefix(l.input[l.pos:], leftTrimMarker)
	afterMarker := Pos(0)
	if trimSpace {
		afterMarker = trimMarkerLen
	}
	if strings.HasPrefix(l.input[l.pos+afterMarker:], leftComment) {
		l.pos += afterMarker
		l.ignore()
		return lexComment
	}
	l.emit(itemLeftDelim)
	l.pos += afterMarker
	l.ignore()
	l.parenDepth = 0
	return lexInsideAction
}

// lexComment scans a comment. The left comment marker is known to be present.
func lexComment(l *lexer) stateFn {
	l.pos += Pos(len(leftComment))
	i := strings.Index(l.input[l.pos:], rightComment)
	if i < 0 {
		return l.errorf("unclosed comment")
	}
	l.pos += Pos(i + len(rightComment))
	delim, trimSpace := l.atRightDelim()
	if !delim {
		return l.errorf("comment ends before closing delimiter")
	}
	if trimSpace {
		l.pos += trimMarkerLen
	}
	l.pos += Pos(len(l.rightDelim))
	if trimSpace {
		l.pos += leftTrimLength(l.input[l.pos:])
	}
	l.ignore()
	return lexText
}

// lexRightDelim scans the right delimiter, which is known to be present, possibly with a trim marker.
func lexRightDelim(l *lexer) stateFn {
	trimSpace := strings.HasPrefix(l.input[l.pos:], rightTrimMarker)
	if trimSpace {
		l.pos += trimMarkerLen
		l.ignore()
	}
	l.pos += Pos(len(l.rightDelim))
	l.emit(itemRightDelim)
	if trimSpace {
		l.pos += leftTrimLength(l.input[l.pos:])
		l.ignore()
	}
	return lexText
}

// lexInsideAction scans the elements inside action delimiters.
func lexInsideAction(l *lexer) stateFn {
	// Either number, quoted string, or identifier.
	// Spaces separate arguments; runs of spaces turn into itemSpace.
	// Pipe symbols separate and are emitted.
	delim, _ := l.atRightDelim()
	if delim {
		if l.parenDepth == 0 {
			return lexRightDelim
		}
		return l.errorf("unclosed left paren")
	}
	switch r := l.next(); {
	case r == eof || isEndOfLine(r):
		return l.errorf("unclosed action")
	case isSpace(r):
		return lexSpace
	case r == ':':
		if l.next() != '=' {
			return l.errorf("expected :=")
		}
		l.emit(itemColonEquals)
	case r == '|':
		l.emit(itemPipe)
	case r == '"':
		return lexQuote
	case r == '`':
		return lexRawQuote
	case r == '$':
		return lexVariable
	case r == '\'':
		return lexChar
	case r == '.':
		// special look-ahead for ".field" so we don't break l.backup().
		if l.pos < Pos(len(l.input)) {
			r := l.input[l.pos]
			if r < '0' || '9' < r {
				return lexField
			}
		}
		fallthrough // '.' can start a number.
	case r == '+' || r == '-' || ('0' <= r && r <= '9'):
		l.backup()
		return lexNumber
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	case r == '(':
		l.emit(itemLeftParen)
		l.parenDepth++
	case r == ')':
		l.emit(itemRightParen)
		l.parenDepth--
		if l.parenDepth < 0 {
			return l.errorf("unexpected right paren %#U", r)
		}
	case r <= unicode.MaxASCII && unicode.IsPrint(r):
		l.emit(itemChar)
		return lexInsideAction
	default:
		return l.errorf("unrecognized character in action: %#U", r)
	}
	return lexInsideAction
}

// lexSpace scans a run of space characters.
// One space has already been seen.
func lexSpace(l *lexer) stateFn {
	for isSpace(l.peek()) {
		l.next()
	}
	l.emit(itemSpace)
	return lexInsideAction
}

*/
// lexIdentifier scans an alphanumeric.
func lexRuleIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !l.atTerminator() {
				return l.errorf("bad character %#U", r)
			}
			// Following paren means it's a function name
			if l.peek() == '(' {
				if _, in := builtinFuncs[word]; in {
					l.emit(BUILTIN_FUNC_NAME)
					return lexRule
				}
				l.emit(FUNC_NAME)
				return lexRule
			}
			switch {
			case key[word] > itemKeyword:
				l.emit(key[word])
			case word[0] == '.':
				l.emit(itemField)
			case word == "true", word == "false":
				l.emit(itemBool)
			default:
				l.emit(NAME)
			}
			break Loop
		}
	}
	return lexRule
}

/*
// lexField scans a field: .Alphanumeric.
// The . has been scanned.
func lexField(l *lexer) stateFn {
	return lexFieldOrVariable(l, itemField)
}

// lexVariable scans a Variable: $Alphanumeric.
// The $ has been scanned.
func lexVariable(l *lexer) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "$".
		l.emit(itemVariable)
		return lexInsideAction
	}
	return lexFieldOrVariable(l, itemVariable)
}

// lexVariable scans a field or variable: [.$]Alphanumeric.
// The . or $ has been scanned.
func lexFieldOrVariable(l *lexer, typ itemType) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "." or "$".
		if typ == itemVariable {
			l.emit(itemVariable)
			//} else {
			//l.emit(itemDot)
		}
		return lexInsideAction
	}
	var r rune
	for {
		r = l.next()
		if !isAlphaNumeric(r) {
			l.backup()
			break
		}
	}
	if !l.atTerminator() {
		return l.errorf("bad character %#U", r)
	}
	l.emit(typ)
	//return lexInsideAction
	return nil
}
*/

// atTerminator reports whether the input is at valid termination character to
// appear after an identifier. Breaks .X.Y into two pieces. Also catches cases
// like "$x+2" not being acceptable without a space, in case we decide one
// day to implement arithmetic.
func (l *lexer) atTerminator() bool {
	r := l.peek()
	if isSpace(r) || isEndOfLine(r) {
		return true
	}
	switch r {
	case eof, '.', ',', '|', ':', ')', '(', '{', '}', ';', '%', '+', '-', '*', '/', '<', '>', '=':
		return true
	}
	// Does r start the delimiter? This can be ambiguous (with delim=="//", $x/2 will
	// succeed but should fail) but only in extremely rare cases caused by willfully
	// bad choice of delimiter.
	if rd, _ := utf8.DecodeRuneInString(l.rightDelim); rd == r {
		return true
	}
	return false
}

/*
// lexChar scans a character constant. The initial quote is already
// scanned. Syntax checking is done by the parser.
func lexChar(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated character constant")
		case '\'':
			break Loop
		}
	}
	l.emit(itemCharConstant)
	return lexInsideAction
}

// lexNumber scans a number: decimal, octal, hex, float, or imaginary. This
// isn't a perfect number scanner - for instance it accepts "." and "0x0.2"
// and "089" - but when it's wrong the input is invalid and the parser (via
// strconv) will notice.
func lexNumber(l *lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	if sign := l.peek(); sign == '+' || sign == '-' {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || l.input[l.pos-1] != 'i' {
			return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
		}
		l.emit(itemComplex)
	} else {
		l.emit(itemNumber)
	}
	return lexInsideAction
}

func (l *lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return false
	}
	return true
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(itemString)
	return lexInsideAction
}

// lexRawQuote scans a raw quoted string.
func lexRawQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof:
			return l.errorf("unterminated raw quoted string")
		case '`':
			break Loop
		}
	}
	l.emit(itemRawString)
	return lexInsideAction
}

*/
// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func (l *lexer) Error(s string) {
	fmt.Println(s)
}
