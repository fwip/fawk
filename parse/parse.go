//line awk.y:2
package parse

import __yyfmt__ "fmt"

//line awk.y:2
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

//line awk.y:61
type yySymType struct {
	yys  int
	node *node
}

const PARM = 57346
const ARRAY = 57347
const UFUNC = 57348
const FIELD = 57349
const IN = 57350
const INDEX = 57351
const CONCAT = 57352
const NOT = 57353
const AND = 57354
const OR = 57355
const EXP = 57356
const QUEST = 57357
const EQ = 57358
const NE = 57359
const GE = 57360
const LE = 57361
const GT = 57362
const LT = 57363
const ADD = 57364
const SUB = 57365
const MUL = 57366
const DIV = 57367
const REM = 57368
const INC = 57369
const DEC = 57370
const PRE_INC = 57371
const PRE_DEC = 57372
const GETLINE = 57373
const CALLFUNC = 57374
const RE = 57375
const TILDE = 57376
const NRE = 57377
const ASG = 57378
const PRINT = 57379
const PRINTF = 57380
const EXIT = 57381
const RETURN = 57382
const BREAK = 57383
const CONTINUE = 57384
const NEXT = 57385
const DELETE = 57386
const WHILE = 57387
const DO = 57388
const FOR = 57389
const FORIN = 57390
const IF = 57391
const CONSTANT = 57392
const VAR = 57393
const FUNC = 57394
const DEFFUNC = 57395
const BEGIN = 57396
const END = 57397
const CLOSE = 57398
const ELSE = 57399
const PACT = 57400
const DOT = 57401
const CALLUFUNC = 57402
const KEYWORD = 57403
const SVAR = 57404
const PIPESYM = 57405
const BAR = 57406
const CARAT = 57407
const LANGLE = 57408
const RANGLE = 57409
const PLUSC = 57410
const HYPHEN = 57411
const STAR = 57412
const SLASH = 57413
const PERCENT = 57414
const EXCLAMATION = 57415
const DOLLAR = 57416
const LSQUARE = 57417
const RSQUARE = 57418
const LPAREN = 57419
const RPAREN = 57420
const SEMI = 57421
const LBRACE = 57422
const RBRACE = 57423
const COMMA = 57424
const PIPE = 57425
const WRITE = 57426
const APPEND = 57427
const AADD = 57428
const ASUB = 57429
const AMUL = 57430
const ADIV = 57431
const AREM = 57432
const AEXP = 57433
const COLON = 57434
const UPLUS = 57435
const UMINUS = 57436
const URE = 57437

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PARM",
	"ARRAY",
	"UFUNC",
	"FIELD",
	"IN",
	"INDEX",
	"CONCAT",
	"NOT",
	"AND",
	"OR",
	"EXP",
	"QUEST",
	"EQ",
	"NE",
	"GE",
	"LE",
	"GT",
	"LT",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"REM",
	"INC",
	"DEC",
	"PRE_INC",
	"PRE_DEC",
	"GETLINE",
	"CALLFUNC",
	"RE",
	"TILDE",
	"NRE",
	"ASG",
	"PRINT",
	"PRINTF",
	"EXIT",
	"RETURN",
	"BREAK",
	"CONTINUE",
	"NEXT",
	"DELETE",
	"WHILE",
	"DO",
	"FOR",
	"FORIN",
	"IF",
	"CONSTANT",
	"VAR",
	"FUNC",
	"DEFFUNC",
	"BEGIN",
	"END",
	"CLOSE",
	"ELSE",
	"PACT",
	"DOT",
	"CALLUFUNC",
	"KEYWORD",
	"SVAR",
	"PIPESYM",
	"BAR",
	"CARAT",
	"LANGLE",
	"RANGLE",
	"PLUSC",
	"HYPHEN",
	"STAR",
	"SLASH",
	"PERCENT",
	"EXCLAMATION",
	"DOLLAR",
	"LSQUARE",
	"RSQUARE",
	"LPAREN",
	"RPAREN",
	"SEMI",
	"LBRACE",
	"RBRACE",
	"COMMA",
	"PIPE",
	"WRITE",
	"APPEND",
	"AADD",
	"ASUB",
	"AMUL",
	"ADIV",
	"AREM",
	"AEXP",
	"COLON",
	"UPLUS",
	"UMINUS",
	"URE",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line awk.y:544

