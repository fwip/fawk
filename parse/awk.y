%{
package parse

import "fmt"
/*
 * CDDL HEADER START
 *
 * The contents of this file are subject to the terms of the
 * Common Development and Distribution License, Version 1.0 only
 * (the "License").  You may not use this file except in compliance
 * with the License.
 *
 * You can obtain a copy of the license at usr/src/OPENSOLARIS.LICENSE
 * or http://www.opensolaris.org/os/licensing.
 * See the License for the specific language governing permissions
 * and limitations under the License.
 *
 * When distributing Covered Code, include this CDDL HEADER in each
 * file and include the License file at usr/src/OPENSOLARIS.LICENSE.
 * If applicable, add the following below this CDDL HEADER, with the
 * fields enclosed by brackets "[]" replaced with your own identifying
 * information: Portions Copyright [yyyy] [name of copyright owner]
 *
 * CDDL HEADER END
 */
/*
 * awk -- YACC grammar
 *
 * Copyright (c) 1995 by Sun Microsystems, Inc.
 *
 * Copyright 1986, 1992 by Mortice Kern Systems Inc.  All rights reserved.
 *
 * This Software is unpublished, valuable, confidential property of
 * Mortice Kern Systems Inc.  Use is authorized only in accordance
 * with the terms and conditions of the source licence agreement
 * protecting this Software.  Any unauthorized use or disclosure of
 * this Software is strictly prohibited and will result in the
 * termination of the licence agreement.
 *
 * NOTE: this grammar correctly produces NO shift/reduce conflicts from YACC.
 *
 */

/*
 * Do not use any character constants as tokens, so the resulting C file
 * is codeset independent.
 */

var yytree *node
var NNULL *node

var doing_begin int
var npattern int
var funparm int

var redelim rune = '/'
var catterm rune

%}

%union	{
	node	*node;
};

/*
 * Do not use any character constants as tokens, so the resulting C file
 * is codeset independent.
 *
 * Declare terminal symbols before their operator
 * precedences to get them in a contiguous block
 * for giant switches in action() and exprreduce().
 */
/* Tokens from exprreduce() */
%token	<node>	PARM ARRAY UFUNC FIELD IN INDEX CONCAT
%token	<node>	NOT AND OR EXP QUEST
%token	<node>	EQ NE GE LE GT LT
%token	<node>	ADD SUB MUL DIV REM INC DEC PRE_INC PRE_DEC
%token	<node>	GETLINE CALLFUNC RE TILDE NRE

/* Tokens shared by exprreduce() and action() */
%token		ASG

/* Tokens from action() */
%token	<node>	PRINT PRINTF
%token	<node>	EXIT RETURN BREAK CONTINUE NEXT
%token	<node>	DELETE WHILE DO FOR FORIN IF

/*
 * Terminal symbols not used in action() and exprrreduce()
 * switch statements.
 */
%token	<node>	CONSTANT VAR FUNC NAME
%token	<node>	DEFFUNC BEGIN END CLOSE ELSE PACT
%right		ELSE
%token		DOT CALLUFUNC

/*
 * Tokens not used in grammar
 */
%token		KEYWORD SVAR
%token		PIPESYM

/*
 * Tokens representing character constants
 * TILDE, '~', taken care of above
 */
%token BAR		/* '|' */
       CARAT		/* '^' */
       LANGLE		/* '<' */
       RANGLE		/* '>' */
       PLUSC		/* '+' */
       HYPHEN		/* '-' */
       STAR		/* '*' */
       SLASH		/* '/' */
       PERCENT		/* '%' */
       EXCLAMATION	/* '!' */
       DOLLAR		/* '$' */
       LSQUARE		/* '[' */
       RSQUARE		/* ']' */
       LPAREN		/* '(' */
       RPAREN		/* ')' */
       SEMI		/* ';' */
       LBRACE		/* '{' */
       RBRACE		/* '}' */
			 NEWLINE

/*
 * Priorities of operators
 * Lowest to highest
 */
%left	COMMA
%right	BAR PIPE WRITE APPEND
%right	ASG AADD ASUB AMUL ADIV AREM AEXP
%right	QUEST COLON
%left	OR
%left	AND
%left	IN
%left	CARAT
%left	TILDE NRE
%left	EQ NE LANGLE RANGLE GE LE
%left	CONCAT
%left	PLUSC HYPHEN
%left	STAR SLASH PERCENT
%right	UPLUS UMINUS
%right	EXCLAMATION
%right	EXP
%right	INC DEC URE
%left	DOLLAR LSQUARE RSQUARE
%left	LPAREN RPAREN

