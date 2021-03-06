//line parse.y:2
package parse

import __yyfmt__ "fmt"

//line parse.y:2
import "os"

//line parse.y:8
type yySymType struct {
	yys int
	pos Pos    // The starting position, in bytes, of this item in the input string.
	val string // The value of this item.
}

const NAME = 57346
const NUMBER = 57347
const STRING = 57348
const ERE = 57349
const FUNC_NAME = 57350
const Begin = 57351
const End = 57352
const Break = 57353
const Continue = 57354
const Delete = 57355
const Do = 57356
const Else = 57357
const Exit = 57358
const For = 57359
const Function = 57360
const If = 57361
const In = 57362
const Next = 57363
const Print = 57364
const Printf = 57365
const Return = 57366
const While = 57367
const BUILTIN_FUNC_NAME = 57368
const GETLINE = 57369
const ADD_ASSIGN = 57370
const SUB_ASSIGN = 57371
const MUL_ASSIGN = 57372
const DIV_ASSIGN = 57373
const MOD_ASSIGN = 57374
const POW_ASSIGN = 57375
const OR = 57376
const AND = 57377
const NO_MATCH = 57378
const EQ = 57379
const LE = 57380
const GE = 57381
const NE = 57382
const INCR = 57383
const DECR = 57384
const APPEND = 57385
const NEWLINE = 57386

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NAME",
	"NUMBER",
	"STRING",
	"ERE",
	"FUNC_NAME",
	"Begin",
	"End",
	"Break",
	"Continue",
	"Delete",
	"Do",
	"Else",
	"Exit",
	"For",
	"Function",
	"If",
	"In",
	"Next",
	"Print",
	"Printf",
	"Return",
	"While",
	"BUILTIN_FUNC_NAME",
	"GETLINE",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
	"POW_ASSIGN",
	"OR",
	"AND",
	"NO_MATCH",
	"EQ",
	"LE",
	"GE",
	"NE",
	"INCR",
	"DECR",
	"APPEND",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'['",
	"']'",
	"','",
	"';'",
	"NEWLINE",
	"'!'",
	"'>'",
	"'<'",
	"'|'",
	"'?'",
	"':'",
	"'~'",
	"'$'",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'%'",
	"'/'",
	"'^'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse.y:392

func parse(s string) {
	l := lex(s, os.Stderr)
	yyErrorVerbose = true
	yyDebug = 5
	yyParse(l)
}