/*
 * Flip a left-recursively generated list
 * so that it can easily be traversed from left
 * to right without recursion.
 */

func fliplist(np *node) *node {
	if np == nil || np.isLeaf() {
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
func uexit(np *node) {
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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 109
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 983

var yyAct = [...]int{

	45, 8, 32, 202, 169, 86, 159, 110, 217, 150,
	16, 234, 182, 81, 150, 172, 170, 171, 229, 30,
	89, 90, 91, 181, 172, 170, 171, 200, 198, 113,
	8, 150, 150, 183, 101, 115, 230, 151, 104, 150,
	31, 150, 111, 111, 116, 116, 223, 222, 80, 218,
	112, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 199,
	141, 99, 143, 144, 145, 146, 147, 148, 149, 140,
	117, 78, 79, 194, 189, 153, 116, 168, 120, 167,
	71, 166, 108, 11, 219, 116, 116, 116, 172, 170,
	171, 107, 101, 111, 163, 106, 165, 29, 216, 83,
	161, 87, 187, 162, 186, 81, 92, 93, 185, 81,
	192, 154, 174, 101, 156, 157, 158, 177, 105, 103,
	102, 188, 96, 95, 109, 94, 85, 155, 221, 28,
	72, 73, 74, 75, 76, 77, 49, 164, 179, 190,
	173, 180, 142, 2, 175, 28, 184, 56, 14, 97,
	181, 233, 224, 121, 1, 13, 7, 201, 4, 3,
	0, 195, 196, 197, 82, 0, 0, 0, 0, 204,
	69, 0, 0, 98, 56, 0, 88, 0, 203, 56,
	111, 48, 205, 212, 211, 56, 213, 209, 206, 207,
	208, 210, 88, 214, 68, 215, 69, 0, 57, 58,
	56, 59, 60, 61, 64, 65, 85, 0, 0, 111,
	0, 225, 227, 0, 228, 15, 226, 0, 203, 0,
	66, 67, 0, 232, 101, 0, 87, 0, 51, 52,
	53, 54, 55, 51, 52, 53, 54, 55, 0, 0,
	0, 53, 54, 55, 0, 0, 28, 0, 25, 0,
	70, 0, 63, 62, 51, 52, 53, 54, 55, 0,
	231, 0, 0, 0, 0, 0, 0, 0, 0, 23,
	22, 0, 0, 17, 0, 0, 0, 0, 178, 43,
	44, 42, 41, 38, 37, 39, 40, 34, 35, 33,
	0, 36, 18, 26, 24, 0, 0, 0, 0, 0,
	0, 28, 0, 25, 0, 0, 0, 0, 0, 0,
	21, 20, 0, 27, 0, 19, 15, 0, 0, 12,
	0, 46, 47, 176, 23, 22, 0, 0, 17, 0,
	0, 0, 0, 0, 43, 44, 42, 41, 38, 37,
	39, 40, 34, 35, 33, 0, 36, 18, 26, 24,
	0, 0, 0, 0, 0, 0, 28, 0, 25, 0,
	0, 0, 0, 0, 0, 21, 20, 0, 27, 0,
	19, 15, 0, 0, 12, 0, 46, 47, 160, 23,
	22, 0, 0, 17, 0, 0, 0, 0, 0, 43,
	44, 42, 41, 38, 37, 39, 40, 34, 35, 33,
	0, 36, 18, 26, 24, 0, 0, 0, 0, 0,
	0, 28, 0, 25, 0, 0, 0, 0, 0, 0,
	21, 20, 0, 27, 0, 19, 15, 0, 0, 12,
	0, 46, 47, 100, 23, 22, 0, 0, 17, 0,
	0, 0, 0, 0, 43, 44, 42, 41, 38, 37,
	39, 40, 34, 35, 33, 0, 36, 18, 26, 24,
	0, 0, 68, 0, 69, 0, 57, 58, 56, 59,
	60, 61, 64, 65, 0, 21, 20, 0, 27, 0,
	19, 15, 0, 0, 12, 0, 46, 47, 66, 67,
	68, 0, 69, 0, 57, 58, 56, 59, 60, 61,
	64, 65, 0, 0, 68, 0, 69, 0, 57, 58,
	56, 59, 60, 61, 64, 65, 66, 67, 70, 0,
	63, 62, 51, 52, 53, 54, 55, 0, 0, 0,
	66, 67, 0, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 63, 62,
	51, 52, 53, 54, 55, 0, 0, 0, 0, 0,
	70, 119, 63, 62, 51, 52, 53, 54, 55, 0,
	0, 68, 0, 69, 220, 57, 58, 56, 59, 60,
	61, 64, 65, 0, 68, 0, 69, 0, 57, 58,
	56, 59, 60, 61, 64, 65, 0, 66, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 0, 69, 0, 57, 58, 56, 59,
	60, 61, 64, 65, 0, 0, 0, 70, 0, 63,
	62, 51, 52, 53, 54, 55, 0, 0, 66, 67,
	70, 193, 63, 62, 51, 52, 53, 54, 55, 0,
	0, 0, 0, 0, 191, 0, 28, 0, 25, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	63, 62, 51, 52, 53, 54, 55, 0, 0, 23,
	22, 0, 152, 17, 68, 0, 69, 0, 57, 58,
	56, 59, 60, 61, 64, 65, 0, 0, 0, 0,
	0, 0, 18, 26, 24, 6, 9, 10, 0, 0,
	66, 67, 28, 0, 25, 0, 0, 0, 0, 0,
	21, 20, 0, 27, 0, 19, 15, 0, 0, 12,
	0, 0, 5, 0, 0, 23, 22, 0, 0, 17,
	70, 0, 63, 62, 51, 52, 53, 54, 55, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 26,
	24, 68, 0, 69, 0, 57, 58, 56, 59, 60,
	61, 64, 65, 0, 0, 0, 21, 20, 0, 27,
	0, 19, 15, 0, 28, 12, 25, 66, 67, 0,
	0, 28, 0, 25, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 22, 0,
	0, 17, 0, 0, 23, 22, 0, 0, 17, 63,
	62, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	18, 26, 24, 0, 0, 0, 0, 18, 26, 24,
	0, 0, 28, 0, 25, 0, 0, 0, 21, 20,
	0, 27, 0, 19, 15, 21, 20, 118, 27, 0,
	19, 15, 0, 0, 114, 23, 22, 68, 0, 69,
	0, 57, 0, 56, 0, 60, 61, 64, 65, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 26,
	24, 69, 0, 66, 67, 56, 0, 60, 61, 64,
	65, 0, 0, 0, 0, 0, 21, 20, 68, 27,
	69, 19, 15, 0, 56, 84, 60, 61, 64, 65,
	0, 0, 0, 0, 0, 63, 62, 51, 52, 53,
	54, 55, 0, 0, 66, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 62, 51,
	52, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 62, 51, 52,
	53, 54, 55,
}
var yyPact = [...]int{

	662, -1000, -1000, 28, -61, 417, 140, -1000, 464, -1000,
	-1000, 54, 718, -1000, -1000, 848, 61, 151, -1000, 718,
	718, 718, 151, 151, 58, 56, 55, -1000, -1000, 662,
	417, 362, -1000, 53, 52, 417, 51, 26, 22, 13,
	151, 718, 718, 797, 790, 492, -1000, 417, -1000, -1000,
	718, 718, 718, 718, 718, 718, 718, 718, 718, 718,
	718, 718, 718, 718, 718, 718, 718, 718, 135, 718,
	121, 718, 718, 718, 718, 718, 718, 718, -1000, -1000,
	-41, 614, -1000, -1000, 718, 718, 71, -1000, -1000, 143,
	143, 143, -1000, -1000, 718, 718, 718, -89, -1000, 307,
	-1000, -1000, 718, 718, 102, 718, -1000, -1000, -1000, 12,
	10, 686, 8, -59, 718, -73, 686, -68, 718, -1000,
	252, 50, 686, 181, 181, 143, 143, 143, 143, 910,
	869, 196, 170, 170, 170, 170, 170, 170, 891, 891,
	-1000, 175, 151, 763, 763, 763, 763, 763, 763, 763,
	718, 152, -1000, 614, -43, 718, 40, 36, 34, 60,
	-1000, 5, 141, 586, 43, 573, -1000, -1000, -1000, 4,
	718, 718, 718, -50, -10, -51, -1000, 135, 718, -1000,
	686, 135, -1000, -1000, 170, -1000, -1000, -1000, -1000, 718,
	135, 417, 718, 417, -1000, 686, 686, 686, 15, -1000,
	15, 30, -1000, -74, 763, -1000, -1000, -1000, -1000, -30,
	16, -1000, 506, 81, -32, -33, -1000, 135, 718, 417,
	-1000, 417, -1000, -1000, -62, -1000, -42, -1000, -1000, 417,
	417, 417, -1000, -70, -1000,
}
var yyPgo = [...]int{

	0, 153, 169, 168, 0, 158, 93, 7, 167, 3,
	2, 40, 4, 35, 29, 166, 165, 5, 10, 164,
	163, 162, 161, 12, 159,
}
var yyR1 = [...]int{

	0, 19, 1, 1, 2, 2, 2, 20, 21, 22,
	2, 2, 2, 3, 3, 15, 15, 15, 14, 14,
	13, 13, 8, 8, 9, 9, 7, 7, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 6,
	6, 6, 18, 18, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 24, 5, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 11, 11, 12,
	12, 12, 12, 16, 16, 16, 17, 17, 23,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 4, 3, 1, 0, 0, 0,
	11, 2, 0, 1, 3, 1, 1, 1, 1, 0,
	1, 3, 0, 1, 1, 3, 1, 0, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 1, 1, 3, 2,
	1, 4, 1, 1, 1, 1, 4, 2, 2, 2,
	2, 2, 1, 5, 5, 5, 0, 4, 9, 7,
	5, 6, 7, 5, 2, 2, 2, 3, 3, 3,
	4, 6, 4, 6, 2, 1, 3, 1, 2, 2,
	2, 2, 0, 2, 4, 4, 1, 0, 0,
}
var yyChk = [...]int{

	-1000, -19, -1, -2, -3, 80, 53, -15, -4, 54,
	55, -6, 77, -16, -5, 74, -18, 31, 50, 73,
	69, 68, 28, 27, 52, 6, 51, 71, 4, 79,
	80, -11, -10, 47, 45, 46, 49, 42, 41, 43,
	44, 40, 39, 37, 38, -4, 79, 80, 51, 6,
	82, 68, 69, 70, 71, 72, 14, 12, 13, 15,
	16, 17, 67, 66, 18, 19, 34, 35, 8, 10,
	64, 36, 86, 87, 88, 89, 90, 91, 27, 28,
	-13, -4, -5, -6, 77, 75, -17, -6, 51, -4,
	-4, -4, -6, -6, 77, 77, 77, -24, -1, -11,
	81, -10, 77, 77, -10, 77, 79, 79, 79, -6,
	-7, -4, -7, -14, 77, -13, -4, -13, 77, 79,
	-11, -20, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-18, -4, 31, -4, -4, -4, -4, -4, -4, -4,
	82, 78, 78, -4, -13, 66, -14, -14, -14, 95,
	81, -7, -18, -4, 45, -4, 79, 79, 79, -12,
	84, 85, 83, -13, -12, -13, 81, 77, 92, -17,
	-4, 8, -23, 76, -4, 78, 78, 78, 71, 79,
	8, 78, 77, 78, 79, -4, -4, -4, 78, 79,
	78, -8, -9, -18, -4, -18, -23, -23, -23, -7,
	-18, -10, -4, -10, -12, -12, 78, 82, 79, 78,
	78, 57, 79, 79, -21, -9, -7, -10, -10, 80,
	78, -11, -10, -22, 81,
}
var yyDef = [...]int{

	12, -2, 1, 2, 6, 0, 0, 13, 17, 15,
	16, 64, 0, 56, 57, 0, 60, 107, 65, 0,
	0, 0, 0, 0, 72, 0, 62, 76, 63, 12,
	0, 0, 97, 0, 0, 0, 0, 0, 0, 0,
	0, 27, 27, 19, 0, 0, 95, 0, 7, 11,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 20, 59, 64, 0, 0, 103, 106, 62, 67,
	68, 69, 70, 71, 19, 19, 19, 0, 3, 0,
	5, 98, 27, 0, 0, 0, 84, 85, 86, 0,
	0, 26, 0, 102, 0, 18, 20, 102, 0, 94,
	0, 0, 14, 28, 29, 30, 31, 32, 33, 34,
	35, 0, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 58, 107, 37, 38, 39, 40, 41, 42, 43,
	0, 0, 108, 0, 0, 0, 0, 0, 0, 0,
	4, 0, 60, 0, 0, 0, 87, 88, 89, 0,
	0, 0, 0, 0, 0, 0, 96, 22, 0, 104,
	21, 0, 66, 61, 105, 108, 108, 108, 77, 27,
	0, 0, 0, 0, 90, 99, 100, 101, 102, 92,
	102, 0, 23, 24, 36, 55, 73, 74, 75, 0,
	0, 80, 0, 83, 0, 0, 8, 0, 27, 0,
	81, 0, 91, 93, 0, 25, 0, 79, 82, 0,
	0, 9, 78, 0, 10,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:159
		{
			yytree = fliplist(yytree)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:164
		{
			yytree = yyDollar[1].node
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:167
		{
			if yyDollar[1].node != NNULL {
				if yytree != NNULL {
					yytree = newNode(COMMA, yyDollar[1].node, yytree)
				} else {
					yytree = yyDollar[1].node
				}
			}
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:178
		{
			yyVAL.node = newNode(PACT, yyDollar[1].node, yyDollar[3].node)
			doing_begin = 0
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:182
		{
			npattern++
			yyVAL.node = newNode(PACT, NNULL, yyDollar[2].node)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:186
		{
			yyVAL.node = newNode(PACT, yyDollar[1].node, newNode(PRINT, NNULL, NNULL))
			doing_begin = 0
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:191
		{
			yyDollar[2].node.typ = UFUNC
			funparm = 1
		}
	case 8:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line awk.y:193
		{
			funparm = 0
		}
	case 9:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line awk.y:194
		{
			uexit(yyDollar[5].node)
		}
	case 10:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line awk.y:194
		{
			yyDollar[2].node.ufunc = newNode(DEFFUNC, yyDollar[5].node, fliplist(yyDollar[9].node))
			yyVAL.node = NNULL
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:198
		{
			//awkerr((char *) gettext("function \"%S\" redefined"), $2.name);
			fmt.Println("yargh! Ye redefined a function matey", yyDollar[2].node.name)
			/* NOTREACHED */
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:202
		{
			yyVAL.node = NNULL
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:209
		{
			npattern++
			yyVAL.node = newNode(COMMA, yyDollar[1].node, yyDollar[3].node)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:216
		{
			yyVAL.node = newNode(BEGIN, NNULL, NNULL)
			doing_begin++
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:220
		{
			npattern++
			yyVAL.node = newNode(END, NNULL, NNULL)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:224
		{
			npattern++
			yyVAL.node = yyDollar[1].node
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:232
		{
			yyVAL.node = NNULL
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:239
		{
			yyVAL.node = newNode(COMMA, yyDollar[1].node, yyDollar[3].node)
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:245
		{
			yyVAL.node = NNULL
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:253
		{
			yyVAL.node = newNode(COMMA, yyDollar[1].node, yyDollar[3].node)
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:260
		{
			yyVAL.node = NNULL
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:269
		{
			yyVAL.node = newNode(ADD, yyDollar[1].node, yyDollar[3].node)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:272
		{
			yyVAL.node = newNode(SUB, yyDollar[1].node, yyDollar[3].node)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:275
		{
			yyVAL.node = newNode(MUL, yyDollar[1].node, yyDollar[3].node)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:278
		{
			yyVAL.node = newNode(DIV, yyDollar[1].node, yyDollar[3].node)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:281
		{
			yyVAL.node = newNode(REM, yyDollar[1].node, yyDollar[3].node)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:284
		{
			yyVAL.node = newNode(EXP, yyDollar[1].node, yyDollar[3].node)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:287
		{
			yyVAL.node = newNode(AND, yyDollar[1].node, yyDollar[3].node)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:290
		{
			yyVAL.node = newNode(OR, yyDollar[1].node, yyDollar[3].node)
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:293
		{
			yyVAL.node = newNode(QUEST, yyDollar[1].node, newNode(COLON, yyDollar[3].node, yyDollar[5].node))
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:296
		{
			yyVAL.node = newNode(ASG, yyDollar[1].node, yyDollar[3].node)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:299
		{
			yyVAL.node = newNode(AADD, yyDollar[1].node, yyDollar[3].node)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:302
		{
			yyVAL.node = newNode(ASUB, yyDollar[1].node, yyDollar[3].node)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:305
		{
			yyVAL.node = newNode(AMUL, yyDollar[1].node, yyDollar[3].node)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:308
		{
			yyVAL.node = newNode(ADIV, yyDollar[1].node, yyDollar[3].node)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:311
		{
			yyVAL.node = newNode(AREM, yyDollar[1].node, yyDollar[3].node)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:314
		{
			yyVAL.node = newNode(AEXP, yyDollar[1].node, yyDollar[3].node)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:317
		{
			yyVAL.node = newNode(INC, yyDollar[1].node, NNULL)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:320
		{
			yyVAL.node = newNode(DEC, yyDollar[1].node, NNULL)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:323
		{
			yyVAL.node = newNode(EQ, yyDollar[1].node, yyDollar[3].node)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:326
		{
			yyVAL.node = newNode(NE, yyDollar[1].node, yyDollar[3].node)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:329
		{
			yyVAL.node = newNode(GT, yyDollar[1].node, yyDollar[3].node)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:332
		{
			yyVAL.node = newNode(LT, yyDollar[1].node, yyDollar[3].node)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:335
		{
			yyVAL.node = newNode(GE, yyDollar[1].node, yyDollar[3].node)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:338
		{
			yyVAL.node = newNode(LE, yyDollar[1].node, yyDollar[3].node)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:341
		{
			yyVAL.node = newNode(TILDE, yyDollar[1].node, yyDollar[3].node)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:344
		{
			yyVAL.node = newNode(NRE, yyDollar[1].node, yyDollar[3].node)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:347
		{
			yyVAL.node = newNode(IN, yyDollar[3].node, yyDollar[1].node)
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:350
		{
			yyVAL.node = newNode(IN, yyDollar[5].node, yyDollar[2].node)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:355
		{
			yyVAL.node = newNode(CONCAT, yyDollar[1].node, yyDollar[3].node)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:361
		{
			yyVAL.node = newNode(FIELD, yyDollar[2].node, NNULL)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:368
		{
			yyVAL.node = newNode(INDEX, yyDollar[1].node, yyDollar[3].node)
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:381
		{
			yyVAL.node = yyDollar[2].node
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:384
		{
			yyVAL.node = newNode(NOT, yyDollar[2].node, NNULL)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:387
		{
			yyVAL.node = newNode(SUB, const0, yyDollar[2].node)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:390
		{
			yyVAL.node = yyDollar[2].node
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:393
		{
			yyVAL.node = newNode(PRE_DEC, yyDollar[2].node, NNULL)
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:396
		{
			yyVAL.node = newNode(PRE_INC, yyDollar[2].node, NNULL)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:399
		{
			yyVAL.node = newNode(CALLFUNC, yyDollar[1].node, NNULL)
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:402
		{
			yyVAL.node = newNode(CALLFUNC, yyDollar[1].node, yyDollar[3].node)
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:405
		{
			yyVAL.node = newNode(CALLUFUNC, yyDollar[1].node, yyDollar[3].node)
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:408
		{
			yyVAL.node = newNode(CALLUFUNC, yyDollar[1].node, yyDollar[3].node)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:411
		{
			redelim = '/'
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:411
		{
			yyVAL.node = yyDollar[3].node
		}
	case 78:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line awk.y:417
		{
			yyVAL.node = newNode(FOR, newNode(COMMA, yyDollar[3].node, newNode(COMMA, yyDollar[5].node, yyDollar[7].node)), yyDollar[9].node)
		}
	case 79:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line awk.y:420
		{
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
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:440
		{
			yyVAL.node = newNode(WHILE, yyDollar[3].node, yyDollar[5].node)
		}
	case 81:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line awk.y:443
		{
			yyVAL.node = newNode(DO, yyDollar[5].node, yyDollar[2].node)
		}
	case 82:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line awk.y:446
		{
			yyVAL.node = newNode(IF, yyDollar[3].node, newNode(ELSE, yyDollar[5].node, yyDollar[7].node))
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line awk.y:449
		{
			yyVAL.node = newNode(IF, yyDollar[3].node, newNode(ELSE, yyDollar[5].node, NNULL))
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:452
		{
			yyVAL.node = newNode(CONTINUE, NNULL, NNULL)
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:455
		{
			yyVAL.node = newNode(BREAK, NNULL, NNULL)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:458
		{
			yyVAL.node = newNode(NEXT, NNULL, NNULL)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:461
		{
			yyVAL.node = newNode(DELETE, yyDollar[2].node, NNULL)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:464
		{
			yyVAL.node = newNode(RETURN, yyDollar[2].node, NNULL)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:467
		{
			yyVAL.node = newNode(EXIT, yyDollar[2].node, NNULL)
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:470
		{
			yyVAL.node = newNode(PRINT, yyDollar[2].node, yyDollar[3].node)
		}
	case 91:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line awk.y:473
		{
			yyVAL.node = newNode(PRINT, yyDollar[3].node, yyDollar[5].node)
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:476
		{
			yyVAL.node = newNode(PRINTF, yyDollar[2].node, yyDollar[3].node)
		}
	case 93:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line awk.y:479
		{
			yyVAL.node = newNode(PRINTF, yyDollar[3].node, yyDollar[5].node)
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:482
		{
			yyVAL.node = yyDollar[1].node
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line awk.y:485
		{
			yyVAL.node = NNULL
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line awk.y:488
		{
			yyVAL.node = yyDollar[2].node
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:496
		{
			if yyDollar[1].node == NNULL {
				yyVAL.node = yyDollar[2].node
			} else if yyDollar[2].node == NNULL {
				yyVAL.node = yyDollar[1].node
			} else {
				yyVAL.node = newNode(COMMA, yyDollar[1].node, yyDollar[2].node)
			}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:508
		{
			yyVAL.node = newNode(WRITE, yyDollar[2].node, NNULL)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:511
		{
			yyVAL.node = newNode(APPEND, yyDollar[2].node, NNULL)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:514
		{
			yyVAL.node = newNode(PIPE, yyDollar[2].node, NNULL)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:517
		{
			yyVAL.node = NNULL
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line awk.y:523
		{
			yyVAL.node = newNode(GETLINE, yyDollar[2].node, NNULL)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:526
		{
			yyVAL.node = newNode(GETLINE, yyDollar[4].node, newNode(PIPESYM, yyDollar[1].node, NNULL))
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line awk.y:529
		{
			yyVAL.node = newNode(GETLINE, yyDollar[2].node, newNode(LT, yyDollar[4].node, NNULL))
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:536
		{
			yyVAL.node = NNULL
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line awk.y:542
		{
			catterm = 1
		}
	}
	goto yystack /* stack new state and value */
}