%type	<node>	prog rule pattern expr rvalue lvalue fexpr varlist varlist2
%type	<node>	statement statlist fileout exprlist eexprlist simplepattern terminator
%type	<node>	getline optvar var
%type	<node>	dummy

%start	dummy
%%

dummy:
		prog			= {
			yytree = fliplist(yytree);
		}
		;
prog:
	  rule	= {
		yytree = $1;
	}
	| rule terminator prog		= {
		if ($1 != NNULL) {
			if (yytree != NNULL){
				yytree = newNode(COMMA, $1, yytree)
			} else {
				yytree = $1;
			}
		}
	}
	;

rule:	  pattern LBRACE statlist RBRACE = {
		$$ = newNode(PACT, $1, $3);
		doing_begin = 0;
	}
	| LBRACE statlist RBRACE		= {
		npattern++;
		$$ = newNode(PACT, NNULL, $2);
	}
	| pattern				= {
		$$ = newNode(PACT, $1, newNode(PRINT, NNULL, NNULL));
		doing_begin = 0;
	}
	| DEFFUNC VAR
		{ $2.typ = UFUNC; funparm = 1; }
	    LPAREN varlist RPAREN
		{ funparm = 0; }
	    LBRACE statlist { uexit($5); } RBRACE = {
		$2.ufunc = newNode(DEFFUNC, $5, fliplist($9));
		$$ = NNULL;
	}
	| DEFFUNC UFUNC				= {
		//awkerr((char *) gettext("function \"%S\" redefined"), $2.name);
		fmt.Println("yargh! Ye redefined a function matey", $2.name);
		/* NOTREACHED */
	}
	|					= {
		$$ = NNULL;
	}
	;

pattern:
	  simplepattern
	| expr COMMA expr			= {
		npattern++;
		$$ = newNode(COMMA, $1, $3);
	}
	;

simplepattern:
	  BEGIN					= {
		$$ = newNode(BEGIN, NNULL, NNULL);
		doing_begin++;
	}
	| END					= {
		npattern++;
		$$ = newNode(END, NNULL, NNULL);
	}
	| expr					 = {
		npattern++;
		$$ = $1;
	}
	;

eexprlist:
	  exprlist
	|					= {
		$$ = NNULL;
	}
	;

exprlist:
	  expr %prec COMMA
	| exprlist COMMA expr			= {
		$$ = newNode(COMMA, $1, $3);
	}
	;

varlist:
	  					= {
		$$ = NNULL;
	}
	| varlist2
	;

varlist2:
	  var
	| var COMMA varlist2			= {
		$$ = newNode(COMMA, $1, $3);
	}
	;

fexpr:
	  expr
	|					= {
		$$ = NNULL;
	}
	;

/*
 * Normal expression (includes regular expression)
 */
expr:
	  expr PLUSC expr			= {
		$$ = newNode(ADD, $1, $3);
	}
	| expr HYPHEN expr			= {
		$$ = newNode(SUB, $1, $3);
	}
	| expr STAR expr			= {
		$$ = newNode(MUL, $1, $3);
	}
	| expr SLASH expr			= {
		$$ = newNode(DIV, $1, $3);
	}
	| expr PERCENT expr			= {
		$$ = newNode(REM, $1, $3);
	}
	| expr EXP expr				= {
		$$ = newNode(EXP, $1, $3);
	}
	| expr AND expr				= {
		$$ = newNode(AND, $1, $3);
	}
	| expr OR expr				= {
		$$ = newNode(OR, $1, $3);
	}
	| expr QUEST expr COLON expr		= {
		$$ = newNode(QUEST, $1, newNode(COLON, $3, $5));
	}
	| lvalue ASG expr			= {
		$$ = newNode(ASG, $1, $3);
	}
	| lvalue AADD expr			= {
		$$ = newNode(AADD, $1, $3);
	}
	| lvalue ASUB expr			= {
		$$ = newNode(ASUB, $1, $3);
	}
	| lvalue AMUL expr			= {
		$$ = newNode(AMUL, $1, $3);
	}
	| lvalue ADIV expr			= {
		$$ = newNode(ADIV, $1, $3);
	}
	| lvalue AREM expr			= {
		$$ = newNode(AREM, $1, $3);
	}
	| lvalue AEXP expr			= {
		$$ = newNode(AEXP, $1, $3);
	}
	| lvalue INC				= {
		$$ = newNode(INC, $1, NNULL);
	}
	| lvalue DEC				= {
		$$ = newNode(DEC, $1, NNULL);
	}
	| expr EQ expr				= {
		$$ = newNode(EQ, $1, $3);
	}
	| expr NE expr				= {
		$$ = newNode(NE, $1, $3);
	}
	| expr RANGLE expr			= {
		$$ = newNode(GT, $1, $3);
	}
	| expr LANGLE expr			= {
		$$ = newNode(LT, $1, $3);
	}
	| expr GE expr				= {
		$$ = newNode(GE, $1, $3);
	}
	| expr LE expr				= {
		$$ = newNode(LE, $1, $3);
	}
	| expr TILDE expr			= {
		$$ = newNode(TILDE, $1, $3);
	}
	| expr NRE expr				= {
		$$ = newNode(NRE, $1, $3);
	}
	| expr IN var				= {
		$$ = newNode(IN, $3, $1);
	}
	| LPAREN exprlist RPAREN IN var		= {
		$$ = newNode(IN, $5, $2);
	}
	| getline
	| rvalue
	| expr CONCAT expr			= {
		$$ = newNode(CONCAT, $1, $3);
	}
	;