// Required by yacc
func (l *lexer) Lex(lval *yySymType) int {
	it := yySymType(l.nextItem())
	lval = &it
	return int(it.yys)
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 208
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1374

var yyAct = [...]int{

	134, 119, 189, 12, 12, 225, 126, 257, 398, 125,
	211, 192, 222, 190, 196, 393, 87, 88, 29, 89,
	91, 389, 31, 245, 242, 218, 107, 9, 35, 120,
	90, 106, 6, 9, 38, 39, 217, 42, 219, 210,
	209, 35, 111, 112, 38, 39, 423, 405, 140, 141,
	142, 143, 144, 145, 118, 146, 147, 148, 149, 150,
	151, 152, 153, 400, 42, 273, 157, 328, 159, 160,
	161, 162, 163, 164, 30, 165, 166, 167, 168, 169,
	170, 171, 172, 4, 390, 372, 176, 180, 180, 158,
	253, 336, 180, 45, 178, 182, 183, 184, 185, 186,
	187, 188, 178, 46, 191, 191, 191, 193, 195, 177,
	335, 181, 14, 180, 180, 406, 384, 179, 202, 194,
	178, 329, 251, 267, 105, 201, 53, 73, 434, 199,
	139, 212, 212, 424, 419, 413, 404, 403, 401, 395,
	239, 394, 213, 105, 327, 326, 254, 155, 156, 105,
	252, 237, 250, 249, 379, 333, 240, 241, 310, 311,
	308, 309, 307, 306, 332, 331, 73, 174, 175, 316,
	315, 304, 305, 207, 206, 243, 244, 205, 114, 98,
	99, 96, 97, 95, 94, 113, 73, 104, 103, 203,
	32, 312, 92, 93, 36, 330, 385, 248, 428, 410,
	417, 40, 41, 407, 204, 43, 255, 256, 208, 44,
	378, 359, 100, 346, 259, 324, 214, 322, 268, 269,
	270, 215, 272, 262, 266, 198, 173, 154, 109, 110,
	280, 294, 7, 271, 5, 224, 34, 28, 33, 272,
	17, 13, 220, 318, 216, 136, 319, 320, 321, 22,
	317, 135, 37, 117, 116, 197, 3, 2, 1, 0,
	0, 0, 246, 0, 247, 0, 0, 0, 191, 0,
	0, 0, 0, 0, 101, 102, 0, 0, 0, 0,
	0, 334, 108, 0, 0, 0, 294, 0, 272, 0,
	0, 0, 223, 260, 261, 0, 0, 0, 0, 350,
	294, 0, 0, 0, 0, 29, 20, 21, 23, 26,
	0, 0, 0, 0, 133, 0, 191, 191, 370, 371,
	0, 0, 0, 137, 138, 0, 0, 27, 32, 212,
	0, 0, 380, 381, 0, 0, 323, 0, 325, 377,
	259, 382, 24, 25, 0, 0, 0, 18, 0, 0,
	0, 0, 0, 373, 19, 374, 0, 337, 0, 0,
	0, 30, 0, 15, 16, 0, 0, 0, 347, 348,
	0, 0, 0, 0, 0, 0, 0, 397, 399, 0,
	402, 0, 360, 361, 396, 0, 0, 231, 231, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 212, 259, 411, 0,
	375, 376, 0, 0, 421, 0, 416, 0, 0, 0,
	426, 420, 0, 422, 266, 266, 0, 425, 0, 432,
	259, 429, 0, 0, 0, 0, 431, 0, 0, 420,
	266, 0, 425, 266, 431, 266, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 29, 20,
	21, 23, 26, 10, 11, 0, 0, 0, 0, 0,
	0, 0, 8, 0, 231, 231, 231, 231, 231, 0,
	27, 32, 0, 313, 314, 412, 0, 414, 415, 0,
	0, 0, 0, 0, 418, 24, 25, 0, 9, 0,
	18, 0, 0, 427, 0, 0, 0, 19, 430, 0,
	0, 0, 433, 0, 30, 0, 15, 16, 435, 301,
	302, 303, 0, 0, 231, 231, 231, 231, 231, 231,
	231, 231, 231, 0, 0, 0, 231, 0, 231, 231,
	231, 231, 231, 231, 231, 231, 231, 0, 0, 0,
	231, 0, 0, 0, 0, 0, 231, 231, 231, 231,
	231, 231, 231, 0, 0, 0, 0, 338, 339, 340,
	341, 342, 343, 0, 344, 345, 0, 0, 0, 349,
	0, 351, 352, 353, 354, 355, 356, 231, 357, 358,
	0, 0, 0, 362, 0, 0, 0, 231, 231, 363,
	364, 365, 366, 367, 368, 369, 0, 0, 0, 0,
	231, 231, 0, 383, 20, 21, 23, 26, 0, 0,
	0, 0, 133, 0, 0, 0, 0, 0, 0, 0,
	386, 137, 138, 0, 0, 27, 32, 0, 0, 231,
	387, 388, 0, 231, 0, 0, 0, 0, 0, 0,
	24, 25, 0, 391, 392, 18, 0, 29, 20, 21,
	23, 26, 19, 0, 0, 0, 0, 0, 0, 30,
	0, 15, 16, 82, 0, 0, 0, 0, 0, 27,
	32, 0, 408, 0, 0, 0, 409, 84, 83, 81,
	77, 75, 79, 76, 24, 25, 0, 0, 0, 18,
	0, 0, 0, 0, 0, 0, 19, 78, 74, 86,
	85, 0, 80, 30, 0, 71, 72, 68, 70, 69,
	67, 29, 20, 21, 23, 26, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 0, 0,
	0, 0, 0, 27, 32, 0, 0, 0, 0, 0,
	0, 64, 63, 61, 57, 55, 59, 56, 24, 25,
	0, 0, 0, 18, 0, 0, 0, 0, 0, 0,
	19, 58, 54, 66, 65, 0, 60, 30, 0, 51,
	52, 48, 50, 49, 47, 29, 20, 21, 23, 26,
	0, 0, 127, 128, 133, 132, 0, 130, 123, 0,
	121, 0, 129, 137, 138, 131, 122, 27, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 24, 25, 0, 9, 115, 18, 0, 0,
	0, 0, 124, 35, 19, 0, 0, 0, 0, 0,
	0, 30, 0, 15, 16, 29, 20, 21, 23, 26,
	0, 0, 127, 128, 133, 132, 0, 130, 265, 0,
	263, 0, 129, 137, 138, 131, 264, 27, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 24, 25, 0, 9, 0, 18, 0, 0,
	0, 0, 124, 35, 19, 0, 0, 0, 0, 0,
	0, 30, 0, 15, 16, 29, 20, 21, 23, 26,
	0, 0, 127, 128, 133, 132, 0, 130, 123, 0,
	121, 0, 129, 137, 138, 131, 122, 27, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 24, 25, 0, 9, 0, 18, 0, 0,
	0, 0, 124, 35, 19, 0, 0, 0, 0, 0,
	0, 30, 0, 15, 16, 29, 20, 21, 23, 26,
	0, 0, 127, 128, 133, 132, 0, 130, 123, 0,
	121, 0, 129, 137, 138, 131, 122, 27, 32, 0,
	0, 0, 0, 0, 0, 29, 229, 230, 232, 235,
	0, 0, 24, 25, 0, 9, 200, 18, 0, 0,
	0, 297, 124, 0, 19, 0, 0, 236, 0, 0,
	0, 30, 0, 15, 16, 299, 298, 296, 0, 0,
	0, 0, 233, 234, 0, 0, 0, 287, 0, 0,
	0, 0, 0, 0, 228, 0, 0, 0, 300, 0,
	295, 30, 0, 292, 293, 289, 291, 290, 288, 29,
	229, 230, 232, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 283, 0, 0, 0, 0,
	0, 236, 0, 0, 0, 0, 0, 0, 0, 285,
	284, 282, 0, 0, 0, 0, 233, 234, 0, 0,
	0, 287, 0, 0, 0, 0, 0, 0, 228, 0,
	0, 0, 286, 0, 281, 30, 0, 278, 279, 275,
	277, 276, 274, 258, 20, 21, 23, 26, 0, 0,
	0, 0, 133, 0, 0, 0, 0, 0, 0, 0,
	0, 137, 138, 0, 0, 27, 32, 0, 29, 20,
	21, 23, 26, 10, 11, 0, 0, 0, 0, 0,
	24, 25, 8, 0, 0, 18, 0, 0, 0, 0,
	27, 32, 19, 29, 20, 21, 23, 26, 0, 30,
	0, 15, 16, 0, 0, 24, 25, 0, 0, 0,
	18, 0, 0, 0, 0, 27, 32, 19, 0, 0,
	0, 0, 0, 0, 30, 0, 15, 16, 0, 0,
	24, 25, 0, 0, 0, 18, 29, 20, 21, 23,
	26, 35, 19, 0, 0, 0, 0, 0, 0, 30,
	0, 15, 16, 0, 0, 0, 0, 0, 27, 32,
	29, 229, 230, 232, 235, 0, 0, 0, 29, 229,
	230, 232, 235, 24, 25, 0, 0, 0, 18, 0,
	0, 0, 236, 0, 0, 19, 0, 0, 0, 0,
	236, 0, 30, 0, 15, 16, 0, 233, 234, 0,
	0, 0, 287, 0, 0, 233, 234, 0, 35, 228,
	287, 29, 229, 230, 232, 235, 30, 228, 226, 227,
	0, 0, 0, 0, 30, 0, 226, 227, 0, 0,
	0, 0, 0, 236, 29, 229, 230, 232, 235, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 233, 234,
	0, 0, 0, 238, 0, 0, 236, 0, 0, 0,
	228, 0, 0, 0, 0, 0, 0, 30, 0, 226,
	227, 233, 234, 0, 0, 0, 221, 0, 0, 0,
	0, 0, 0, 228, 0, 0, 0, 0, 0, 0,
	30, 0, 226, 227,
}
var yyPact = [...]int{

	-1000, -1000, 454, 1144, -24, -7, -7, -17, 201, -1000,
	-1000, -1000, 53, 717, 653, 1212, 1212, -1000, 1212, 1212,
	-1000, -1000, 151, -1000, 14, 14, 142, 141, -1000, 76,
	1212, -29, 14, -7, -17, -1000, -1000, -9, -1000, -1000,
	-1000, -1000, -1000, 139, 132, 781, -1000, 1212, 1212, 1212,
	1212, 1212, 1212, 653, 1212, 1212, 1212, 1212, 1212, 1212,
	1212, 1212, 223, -1000, -1000, 1212, 163, 1212, 1212, 1212,
	1212, 1212, 1212, 653, 1212, 1212, 1212, 1212, 1212, 1212,
	1212, 1212, 222, -1000, -1000, 1212, 163, -1000, -1000, 70,
	64, -1000, -1000, -1000, 1212, 1212, 1212, 1212, 1212, 1212,
	1212, -1000, -1000, 1212, 1212, 1212, -1000, 1212, -1000, -1000,
	-1000, -1000, -1000, 221, 221, -1000, 961, 144, -1000, -1000,
	-1000, 131, 128, 127, -1000, -12, -1000, -1000, -1000, -1000,
	1212, 1212, -1000, 217, -1000, -1000, -18, 1310, 1287, 1169,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 1169, 1169, -34, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 1169, 1169, -35, -1000, -1000, -1000,
	-1000, 177, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 106,
	-1000, 52, 42, 105, 73, -1000, 103, 40, -1000, 99,
	-1000, -1000, -1000, -1000, -24, 1212, 1212, 1119, -24, -1000,
	-1000, -1000, -1000, -1000, 841, 75, -1000, 1212, 1212, 1212,
	-1000, 1212, 15, -1000, 1055, 991, 1244, 1244, 1244, -1000,
	-1000, 130, -1000, 14, 14, 124, 123, 15, 1212, -1000,
	-1000, -1000, 1212, -1000, -1000, 1212, 1169, 1169, 213, -1000,
	-1000, -1000, -1000, 211, -1000, 98, 97, 16, 101, -1000,
	-24, -24, 170, 119, 118, 109, -12, 1212, -1000, -1000,
	-1000, 63, 44, -1000, 1244, 1244, 1244, 1244, 1244, 1244,
	991, 1244, 1244, 209, -1000, -1000, 1244, 1212, 1244, 1244,
	1244, 1244, 1244, 1244, 991, 1244, 1244, 207, -1000, -1000,
	1244, -1000, -1000, -1000, -1000, -1000, 1244, 1244, 1244, 1244,
	1244, 1244, 1244, -1000, -1000, 1212, 1212, 38, -1000, -1000,
	-1000, -1000, -1000, -11, -1000, -11, -1000, -1000, 1212, 206,
	108, 1212, 1212, 609, 67, 176, -1000, 1236, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 1236, 1236, -37,
	37, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1236, 1236, -43, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	94, 92, 176, -1000, -1000, 901, 901, 12, 91, 1212,
	90, 89, -4, 95, -1000, 199, -1000, -1000, -1000, 1244,
	176, -1000, -1000, 1244, -1000, -1000, 184, -1000, -1000, -1000,
	301, -1000, 88, -1000, -1000, 1212, 196, -1000, -1000, -1000,
	-1000, 87, 901, -1000, 841, 841, -5, 86, 901, -1000,
	-1000, -1000, 183, 301, -1000, -1000, -1000, 901, -1000, 81,
	841, -1000, -1000, 841, -1000, 841,
}
var yyPgo = [...]int{

	0, 258, 257, 256, 83, 234, 194, 29, 232, 14,
	255, 0, 254, 253, 252, 8, 1, 7, 10, 9,
	6, 13, 251, 245, 244, 242, 11, 12, 2, 241,
	112, 240, 249, 237, 292, 235, 5, 22,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 3, 3, 5,
	5, 5, 9, 9, 10, 10, 8, 8, 8, 8,
	7, 7, 7, 6, 6, 14, 14, 14, 14, 12,
	12, 13, 13, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 16, 16, 16, 16, 16, 16, 19, 19,
	19, 19, 19, 19, 19, 17, 17, 20, 20, 20,
	22, 22, 23, 23, 23, 23, 24, 24, 24, 28,
	28, 21, 21, 26, 26, 18, 18, 11, 11, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 25, 25, 27, 27, 34, 34, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 32, 32, 32,
	33, 33, 33, 31, 37, 37, 4, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 3, 3, 3, 2,
	7, 7, 0, 1, 1, 3, 1, 1, 1, 4,
	3, 4, 4, 0, 1, 2, 2, 1, 1, 1,
	2, 1, 2, 2, 6, 9, 6, 10, 8, 2,
	3, 3, 1, 6, 9, 6, 10, 8, 1, 1,
	1, 1, 2, 2, 7, 0, 1, 5, 1, 1,
	1, 2, 2, 4, 2, 4, 2, 2, 2, 0,
	1, 1, 1, 4, 4, 0, 1, 1, 1, 2,
	2, 3, 3, 3, 3, 3, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 4, 5,
	1, 3, 2, 3, 3, 3, 3, 3, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	4, 4, 5, 1, 1, 1, 1, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 4,
	1, 1, 0, 1, 1, 4, 1, 1, 2, 2,
	3, 3, 3, 3, 3, 3, 2, 3, 3, 3,
	4, 4, 5, 3, 2, 3, 3, 3, 3, 3,
	3, 2, 3, 3, 3, 5, 4, 4, 5, 1,
	1, 1, 1, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 1, 1, 4, 2,
	1, 3, 3, 3, 1, 2, 0, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -7, -8, 18, 44,
	9, 10, -11, -29, -30, 62, 63, -31, 46, 53,
	5, 6, -32, 7, 41, 42, 8, 26, -33, 4,
	60, -37, 27, -5, -8, 52, -6, -14, 51, 52,
	-6, -6, -7, 4, 8, -4, 50, 67, 64, 66,
	65, 62, 63, -30, 55, 38, 40, 37, 54, 39,
	59, 36, 20, 35, 34, 57, 56, 67, 64, 66,
	65, 62, 63, -30, 55, 38, 40, 37, 54, 39,
	59, 36, 20, 35, 34, 57, 56, -11, -11, -11,
	-26, -11, 41, 42, 33, 32, 30, 31, 28, 29,
	61, -32, -32, 46, 46, 48, -11, 55, -32, -6,
	-6, 51, 52, 46, 46, 45, -12, -13, -15, -16,
	-7, 19, 25, 17, 51, -19, -20, 11, 12, 21,
	16, 24, 14, 13, -11, -22, -23, 22, 23, -4,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, 4, -4, -4, -11, -37, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, 4, -4, -4, -11, -37, 50, 47,
	50, 47, -11, -11, -11, -11, -11, -11, -11, -28,
	-21, -11, -26, -28, -21, -11, -9, -10, 4, -9,
	45, -15, -16, 45, -4, 46, 46, 46, -4, 52,
	51, -18, -11, -18, -4, 4, -24, 54, 43, 56,
	-25, 46, -27, -34, -35, -36, 62, 63, 53, 5,
	6, -32, 7, 41, 42, 8, 26, -27, 46, -11,
	-11, -11, 58, -11, -11, 58, -4, -4, 20, 47,
	47, 49, 47, 50, 47, -11, -11, -17, 4, -20,
	-4, -4, -15, 19, 25, 17, -19, 48, -11, -11,
	-11, -26, -11, 50, 67, 64, 66, 65, 62, 63,
	-36, 59, 36, 20, 35, 34, 57, 46, 67, 64,
	66, 65, 62, 63, -36, 59, 36, 20, 35, 34,
	57, -34, -34, -34, 41, 42, 33, 32, 30, 31,
	28, 29, 61, -32, -32, 46, 46, -26, -11, -11,
	-11, -11, 4, -4, 4, -4, 47, 47, 51, 20,
	25, 46, 46, 46, -21, 47, 47, -4, -34, -34,
	-34, -34, -34, -34, -34, -34, 4, -4, -4, -34,
	-26, -34, -34, -34, -34, -34, -34, -34, -34, 4,
	-4, -4, -34, -34, -34, -34, -34, -34, -34, -34,
	-28, -28, 47, -7, -7, -4, -4, -18, 4, 46,
	-11, -11, -17, 4, 49, 20, -34, -34, -34, 58,
	47, -34, -34, 58, 47, 47, -15, -16, -15, -16,
	51, 47, -11, 47, 47, 51, 20, 4, -34, -34,
	15, -17, -4, 47, -4, -4, -18, 4, -4, 47,
	-15, -16, -15, 51, 47, -15, -16, -4, 15, -17,
	-4, -15, -16, -4, 47, -4,
}
var yyDef = [...]int{

	206, -2, 1, 2, 3, 23, 23, 23, 0, 206,
	16, 17, 18, 77, 78, 0, 0, 100, 0, 0,
	123, 124, 125, 126, 0, 0, 0, 140, 141, 197,
	0, 200, 204, 23, 23, 207, 5, 24, 27, 28,
	6, 7, 9, 0, 0, 0, 206, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 206, 206, 0, 0, 0, 0, 0,
	0, 0, 0, 109, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 206, 206, 0, 0, 79, 80, 0,
	0, 102, 127, 128, 0, 0, 0, 0, 0, 0,
	0, 129, 130, 69, 69, 0, 199, 0, 205, 4,
	8, 25, 26, 12, 12, 20, 0, 0, 29, 31,
	206, 0, 0, 0, 206, 42, 48, 49, 50, 51,
	75, 75, 206, 0, 58, 59, 60, 142, 0, 0,
	81, 82, 83, 84, 85, 86, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 0, 0, 0, 203, 103,
	104, 105, 106, 107, 108, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 0, 0, 0, 202, 206, 101,
	206, 0, 131, 132, 133, 134, 135, 136, 137, 0,
	70, 71, 72, 0, 0, 201, 0, 13, 14, 0,
	21, 30, 32, 22, 33, 0, 0, 55, 39, 206,
	206, 52, 76, 53, 0, 0, 61, 0, 0, 0,
	62, 0, 143, 144, 146, 147, 0, 0, 0, 179,
	180, 181, 182, 0, 0, 0, 196, 64, 0, 19,
	97, 98, 0, 120, 121, 0, 0, 0, 0, 138,
	139, 198, 206, 0, 206, 0, 0, 0, 197, 56,
	40, 41, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 0, 0, 206, 0, 0, 0, 0, 0, 0,
	156, 0, 0, 0, 206, 206, 0, 0, 0, 0,
	0, 0, 0, 0, 171, 0, 0, 0, 206, 206,
	0, 148, 149, 164, 183, 184, 0, 0, 0, 0,
	0, 0, 0, 185, 186, 69, 69, 0, 99, 122,
	73, 74, 119, 0, 15, 0, 206, 206, 75, 0,
	0, 0, 0, 55, 0, 63, 163, 0, 150, 151,
	152, 153, 154, 155, 157, 158, 159, 0, 0, 0,
	0, 165, 166, 167, 168, 169, 170, 172, 173, 174,
	0, 0, 0, 187, 188, 189, 190, 191, 192, 193,
	0, 0, 65, 10, 11, 0, 0, 0, 0, 0,
	0, 0, 0, 197, 57, 0, 145, 160, 161, 0,
	0, 176, 177, 0, 194, 195, 34, 43, 36, 45,
	55, 206, 0, 206, 206, 75, 0, 175, 162, 178,
	206, 0, 0, 54, 0, 0, 0, 0, 0, 206,
	38, 47, 34, 55, 206, 35, 44, 0, 206, 0,
	0, 37, 46, 0, 206, 0,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 3, 3, 60, 65, 3, 3,
	46, 47, 64, 62, 50, 63, 3, 66, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 58, 51,
	55, 61, 54, 57, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 48, 3, 49, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 56, 45, 59,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 52,
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

	}
	goto yystack /* stack new state and value */
}