lvalue:
	  DOLLAR rvalue				= {
		$$ = newNode(FIELD, $2, NNULL);
	}
	/*
	 * Prevents conflict with FOR LPAREN var IN var RPAREN production
	 */
	| var %prec COMMA
	| var LSQUARE exprlist RSQUARE		= {
		$$ = newNode(INDEX, $1, $3);
	}
	;

var:
	  VAR
	| PARM
	/*| NAME*/
	;

rvalue:
	  lvalue %prec COMMA
	| CONSTANT
	| NAME
	| LPAREN expr RPAREN term		= {
		$$ = $2;
	}
	| EXCLAMATION expr			= {
		$$ = newNode(NOT, $2, NNULL);
	}
	| HYPHEN expr %prec UMINUS		= {
		$$ = newNode(SUB, const0, $2);
	}
	| PLUSC expr %prec UPLUS		= {
		$$ = $2;
	}
	| DEC lvalue				= {
		$$ = newNode(PRE_DEC, $2, NNULL);
	}
	| INC lvalue				= {
		$$ = newNode(PRE_INC, $2, NNULL);
	}
	| FUNC					= {
		$$ = newNode(CALLFUNC, $1, NNULL);
	}
	| FUNC LPAREN eexprlist RPAREN term	= {
		$$ = newNode(CALLFUNC, $1, $3);
	}
	| UFUNC LPAREN eexprlist RPAREN term	= {
		$$ = newNode(CALLUFUNC, $1, $3);
	}
	| VAR LPAREN eexprlist RPAREN term	= {
		$$ = newNode(CALLUFUNC, $1, $3);
	}
	| URE = {
		$$ = $<node>1;
	}
	;

statement:
	  FOR LPAREN fexpr SEMI fexpr SEMI fexpr RPAREN statement = {
		$$ = newNode(FOR, newNode(COMMA, $3, newNode(COMMA, $5, $7)), $9);
	}
	| FOR LPAREN var IN var RPAREN statement = {
		//register NODE *np;

		/*
		 * attempt to optimize statements for the form
		 *    for (i in x) delete x[i]
		 * to
		 *    delete x
		 */
		/*
		np = $7;
		if (np != NNULL
		 && np->n_type == DELETE
		 && (np = np->n_left)->n_type == INDEX
		 && np->n_left == $5
		 && np->n_right == $3){
			$$ = newNode(DELETE, $5, NNULL);
		} else {
			$$ = newNode(FORIN, newNode(IN, $3, $5), $7);
		}
		*/
	}
	| WHILE LPAREN expr RPAREN statement	= {
		$$ = newNode(WHILE, $3, $5);
	}
	| DO statement WHILE LPAREN expr RPAREN	= {
		$$ = newNode(DO, $5, $2);
	}
	| IF LPAREN expr RPAREN statement ELSE statement = {
		$$ = newNode(IF, $3, newNode(ELSE, $5, $7));
	}
	| IF LPAREN expr RPAREN statement %prec ELSE	= {
		$$ = newNode(IF, $3, newNode(ELSE, $5, NNULL));
	}
	| CONTINUE SEMI				= {
		$$ = newNode(CONTINUE, NNULL, NNULL);
	}
	| BREAK SEMI				= {
		$$ = newNode(BREAK, NNULL, NNULL);
	}
	| NEXT SEMI				= {
		$$ = newNode(NEXT, NNULL, NNULL);
	}
	| DELETE lvalue SEMI			= {
		$$ = newNode(DELETE, $2, NNULL);
	}
	| RETURN fexpr SEMI			= {
		$$ = newNode(RETURN, $2, NNULL);
	}
	| EXIT fexpr SEMI			= {
		$$ = newNode(EXIT, $2, NNULL);
	}
	| PRINT eexprlist fileout SEMI		= {
		$$ = newNode(PRINT, $2, $3);
	}
	| PRINT LPAREN exprlist RPAREN fileout SEMI	= {
		$$ = newNode(PRINT, $3, $5);
	}
	| PRINTF exprlist fileout SEMI		= {
		$$ = newNode(PRINTF, $2, $3);
	}
	| PRINTF LPAREN exprlist RPAREN fileout SEMI	= {
		$$ = newNode(PRINTF, $3, $5);
	}
	| expr SEMI 				= {
		$$ = $1;
	}
	| SEMI					= {
		$$ = NNULL;
	}
	| LBRACE statlist RBRACE		= {
		$$ = $2;
	}
	;


statlist:
	  statement
	| statlist statement			= {
		if ($1 == NNULL) {
			$$ = $2;
		} else if ($2 == NNULL) {
			$$ = $1;
		} else {
			$$ = newNode(COMMA, $1, $2);
		}
	}
	;

fileout:
	  WRITE expr				= {
		$$ = newNode(WRITE, $2, NNULL);
	}
	| APPEND expr				= {
		$$ = newNode(APPEND, $2, NNULL);
	}
	| PIPE expr				= {
		$$ = newNode(PIPE, $2, NNULL);
	}
	|					= {
		$$ = NNULL;
	}
	;

getline:
	  GETLINE optvar %prec WRITE		= {
		$$ = newNode(GETLINE, $2, NNULL);
	}
	| expr BAR GETLINE optvar		= {
		$$ = newNode(GETLINE, $4, newNode(PIPESYM, $1, NNULL));
	}
	| GETLINE optvar LANGLE expr		= {
		$$ = newNode(GETLINE, $2, newNode(LT, $4, NNULL));
	}
	;

optvar:
	  lvalue
	|					= {
		$$ = NNULL;
	}
	;

/*
optsemi:
		SEMI = {
		$$ = NNULL
	}
	| = {
		$$ = NNULL
	}
	;
*/

terminator:
		SEMI = {
		$$ = NNULL
	}
	| NEWLINE = {
		$$ = NNULL
	}
	;


term:
	  {catterm = 1;}
	;
%%
/*
 * Flip a left-recursively generated list
 * so that it can easily be traversed from left
 * to right without recursion.
 */


func fliplist(np *node) *node{
	return np

	if np == nil || np.isLeaf(){
		return np
	}

  // Flip right child
	np.right = fliplist(np.right)

	if np.typ == COMMA {
		for np.left != nil && np.left.typ == COMMA {
			np.left.right = fliplist(np.left.right)
			var spp *node
			for spp = np.left.right; spp != nil && spp.typ == COMMA; {
				fmt.Println("flipping")
			}
			np.left = spp
			spp = np
			np = spp.left
		}

	}

	if np.left != nil && np.typ != FUNC && np.typ != UFUNC {
		np.left = fliplist(np.left)
	}

	return np
}


/* Cleanup when exiting function */
/* noop */
func uexit(np *node){
}

/*
static NODE *
fliplist(np)
register NODE *np;
{
	int type;

	if (np!=NNULL && !isleaf(np->n_flags)
//#if 0
		 //&& (type = np->n_type)!=FUNC && type!=UFUNC
//#endif
	) {
		np->n_right = fliplist(np->n_right);
		if ((type=np->n_type)==COMMA) {
			register NODE *lp;

			while ((lp = np->n_left)!=NNULL && lp->n_type==COMMA) {
				register NODE* *spp;

				lp->n_right = fliplist(lp->n_right);
				for (spp = &lp->n_right; *spp != NNULL && (*spp)->n_type==COMMA; spp = &(*spp)->n_right) ;
				np->n_left = *spp;
				*spp = np;
				np = lp;
			}
		}
		if (np->n_left != NULL && (type = np->n_left->n_type)!= FUNC && type!=UFUNC)
			np->n_left = fliplist(np->n_left);
	}
	return (np);
}
*/
