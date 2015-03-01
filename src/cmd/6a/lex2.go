//wline ./cmd/6a/a.y:32
package main

import __yyfmt__ "fmt"

//wline ./cmd/6a/a.y:32
import (
	"cmd/internal/asm"
	"cmd/internal/obj"
	"cmd/internal/obj/x86"
)

var (
	yyerror  = asm.Yyerror
	nullgen  obj.Addr
	stmtline int32
)

//wline ./cmd/6a/a.y:41
var lastpc *obj.Prog
type Addr2 struct {
	from obj.Addr
	to   obj.Addr
}

type yySymType struct {
	yys   int
	sym   *asm.Sym
	lval  int64
	dval  float64
	sval  string
	addr  obj.Addr
	addr2 Addr2
}

const LTYPE0 = 57346
const LTYPE1 = 57347
const LTYPE2 = 57348
const LTYPE3 = 57349
const LTYPE4 = 57350
const LTYPEC = 57351
const LTYPED = 57352
const LTYPEN = 57353
const LTYPER = 57354
const LTYPET = 57355
const LTYPEG = 57356
const LTYPEPC = 57357
const LTYPES = 57358
const LTYPEM = 57359
const LTYPEI = 57360
const LTYPEXC = 57361
const LTYPEX = 57362
const LTYPERT = 57363
const LTYPEF = 57364
const LCONST = 57365
const LFP = 57366
const LPC = 57367
const LSB = 57368
const LBREG = 57369
const LLREG = 57370
const LSREG = 57371
const LFREG = 57372
const LMREG = 57373
const LXREG = 57374
const LFCONST = 57375
const LSCONST = 57376
const LSP = 57377
const LNAME = 57378
const LLAB = 57379
const LVAR = 57380

var yyToknames = []string{
	"'|'",
	"'^'",
	"'&'",
	"'<'",
	"'>'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"LTYPE0",
	"LTYPE1",
	"LTYPE2",
	"LTYPE3",
	"LTYPE4",
	"LTYPEC",
	"LTYPED",
	"LTYPEN",
	"LTYPER",
	"LTYPET",
	"LTYPEG",
	"LTYPEPC",
	"LTYPES",
	"LTYPEM",
	"LTYPEI",
	"LTYPEXC",
	"LTYPEX",
	"LTYPERT",
	"LTYPEF",
	"LCONST",
	"LFP",
	"LPC",
	"LSB",
	"LBREG",
	"LLREG",
	"LSREG",
	"LFREG",
	"LMREG",
	"LXREG",
	"LFCONST",
	"LSCONST",
	"LSP",
	"LNAME",
	"LLAB",
	"LVAR",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//wline yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 2,
}

const yyNprod = 133
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 593

var yyAct = []int{

	52, 227, 41, 3, 80, 208, 269, 64, 123, 50,
	51, 79, 54, 170, 268, 74, 267, 118, 85, 72,
	83, 263, 73, 255, 253, 98, 241, 84, 81, 239,
	237, 100, 102, 112, 221, 219, 112, 210, 209, 171,
	240, 107, 234, 62, 211, 174, 143, 138, 65, 207,
	111, 119, 115, 113, 112, 231, 67, 169, 120, 121,
	122, 249, 230, 92, 94, 96, 128, 226, 225, 224,
	104, 106, 74, 58, 57, 154, 136, 112, 129, 85,
	153, 83, 151, 150, 139, 141, 149, 148, 84, 81,
	140, 147, 142, 146, 145, 144, 63, 55, 58, 57,
	137, 43, 45, 48, 44, 46, 49, 40, 135, 47,
	69, 134, 56, 127, 155, 40, 34, 37, 53, 31,
	59, 32, 55, 35, 33, 223, 176, 177, 222, 217,
	60, 215, 220, 112, 120, 243, 114, 56, 74, 242,
	216, 236, 183, 76, 173, 59, 58, 57, 256, 166,
	168, 251, 252, 192, 194, 196, 167, 112, 112, 112,
	112, 112, 195, 184, 112, 112, 112, 264, 58, 57,
	55, 212, 257, 248, 197, 198, 199, 200, 201, 182,
	120, 204, 205, 206, 218, 56, 42, 114, 152, 38,
	65, 76, 55, 59, 190, 191, 184, 261, 260, 166,
	168, 229, 258, 112, 112, 75, 167, 56, 89, 235,
	36, 71, 65, 76, 238, 59, 108, 109, 254, 213,
	232, 233, 125, 126, 228, 244, 247, 203, 245, 88,
	124, 181, 125, 126, 246, 158, 159, 160, 175, 250,
	202, 25, 185, 186, 187, 188, 189, 16, 15, 6,
	110, 259, 7, 2, 1, 262, 156, 157, 158, 159,
	160, 265, 266, 105, 9, 10, 11, 12, 13, 17,
	28, 18, 14, 29, 30, 26, 19, 20, 21, 22,
	23, 24, 27, 58, 57, 82, 165, 164, 163, 161,
	162, 156, 157, 158, 159, 160, 4, 103, 8, 101,
	5, 99, 97, 58, 57, 95, 93, 55, 91, 87,
	77, 43, 45, 48, 44, 46, 49, 68, 66, 47,
	86, 61, 56, 70, 214, 0, 78, 55, 53, 0,
	59, 43, 45, 48, 44, 46, 49, 172, 0, 47,
	60, 0, 56, 58, 57, 82, 0, 65, 53, 0,
	59, 43, 45, 48, 44, 46, 49, 0, 0, 47,
	0, 0, 0, 58, 57, 0, 0, 55, 0, 0,
	0, 43, 45, 48, 44, 46, 49, 0, 0, 47,
	86, 0, 56, 58, 57, 0, 0, 55, 53, 0,
	59, 43, 45, 48, 44, 46, 49, 0, 0, 47,
	60, 0, 56, 58, 57, 0, 90, 55, 53, 0,
	59, 43, 45, 48, 44, 46, 49, 58, 133, 47,
	60, 0, 56, 0, 0, 0, 39, 55, 53, 0,
	59, 43, 45, 48, 44, 46, 49, 58, 57, 47,
	60, 55, 56, 0, 58, 57, 0, 0, 53, 0,
	59, 131, 130, 0, 60, 0, 56, 58, 57, 0,
	0, 55, 132, 0, 59, 0, 116, 0, 55, 58,
	57, 0, 0, 117, 0, 0, 56, 0, 0, 0,
	0, 55, 76, 56, 59, 58, 179, 0, 193, 76,
	0, 59, 0, 55, 75, 0, 56, 58, 57, 0,
	0, 0, 76, 180, 59, 0, 0, 0, 56, 55,
	0, 58, 57, 0, 76, 0, 59, 0, 0, 178,
	0, 55, 0, 0, 56, 0, 0, 0, 0, 0,
	76, 0, 59, 0, 60, 55, 56, 0, 0, 0,
	0, 0, 53, 0, 59, 0, 0, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 76, 0, 59, 165,
	164, 163, 161, 162, 156, 157, 158, 159, 160, 164,
	163, 161, 162, 156, 157, 158, 159, 160, 163, 161,
	162, 156, 157, 158, 159, 160, 161, 162, 156, 157,
	158, 159, 160,
}
var yyPact = []int{

	-1000, -1000, 250, -1000, 70, -1000, 74, 66, 72, 65,
	374, 294, 294, 394, 159, -1000, -1000, 274, 354, 294,
	294, 294, 314, -5, -5, -1000, 294, 294, 84, 488,
	488, -1000, 502, -1000, -1000, 502, -1000, -1000, -1000, 394,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -2, 428, -3, -1000, -1000, 502, 502, 502,
	223, -1000, 61, -1000, -1000, 408, -1000, 59, -1000, 56,
	-1000, 448, -1000, 48, -7, 213, 502, -1000, 334, -1000,
	-1000, -1000, 64, -1000, -1000, -8, 223, -1000, -1000, -1000,
	394, -1000, 42, -1000, 41, -1000, 39, -1000, 35, -1000,
	34, -1000, -1000, -1000, 31, -1000, 30, 176, 28, 23,
	250, 555, -1000, 555, -1000, 111, 2, -16, 282, 106,
	-1000, -1000, -1000, -9, 230, 502, 502, -1000, -1000, -1000,
	-1000, -1000, 476, 460, 394, 294, -1000, 448, 128, -1000,
	-1000, -1000, -1000, 161, -9, 394, 394, 394, 394, 394,
	294, 294, 502, 435, 137, -1000, 502, 502, 502, 502,
	502, 233, 219, 502, 502, 502, -6, -17, -18, -10,
	502, -1000, -1000, 208, 95, 213, -1000, -1000, -20, 89,
	-1000, -1000, -1000, -1000, -21, 79, 76, -1000, 17, 16,
	-1000, -1000, 15, 191, 10, -1000, 3, 224, 224, -1000,
	-1000, -1000, 502, 502, 579, 572, 564, -12, 502, -1000,
	-1000, 103, -25, 502, -26, -1000, -1000, -1000, -14, -1000,
	-29, -1000, 101, 96, 502, 314, -5, -1000, 216, 140,
	8, -5, 247, 247, 113, -31, 207, -1000, -32, -1000,
	112, -1000, -1000, -1000, -1000, -1000, -1000, 139, 192, 191,
	-1000, 187, 186, -1000, 502, -1000, -34, -1000, 134, -1000,
	502, 502, -39, -1000, -1000, -41, -49, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 17, 324, 8, 186, 7, 1, 2, 12,
	4, 96, 43, 11, 9, 10, 210, 323, 189, 321,
	318, 317, 310, 309, 308, 306, 305, 302, 301, 299,
	297, 263, 254, 253, 3, 250, 249, 248, 247, 241,
}
var yyR1 = []int{

	0, 32, 33, 32, 35, 34, 34, 34, 34, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	16, 16, 20, 21, 19, 19, 18, 18, 17, 17,
	17, 37, 38, 38, 39, 39, 22, 22, 23, 23,
	24, 24, 25, 25, 26, 26, 26, 27, 28, 29,
	29, 30, 31, 11, 11, 13, 13, 13, 13, 13,
	13, 12, 12, 10, 10, 8, 8, 8, 8, 8,
	8, 8, 6, 6, 6, 6, 6, 6, 6, 5,
	5, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 15, 15, 9, 9, 4, 4, 4, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 7, 7,
	7, 7, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2,
}
var yyR2 = []int{

	0, 0, 0, 3, 0, 4, 1, 2, 2, 3,
	3, 2, 2, 2, 2, 2, 2, 1, 1, 2,
	2, 2, 2, 2, 2, 2, 2, 1, 2, 2,
	0, 1, 3, 3, 2, 1, 2, 1, 2, 1,
	3, 6, 5, 7, 4, 6, 2, 1, 1, 1,
	3, 5, 3, 5, 2, 1, 3, 5, 5, 0,
	1, 3, 3, 1, 1, 1, 1, 2, 2, 1,
	1, 1, 1, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 4, 5, 3, 1,
	1, 1, 4, 4, 4, 6, 9, 9, 3, 3,
	5, 8, 1, 6, 5, 7, 0, 2, 2, 1,
	1, 1, 1, 1, 2, 2, 2, 3, 1, 2,
	3, 4, 1, 3, 3, 3, 3, 3, 4, 4,
	3, 3, 3,
}
var yyChk = []int{

	-1000, -32, -33, -34, 46, 50, -36, 2, 48, 14,
	15, 16, 17, 18, 22, -37, -38, 19, 21, 26,
	27, 28, 29, 30, 31, -39, 25, 32, 20, 23,
	24, 49, 51, 50, 50, 51, -16, 52, -18, 52,
	-11, -8, -5, 37, 40, 38, 41, 45, 39, 42,
	-14, -15, -1, 54, -9, 33, 48, 10, 9, 56,
	46, -19, -12, -11, -6, 53, -20, -12, -21, -11,
	-17, 52, -10, -6, -1, 46, 54, -22, 52, -13,
	-10, -15, 11, -8, -14, -1, 46, -23, -16, -18,
	52, -24, -12, -25, -12, -26, -12, -27, -8, -28,
	-6, -29, -6, -30, -12, -31, -12, -9, -5, -5,
	-35, -2, -1, -2, -11, 54, 38, 45, -2, 54,
	-1, -1, -1, -4, 7, 9, 10, 52, -1, -9,
	44, 43, 54, 10, 52, 52, -10, 52, 54, -4,
	-13, -8, -14, 54, -4, 52, 52, 52, 52, 52,
	52, 52, 12, 52, 52, -34, 9, 10, 11, 12,
	13, 7, 8, 6, 5, 4, 38, 45, 39, 55,
	11, 55, 55, 38, 54, 8, -1, -1, 43, 10,
	43, -11, -12, -10, 35, -11, -11, -11, -11, -11,
	-12, -12, -1, 53, -1, -6, -1, -2, -2, -2,
	-2, -2, 7, 8, -2, -2, -2, 55, 11, 55,
	55, 54, -1, 11, -3, 36, 45, 34, -4, 55,
	43, 55, 49, 49, 52, 52, 52, -7, 33, 10,
	52, 52, -2, -2, 54, -1, 38, 55, -1, 55,
	54, 55, 38, 39, -1, -8, -6, 10, 33, 53,
	-6, 38, 39, 55, 11, 55, 36, 33, 10, -7,
	11, 11, -1, 55, 33, -1, -1, 55, 55, 55,
}
var yyDef = []int{

	1, -2, 0, 3, 0, 6, 0, 0, 0, 30,
	0, 0, 0, 0, 0, 17, 18, 0, 30, 0,
	0, 0, 0, 0, 59, 27, 0, 0, 0, 0,
	0, 4, 0, 7, 8, 0, 11, 31, 12, 0,
	37, 63, 64, 75, 76, 77, 78, 79, 80, 81,
	89, 90, 91, 0, 102, 112, 113, 0, 0, 0,
	106, 13, 35, 71, 72, 0, 14, 0, 15, 0,
	16, 0, 39, 0, 0, 106, 0, 19, 0, 47,
	65, 66, 0, 69, 70, 91, 106, 20, 48, 49,
	31, 21, 0, 22, 0, 23, 55, 24, 0, 25,
	0, 26, 60, 28, 0, 29, 0, 0, 0, 0,
	0, 9, 122, 10, 36, 0, 0, 0, 0, 0,
	114, 115, 116, 0, 0, 0, 0, 34, 82, 83,
	84, 85, 0, 0, 0, 0, 38, 0, 0, 74,
	46, 67, 68, 0, 74, 0, 0, 54, 0, 0,
	0, 0, 0, 0, 0, 5, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 98,
	0, 99, 117, 0, 0, 106, 107, 108, 0, 0,
	88, 32, 33, 40, 0, 50, 52, 56, 0, 0,
	61, 62, 0, 0, 0, 44, 0, 123, 124, 125,
	126, 127, 0, 0, 130, 131, 132, 92, 0, 93,
	94, 0, 0, 0, 0, 109, 110, 111, 0, 86,
	0, 73, 0, 0, 0, 0, 0, 42, 118, 0,
	0, 0, 128, 129, 0, 0, 0, 100, 0, 104,
	0, 87, 51, 53, 57, 58, 41, 0, 119, 0,
	45, 0, 0, 95, 0, 103, 0, 120, 0, 43,
	0, 0, 0, 105, 121, 0, 0, 101, 96, 97,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 53, 13, 6, 3,
	54, 55, 11, 9, 52, 10, 3, 12, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 50,
	7, 51, 8, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 56,
}
var yyTok2 = []int{

	2, 3, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48,
}
var yyTok3 = []int{
	0,
}

//wline yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yylex.Error("syntax error")
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
			yychar = -1
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

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:72
		{
			stmtline = asm.Lineno
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:79
		{
			yyDollar[1].sym = asm.LabelLookup(yyDollar[1].sym)
			if yyDollar[1].sym.Type == LLAB && yyDollar[1].sym.Value != int64(asm.PC) {
				yyerror("redeclaration of %s (%s)", yyDollar[1].sym.Labelname, yyDollar[1].sym.Name)
			}
			yyDollar[1].sym.Type = LLAB
			yyDollar[1].sym.Value = int64(asm.PC)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:94
		{
			yyDollar[1].sym.Type = LVAR
			yyDollar[1].sym.Value = yyDollar[3].lval
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:99
		{
			if yyDollar[1].sym.Value != yyDollar[3].lval {
				yyerror("redeclaration of %s", yyDollar[1].sym.Name)
			}
			yyDollar[1].sym.Value = yyDollar[3].lval
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:105
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:106
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:107
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:108
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:109
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:110
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:113
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:114
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:115
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:116
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:117
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:118
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:119
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:120
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:122
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:123
		{
			outcode(int(yyDollar[1].lval), &yyDollar[2].addr2)
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//wline ./cmd/6a/a.y:126
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = nullgen
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:131
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = nullgen
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:138
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:145
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:152
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = nullgen
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:157
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = nullgen
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:164
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[2].addr
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:169
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[1].addr
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:176
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[2].addr
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:181
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[1].addr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:186
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//wline ./cmd/6a/a.y:193
		{
			var a Addr2
			a.from = yyDollar[2].addr
			a.to = yyDollar[6].addr
			outcode(obj.ADATA, &a)
			if asm.Pass > 1 {
				lastpc.From3.Type = obj.TYPE_CONST
				lastpc.From3.Offset = yyDollar[4].lval
			}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:206
		{
			asm.Settext(yyDollar[2].addr.Sym)
			outcode(obj.ATEXT, &Addr2{yyDollar[2].addr, yyDollar[5].addr})
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//wline ./cmd/6a/a.y:211
		{
			asm.Settext(yyDollar[2].addr.Sym)
			outcode(obj.ATEXT, &Addr2{yyDollar[2].addr, yyDollar[7].addr})
			if asm.Pass > 1 {
				lastpc.From3.Type = obj.TYPE_CONST
				lastpc.From3.Offset = yyDollar[4].lval
			}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:222
		{
			asm.Settext(yyDollar[2].addr.Sym)
			outcode(obj.AGLOBL, &Addr2{yyDollar[2].addr, yyDollar[4].addr})
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
		//wline ./cmd/6a/a.y:227
		{
			asm.Settext(yyDollar[2].addr.Sym)
			outcode(obj.AGLOBL, &Addr2{yyDollar[2].addr, yyDollar[6].addr})
			if asm.Pass > 1 {
				lastpc.From3.Type = obj.TYPE_CONST
				lastpc.From3.Offset = yyDollar[4].lval
			}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:238
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[2].addr
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:243
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = yyDollar[1].addr
		}
	case 48:
		yyVAL.addr2 = yyS[yypt-0].addr2
	case 49:
		yyVAL.addr2 = yyS[yypt-0].addr2
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:254
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:259
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
			if yyVAL.addr2.from.Index != obj.TYPE_NONE {
				yyerror("dp shift with lhs index")
			}
			yyVAL.addr2.from.Index = int16(yyDollar[5].lval)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:270
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:275
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
			if yyVAL.addr2.to.Index != obj.TYPE_NONE {
				yyerror("dp move with lhs index")
			}
			yyVAL.addr2.to.Index = int16(yyDollar[5].lval)
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:286
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = nullgen
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:291
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = nullgen
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:296
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:303
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
			yyVAL.addr2.to.Offset = yyDollar[5].lval
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:311
		{
			yyVAL.addr2.from = yyDollar[3].addr
			yyVAL.addr2.to = yyDollar[5].addr
			if yyDollar[1].addr.Type != obj.TYPE_CONST {
				yyerror("illegal constant")
			}
			yyVAL.addr2.to.Offset = yyDollar[1].addr.Offset
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//wline ./cmd/6a/a.y:321
		{
			yyVAL.addr2.from = nullgen
			yyVAL.addr2.to = nullgen
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:326
		{
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = nullgen
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:333
		{
			if yyDollar[1].addr.Type != obj.TYPE_CONST || yyDollar[3].addr.Type != obj.TYPE_CONST {
				yyerror("arguments to asm.PCDATA must be integer constants")
			}
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:343
		{
			if yyDollar[1].addr.Type != obj.TYPE_CONST {
				yyerror("index for FUNCDATA must be integer constant")
			}
			if yyDollar[3].addr.Type != obj.TYPE_MEM || (yyDollar[3].addr.Name != obj.NAME_EXTERN && yyDollar[3].addr.Name != obj.NAME_STATIC) {
				yyerror("value for FUNCDATA must be symbol reference")
			}
			yyVAL.addr2.from = yyDollar[1].addr
			yyVAL.addr2.to = yyDollar[3].addr
		}
	case 63:
		yyVAL.addr = yyS[yypt-0].addr
	case 64:
		yyVAL.addr = yyS[yypt-0].addr
	case 65:
		yyVAL.addr = yyS[yypt-0].addr
	case 66:
		yyVAL.addr = yyS[yypt-0].addr
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:362
		{
			yyVAL.addr = yyDollar[2].addr
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:366
		{
			yyVAL.addr = yyDollar[2].addr
		}
	case 69:
		yyVAL.addr = yyS[yypt-0].addr
	case 70:
		yyVAL.addr = yyS[yypt-0].addr
	case 71:
		yyVAL.addr = yyS[yypt-0].addr
	case 72:
		yyVAL.addr = yyS[yypt-0].addr
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:378
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_BRANCH
			yyVAL.addr.Offset = yyDollar[1].lval + int64(asm.PC)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:384
		{
			yyDollar[1].sym = asm.LabelLookup(yyDollar[1].sym)
			yyVAL.addr = nullgen
			if asm.Pass == 2 && yyDollar[1].sym.Type != LLAB {
				yyerror("undefined label: %s", yyDollar[1].sym.Labelname)
			}
			yyVAL.addr.Type = obj.TYPE_BRANCH
			yyVAL.addr.Offset = yyDollar[1].sym.Value + yyDollar[2].lval
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:396
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:402
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:408
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:414
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:420
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = x86.REG_SP
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:426
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:432
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_REG
			yyVAL.addr.Reg = int16(yyDollar[1].lval)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:440
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_CONST
			yyVAL.addr.Offset = yyDollar[2].lval
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:446
		{
			yyVAL.addr = yyDollar[2].addr
			yyVAL.addr.Type = obj.TYPE_ADDR
			/*
				if($2.Type == x86.D_AUTO || $2.Type == x86.D_PARAM)
					yyerror("constant cannot be automatic: %s",
						$2.sym.Name);
			*/
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:455
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_SCONST
			yyVAL.addr.U.Sval = (yyDollar[2].sval + "\x00\x00\x00\x00\x00\x00\x00\x00")[:8]
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:461
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_FCONST
			yyVAL.addr.U.Dval = yyDollar[2].dval
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:467
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_FCONST
			yyVAL.addr.U.Dval = yyDollar[3].dval
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:473
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_FCONST
			yyVAL.addr.U.Dval = -yyDollar[4].dval
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:479
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_FCONST
			yyVAL.addr.U.Dval = -yyDollar[3].dval
		}
	case 89:
		yyVAL.addr = yyS[yypt-0].addr
	case 90:
		yyVAL.addr = yyS[yypt-0].addr
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:491
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Offset = yyDollar[1].lval
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:497
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[3].lval)
			yyVAL.addr.Offset = yyDollar[1].lval
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:504
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = x86.REG_SP
			yyVAL.addr.Offset = yyDollar[1].lval
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:511
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[3].lval)
			yyVAL.addr.Offset = yyDollar[1].lval
		}
	case 95:
		yyDollar = yyS[yypt-6 : yypt+1]
		//wline ./cmd/6a/a.y:518
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Offset = yyDollar[1].lval
			yyVAL.addr.Index = int16(yyDollar[3].lval)
			yyVAL.addr.Scale = int16(yyDollar[5].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 96:
		yyDollar = yyS[yypt-9 : yypt+1]
		//wline ./cmd/6a/a.y:527
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[3].lval)
			yyVAL.addr.Offset = yyDollar[1].lval
			yyVAL.addr.Index = int16(yyDollar[6].lval)
			yyVAL.addr.Scale = int16(yyDollar[8].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 97:
		yyDollar = yyS[yypt-9 : yypt+1]
		//wline ./cmd/6a/a.y:537
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[3].lval)
			yyVAL.addr.Offset = yyDollar[1].lval
			yyVAL.addr.Index = int16(yyDollar[6].lval)
			yyVAL.addr.Scale = int16(yyDollar[8].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:547
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[2].lval)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:553
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = x86.REG_SP
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:559
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Index = int16(yyDollar[2].lval)
			yyVAL.addr.Scale = int16(yyDollar[4].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 101:
		yyDollar = yyS[yypt-8 : yypt+1]
		//wline ./cmd/6a/a.y:567
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Reg = int16(yyDollar[2].lval)
			yyVAL.addr.Index = int16(yyDollar[5].lval)
			yyVAL.addr.Scale = int16(yyDollar[7].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:578
		{
			yyVAL.addr = yyDollar[1].addr
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//wline ./cmd/6a/a.y:582
		{
			yyVAL.addr = yyDollar[1].addr
			yyVAL.addr.Index = int16(yyDollar[3].lval)
			yyVAL.addr.Scale = int16(yyDollar[5].lval)
			checkscale(yyVAL.addr.Scale)
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//wline ./cmd/6a/a.y:591
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Name = int8(yyDollar[4].lval)
			yyVAL.addr.Sym = obj.Linklookup(asm.Ctxt, yyDollar[1].sym.Name, 0)
			yyVAL.addr.Offset = yyDollar[2].lval
		}
	case 105:
		yyDollar = yyS[yypt-7 : yypt+1]
		//wline ./cmd/6a/a.y:599
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_MEM
			yyVAL.addr.Name = obj.NAME_STATIC
			yyVAL.addr.Sym = obj.Linklookup(asm.Ctxt, yyDollar[1].sym.Name, 1)
			yyVAL.addr.Offset = yyDollar[4].lval
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//wline ./cmd/6a/a.y:608
		{
			yyVAL.lval = 0
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:612
		{
			yyVAL.lval = yyDollar[2].lval
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:616
		{
			yyVAL.lval = -yyDollar[2].lval
		}
	case 109:
		yyVAL.lval = yyS[yypt-0].lval
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:623
		{
			yyVAL.lval = obj.NAME_AUTO
		}
	case 111:
		yyVAL.lval = yyS[yypt-0].lval
	case 112:
		yyVAL.lval = yyS[yypt-0].lval
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:631
		{
			yyVAL.lval = yyDollar[1].sym.Value
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:635
		{
			yyVAL.lval = -yyDollar[2].lval
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:639
		{
			yyVAL.lval = yyDollar[2].lval
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:643
		{
			yyVAL.lval = ^yyDollar[2].lval
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:647
		{
			yyVAL.lval = yyDollar[2].lval
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//wline ./cmd/6a/a.y:653
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_TEXTSIZE
			yyVAL.addr.Offset = yyDollar[1].lval
			yyVAL.addr.U.Argsize = obj.ArgsSizeUnknown
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//wline ./cmd/6a/a.y:660
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_TEXTSIZE
			yyVAL.addr.Offset = -yyDollar[2].lval
			yyVAL.addr.U.Argsize = obj.ArgsSizeUnknown
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:667
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_TEXTSIZE
			yyVAL.addr.Offset = yyDollar[1].lval
			yyVAL.addr.U.Argsize = int32(yyDollar[3].lval)
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:674
		{
			yyVAL.addr = nullgen
			yyVAL.addr.Type = obj.TYPE_TEXTSIZE
			yyVAL.addr.Offset = -yyDollar[2].lval
			yyVAL.addr.U.Argsize = int32(yyDollar[4].lval)
		}
	case 122:
		yyVAL.lval = yyS[yypt-0].lval
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:684
		{
			yyVAL.lval = yyDollar[1].lval + yyDollar[3].lval
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:688
		{
			yyVAL.lval = yyDollar[1].lval - yyDollar[3].lval
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:692
		{
			yyVAL.lval = yyDollar[1].lval * yyDollar[3].lval
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:696
		{
			yyVAL.lval = yyDollar[1].lval / yyDollar[3].lval
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:700
		{
			yyVAL.lval = yyDollar[1].lval % yyDollar[3].lval
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:704
		{
			yyVAL.lval = yyDollar[1].lval << uint(yyDollar[4].lval)
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		//wline ./cmd/6a/a.y:708
		{
			yyVAL.lval = yyDollar[1].lval >> uint(yyDollar[4].lval)
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:712
		{
			yyVAL.lval = yyDollar[1].lval & yyDollar[3].lval
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:716
		{
			yyVAL.lval = yyDollar[1].lval ^ yyDollar[3].lval
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//wline ./cmd/6a/a.y:720
		{
			yyVAL.lval = yyDollar[1].lval | yyDollar[3].lval
		}
	}
	goto yystack /* stack new state and value */
}

func checkscale(scale int16) {
	switch scale {
	case 1,
		2,
		4,
		8:
		return
	}

	yyerror("scale must be 1248: %d", scale)
}

func cclean() {
	var g2 Addr2

	g2.from = nullgen
	g2.to = nullgen
	outcode(obj.AEND, &g2)
}

func outcode(a int, g2 *Addr2) {
	var p *obj.Prog
	var pl *obj.Plist

	if asm.Pass == 1 {
		goto out
	}

	p = new(obj.Prog)
	*p = obj.Prog{}
	p.Ctxt = asm.Ctxt
	p.As = int16(a)
	p.Lineno = stmtline
	p.From = g2.from
	p.To = g2.to
	p.Pc = int64(asm.PC)

	if lastpc == nil {
		pl = obj.Linknewplist(asm.Ctxt)
		pl.Firstpc = p
	} else {

		lastpc.Link = p
	}
	lastpc = p

out:
	if a != obj.AGLOBL && a != obj.ADATA {
		asm.PC++
	}
}

var lexinit = []asm.Lextab{
	{"SP", LSP, obj.NAME_AUTO},
	{"SB", LSB, obj.NAME_EXTERN},
	{"FP", LFP, obj.NAME_PARAM},
	{"PC", LPC, obj.TYPE_BRANCH},

	{"AL", LBREG, x86.REG_AL},
	{"CL", LBREG, x86.REG_CL},
	{"DL", LBREG, x86.REG_DL},
	{"BL", LBREG, x86.REG_BL},
	/*	"SPB",		LBREG,	REG_SPB,	*/
	{"SIB", LBREG, x86.REG_SIB},
	{"DIB", LBREG, x86.REG_DIB},
	{"BPB", LBREG, x86.REG_BPB},
	{"R8B", LBREG, x86.REG_R8B},
	{"R9B", LBREG, x86.REG_R9B},
	{"R10B", LBREG, x86.REG_R10B},
	{"R11B", LBREG, x86.REG_R11B},
	{"R12B", LBREG, x86.REG_R12B},
	{"R13B", LBREG, x86.REG_R13B},
	{"R14B", LBREG, x86.REG_R14B},
	{"R15B", LBREG, x86.REG_R15B},
	{"AH", LBREG, x86.REG_AH},
	{"CH", LBREG, x86.REG_CH},
	{"DH", LBREG, x86.REG_DH},
	{"BH", LBREG, x86.REG_BH},
	{"AX", LLREG, x86.REG_AX},
	{"CX", LLREG, x86.REG_CX},
	{"DX", LLREG, x86.REG_DX},
	{"BX", LLREG, x86.REG_BX},

	/*	"SP",		LLREG,	REG_SP,	*/
	{"BP", LLREG, x86.REG_BP},
	{"SI", LLREG, x86.REG_SI},
	{"DI", LLREG, x86.REG_DI},
	{"R8", LLREG, x86.REG_R8},
	{"R9", LLREG, x86.REG_R9},
	{"R10", LLREG, x86.REG_R10},
	{"R11", LLREG, x86.REG_R11},
	{"R12", LLREG, x86.REG_R12},
	{"R13", LLREG, x86.REG_R13},
	{"R14", LLREG, x86.REG_R14},
	{"R15", LLREG, x86.REG_R15},
	{"RARG", LLREG, x86.REGARG},
	{"F0", LFREG, x86.REG_F0 + 0},
	{"F1", LFREG, x86.REG_F0 + 1},
	{"F2", LFREG, x86.REG_F0 + 2},
	{"F3", LFREG, x86.REG_F0 + 3},
	{"F4", LFREG, x86.REG_F0 + 4},
	{"F5", LFREG, x86.REG_F0 + 5},
	{"F6", LFREG, x86.REG_F0 + 6},
	{"F7", LFREG, x86.REG_F0 + 7},
	{"M0", LMREG, x86.REG_M0 + 0},
	{"M1", LMREG, x86.REG_M0 + 1},
	{"M2", LMREG, x86.REG_M0 + 2},
	{"M3", LMREG, x86.REG_M0 + 3},
	{"M4", LMREG, x86.REG_M0 + 4},
	{"M5", LMREG, x86.REG_M0 + 5},
	{"M6", LMREG, x86.REG_M0 + 6},
	{"M7", LMREG, x86.REG_M0 + 7},
	{"X0", LXREG, x86.REG_X0 + 0},
	{"X1", LXREG, x86.REG_X0 + 1},
	{"X2", LXREG, x86.REG_X0 + 2},
	{"X3", LXREG, x86.REG_X0 + 3},
	{"X4", LXREG, x86.REG_X0 + 4},
	{"X5", LXREG, x86.REG_X0 + 5},
	{"X6", LXREG, x86.REG_X0 + 6},
	{"X7", LXREG, x86.REG_X0 + 7},
	{"X8", LXREG, x86.REG_X0 + 8},
	{"X9", LXREG, x86.REG_X0 + 9},
	{"X10", LXREG, x86.REG_X0 + 10},
	{"X11", LXREG, x86.REG_X0 + 11},
	{"X12", LXREG, x86.REG_X0 + 12},
	{"X13", LXREG, x86.REG_X0 + 13},
	{"X14", LXREG, x86.REG_X0 + 14},
	{"X15", LXREG, x86.REG_X0 + 15},
	{"CS", LSREG, x86.REG_CS},
	{"SS", LSREG, x86.REG_SS},
	{"DS", LSREG, x86.REG_DS},
	{"ES", LSREG, x86.REG_ES},
	{"FS", LSREG, x86.REG_FS},
	{"GS", LSREG, x86.REG_GS},
	{"GDTR", LBREG, x86.REG_GDTR},
	{"IDTR", LBREG, x86.REG_IDTR},
	{"LDTR", LBREG, x86.REG_LDTR},
	{"MSW", LBREG, x86.REG_MSW},
	{"TASK", LBREG, x86.REG_TASK},
	{"CR0", LBREG, x86.REG_CR + 0},
	{"CR1", LBREG, x86.REG_CR + 1},
	{"CR2", LBREG, x86.REG_CR + 2},
	{"CR3", LBREG, x86.REG_CR + 3},
	{"CR4", LBREG, x86.REG_CR + 4},
	{"CR5", LBREG, x86.REG_CR + 5},
	{"CR6", LBREG, x86.REG_CR + 6},
	{"CR7", LBREG, x86.REG_CR + 7},
	{"CR8", LBREG, x86.REG_CR + 8},
	{"CR9", LBREG, x86.REG_CR + 9},
	{"CR10", LBREG, x86.REG_CR + 10},
	{"CR11", LBREG, x86.REG_CR + 11},
	{"CR12", LBREG, x86.REG_CR + 12},
	{"CR13", LBREG, x86.REG_CR + 13},
	{"CR14", LBREG, x86.REG_CR + 14},
	{"CR15", LBREG, x86.REG_CR + 15},
	{"DR0", LBREG, x86.REG_DR + 0},
	{"DR1", LBREG, x86.REG_DR + 1},
	{"DR2", LBREG, x86.REG_DR + 2},
	{"DR3", LBREG, x86.REG_DR + 3},
	{"DR4", LBREG, x86.REG_DR + 4},
	{"DR5", LBREG, x86.REG_DR + 5},
	{"DR6", LBREG, x86.REG_DR + 6},
	{"DR7", LBREG, x86.REG_DR + 7},
	{"TR0", LBREG, x86.REG_TR + 0},
	{"TR1", LBREG, x86.REG_TR + 1},
	{"TR2", LBREG, x86.REG_TR + 2},
	{"TR3", LBREG, x86.REG_TR + 3},
	{"TR4", LBREG, x86.REG_TR + 4},
	{"TR5", LBREG, x86.REG_TR + 5},
	{"TR6", LBREG, x86.REG_TR + 6},
	{"TR7", LBREG, x86.REG_TR + 7},
	{"TLS", LSREG, x86.REG_TLS},
	{"AAA", LTYPE0, x86.AAAA},
	{"AAD", LTYPE0, x86.AAAD},
	{"AAM", LTYPE0, x86.AAAM},
	{"AAS", LTYPE0, x86.AAAS},
	{"ADCB", LTYPE3, x86.AADCB},
	{"ADCL", LTYPE3, x86.AADCL},
	{"ADCQ", LTYPE3, x86.AADCQ},
	{"ADCW", LTYPE3, x86.AADCW},
	{"ADDB", LTYPE3, x86.AADDB},
	{"ADDL", LTYPE3, x86.AADDL},
	{"ADDQ", LTYPE3, x86.AADDQ},
	{"ADDW", LTYPE3, x86.AADDW},
	{"ADJSP", LTYPE2, x86.AADJSP},
	{"ANDB", LTYPE3, x86.AANDB},
	{"ANDL", LTYPE3, x86.AANDL},
	{"ANDQ", LTYPE3, x86.AANDQ},
	{"ANDW", LTYPE3, x86.AANDW},
	{"ARPL", LTYPE3, x86.AARPL},
	{"BOUNDL", LTYPE3, x86.ABOUNDL},
	{"BOUNDW", LTYPE3, x86.ABOUNDW},
	{"BSFL", LTYPE3, x86.ABSFL},
	{"BSFQ", LTYPE3, x86.ABSFQ},
	{"BSFW", LTYPE3, x86.ABSFW},
	{"BSRL", LTYPE3, x86.ABSRL},
	{"BSRQ", LTYPE3, x86.ABSRQ},
	{"BSRW", LTYPE3, x86.ABSRW},
	{"BSWAPL", LTYPE1, x86.ABSWAPL},
	{"BSWAPQ", LTYPE1, x86.ABSWAPQ},
	{"BTCL", LTYPE3, x86.ABTCL},
	{"BTCQ", LTYPE3, x86.ABTCQ},
	{"BTCW", LTYPE3, x86.ABTCW},
	{"BTL", LTYPE3, x86.ABTL},
	{"BTQ", LTYPE3, x86.ABTQ},
	{"BTRL", LTYPE3, x86.ABTRL},
	{"BTRQ", LTYPE3, x86.ABTRQ},
	{"BTRW", LTYPE3, x86.ABTRW},
	{"BTSL", LTYPE3, x86.ABTSL},
	{"BTSQ", LTYPE3, x86.ABTSQ},
	{"BTSW", LTYPE3, x86.ABTSW},
	{"BTW", LTYPE3, x86.ABTW},
	{"BYTE", LTYPE2, x86.ABYTE},
	{"CALL", LTYPEC, obj.ACALL},
	{"CLC", LTYPE0, x86.ACLC},
	{"CLD", LTYPE0, x86.ACLD},
	{"CLI", LTYPE0, x86.ACLI},
	{"CLTS", LTYPE0, x86.ACLTS},
	{"CMC", LTYPE0, x86.ACMC},
	{"CMPB", LTYPE4, x86.ACMPB},
	{"CMPL", LTYPE4, x86.ACMPL},
	{"CMPQ", LTYPE4, x86.ACMPQ},
	{"CMPW", LTYPE4, x86.ACMPW},
	{"CMPSB", LTYPE0, x86.ACMPSB},
	{"CMPSL", LTYPE0, x86.ACMPSL},
	{"CMPSQ", LTYPE0, x86.ACMPSQ},
	{"CMPSW", LTYPE0, x86.ACMPSW},
	{"CMPXCHG8B", LTYPE1, x86.ACMPXCHG8B},
	{"CMPXCHGB", LTYPE3, x86.ACMPXCHGB}, /* LTYPE3? */
	{"CMPXCHGL", LTYPE3, x86.ACMPXCHGL},
	{"CMPXCHGQ", LTYPE3, x86.ACMPXCHGQ},
	{"CMPXCHGW", LTYPE3, x86.ACMPXCHGW},
	{"CPUID", LTYPE0, x86.ACPUID},
	{"DAA", LTYPE0, x86.ADAA},
	{"DAS", LTYPE0, x86.ADAS},
	{"DATA", LTYPED, obj.ADATA},
	{"DECB", LTYPE1, x86.ADECB},
	{"DECL", LTYPE1, x86.ADECL},
	{"DECQ", LTYPE1, x86.ADECQ},
	{"DECW", LTYPE1, x86.ADECW},
	{"DIVB", LTYPE2, x86.ADIVB},
	{"DIVL", LTYPE2, x86.ADIVL},
	{"DIVQ", LTYPE2, x86.ADIVQ},
	{"DIVW", LTYPE2, x86.ADIVW},
	{"EMMS", LTYPE0, x86.AEMMS},
	{"END", LTYPE0, obj.AEND},
	{"ENTER", LTYPE2, x86.AENTER},
	{"GLOBL", LTYPEG, obj.AGLOBL},
	{"HLT", LTYPE0, x86.AHLT},
	{"IDIVB", LTYPE2, x86.AIDIVB},
	{"IDIVL", LTYPE2, x86.AIDIVL},
	{"IDIVQ", LTYPE2, x86.AIDIVQ},
	{"IDIVW", LTYPE2, x86.AIDIVW},
	{"IMULB", LTYPEI, x86.AIMULB},
	{"IMULL", LTYPEI, x86.AIMULL},
	{"IMULQ", LTYPEI, x86.AIMULQ},
	{"IMUL3Q", LTYPEX, x86.AIMUL3Q},
	{"IMULW", LTYPEI, x86.AIMULW},
	{"INB", LTYPE0, x86.AINB},
	{"INL", LTYPE0, x86.AINL},
	{"INW", LTYPE0, x86.AINW},
	{"INCB", LTYPE1, x86.AINCB},
	{"INCL", LTYPE1, x86.AINCL},
	{"INCQ", LTYPE1, x86.AINCQ},
	{"INCW", LTYPE1, x86.AINCW},
	{"INSB", LTYPE0, x86.AINSB},
	{"INSL", LTYPE0, x86.AINSL},
	{"INSW", LTYPE0, x86.AINSW},
	{"INT", LTYPE2, x86.AINT},
	{"INTO", LTYPE0, x86.AINTO},
	{"INVD", LTYPE0, x86.AINVD},
	{"INVLPG", LTYPE2, x86.AINVLPG},
	{"IRETL", LTYPE0, x86.AIRETL},
	{"IRETQ", LTYPE0, x86.AIRETQ},
	{"IRETW", LTYPE0, x86.AIRETW},
	{"JOS", LTYPER, x86.AJOS},  /* overflow set (OF = 1) */
	{"JO", LTYPER, x86.AJOS},   /* alternate */
	{"JOC", LTYPER, x86.AJOC},  /* overflow clear (OF = 0) */
	{"JNO", LTYPER, x86.AJOC},  /* alternate */
	{"JCS", LTYPER, x86.AJCS},  /* carry set (CF = 1) */
	{"JB", LTYPER, x86.AJCS},   /* alternate */
	{"JC", LTYPER, x86.AJCS},   /* alternate */
	{"JNAE", LTYPER, x86.AJCS}, /* alternate */
	{"JLO", LTYPER, x86.AJCS},  /* alternate */
	{"JCC", LTYPER, x86.AJCC},  /* carry clear (CF = 0) */
	{"JAE", LTYPER, x86.AJCC},  /* alternate */
	{"JNB", LTYPER, x86.AJCC},  /* alternate */
	{"JNC", LTYPER, x86.AJCC},  /* alternate */
	{"JHS", LTYPER, x86.AJCC},  /* alternate */
	{"JEQ", LTYPER, x86.AJEQ},  /* equal (ZF = 1) */
	{"JE", LTYPER, x86.AJEQ},   /* alternate */
	{"JZ", LTYPER, x86.AJEQ},   /* alternate */
	{"JNE", LTYPER, x86.AJNE},  /* not equal (ZF = 0) */
	{"JNZ", LTYPER, x86.AJNE},  /* alternate */
	{"JLS", LTYPER, x86.AJLS},  /* lower or same (unsigned) (CF = 1 || ZF = 1) */
	{"JBE", LTYPER, x86.AJLS},  /* alternate */
	{"JNA", LTYPER, x86.AJLS},  /* alternate */
	{"JHI", LTYPER, x86.AJHI},  /* higher (unsigned) (CF = 0 && ZF = 0) */
	{"JA", LTYPER, x86.AJHI},   /* alternate */
	{"JNBE", LTYPER, x86.AJHI}, /* alternate */
	{"JMI", LTYPER, x86.AJMI},  /* negative (minus) (SF = 1) */
	{"JS", LTYPER, x86.AJMI},   /* alternate */
	{"JPL", LTYPER, x86.AJPL},  /* non-negative (plus) (SF = 0) */
	{"JNS", LTYPER, x86.AJPL},  /* alternate */
	{"JPS", LTYPER, x86.AJPS},  /* parity set (PF = 1) */
	{"JP", LTYPER, x86.AJPS},   /* alternate */
	{"JPE", LTYPER, x86.AJPS},  /* alternate */
	{"JPC", LTYPER, x86.AJPC},  /* parity clear (PF = 0) */
	{"JNP", LTYPER, x86.AJPC},  /* alternate */
	{"JPO", LTYPER, x86.AJPC},  /* alternate */
	{"JLT", LTYPER, x86.AJLT},  /* less than (signed) (SF != OF) */
	{"JL", LTYPER, x86.AJLT},   /* alternate */
	{"JNGE", LTYPER, x86.AJLT}, /* alternate */
	{"JGE", LTYPER, x86.AJGE},  /* greater than or equal (signed) (SF = OF) */
	{"JNL", LTYPER, x86.AJGE},  /* alternate */
	{"JLE", LTYPER, x86.AJLE},  /* less than or equal (signed) (ZF = 1 || SF != OF) */
	{"JNG", LTYPER, x86.AJLE},  /* alternate */
	{"JGT", LTYPER, x86.AJGT},  /* greater than (signed) (ZF = 0 && SF = OF) */
	{"JG", LTYPER, x86.AJGT},   /* alternate */
	{"JNLE", LTYPER, x86.AJGT}, /* alternate */
	{"JCXZL", LTYPER, x86.AJCXZL},
	{"JCXZQ", LTYPER, x86.AJCXZQ},
	{"JMP", LTYPEC, obj.AJMP},
	{"LAHF", LTYPE0, x86.ALAHF},
	{"LARL", LTYPE3, x86.ALARL},
	{"LARW", LTYPE3, x86.ALARW},
	{"LEAL", LTYPE3, x86.ALEAL},
	{"LEAQ", LTYPE3, x86.ALEAQ},
	{"LEAW", LTYPE3, x86.ALEAW},
	{"LEAVEL", LTYPE0, x86.ALEAVEL},
	{"LEAVEQ", LTYPE0, x86.ALEAVEQ},
	{"LEAVEW", LTYPE0, x86.ALEAVEW},
	{"LFENCE", LTYPE0, x86.ALFENCE},
	{"LOCK", LTYPE0, x86.ALOCK},
	{"LODSB", LTYPE0, x86.ALODSB},
	{"LODSL", LTYPE0, x86.ALODSL},
	{"LODSQ", LTYPE0, x86.ALODSQ},
	{"LODSW", LTYPE0, x86.ALODSW},
	{"LONG", LTYPE2, x86.ALONG},
	{"LOOP", LTYPER, x86.ALOOP},
	{"LOOPEQ", LTYPER, x86.ALOOPEQ},
	{"LOOPNE", LTYPER, x86.ALOOPNE},
	{"LSLL", LTYPE3, x86.ALSLL},
	{"LSLW", LTYPE3, x86.ALSLW},
	{"MFENCE", LTYPE0, x86.AMFENCE},
	{"MODE", LTYPE2, x86.AMODE},
	{"MOVB", LTYPE3, x86.AMOVB},
	{"MOVL", LTYPEM, x86.AMOVL},
	{"MOVQ", LTYPEM, x86.AMOVQ},
	{"MOVW", LTYPEM, x86.AMOVW},
	{"MOVBLSX", LTYPE3, x86.AMOVBLSX},
	{"MOVBLZX", LTYPE3, x86.AMOVBLZX},
	{"MOVBQSX", LTYPE3, x86.AMOVBQSX},
	{"MOVBQZX", LTYPE3, x86.AMOVBQZX},
	{"MOVBWSX", LTYPE3, x86.AMOVBWSX},
	{"MOVBWZX", LTYPE3, x86.AMOVBWZX},
	{"MOVLQSX", LTYPE3, x86.AMOVLQSX},
	{"MOVLQZX", LTYPE3, x86.AMOVLQZX},
	{"MOVNTIL", LTYPE3, x86.AMOVNTIL},
	{"MOVNTIQ", LTYPE3, x86.AMOVNTIQ},
	{"MOVQL", LTYPE3, x86.AMOVQL},
	{"MOVWLSX", LTYPE3, x86.AMOVWLSX},
	{"MOVWLZX", LTYPE3, x86.AMOVWLZX},
	{"MOVWQSX", LTYPE3, x86.AMOVWQSX},
	{"MOVWQZX", LTYPE3, x86.AMOVWQZX},
	{"MOVSB", LTYPE0, x86.AMOVSB},
	{"MOVSL", LTYPE0, x86.AMOVSL},
	{"MOVSQ", LTYPE0, x86.AMOVSQ},
	{"MOVSW", LTYPE0, x86.AMOVSW},
	{"MULB", LTYPE2, x86.AMULB},
	{"MULL", LTYPE2, x86.AMULL},
	{"MULQ", LTYPE2, x86.AMULQ},
	{"MULW", LTYPE2, x86.AMULW},
	{"NEGB", LTYPE1, x86.ANEGB},
	{"NEGL", LTYPE1, x86.ANEGL},
	{"NEGQ", LTYPE1, x86.ANEGQ},
	{"NEGW", LTYPE1, x86.ANEGW},
	{"NOP", LTYPEN, obj.ANOP},
	{"NOTB", LTYPE1, x86.ANOTB},
	{"NOTL", LTYPE1, x86.ANOTL},
	{"NOTQ", LTYPE1, x86.ANOTQ},
	{"NOTW", LTYPE1, x86.ANOTW},
	{"ORB", LTYPE3, x86.AORB},
	{"ORL", LTYPE3, x86.AORL},
	{"ORQ", LTYPE3, x86.AORQ},
	{"ORW", LTYPE3, x86.AORW},
	{"OUTB", LTYPE0, x86.AOUTB},
	{"OUTL", LTYPE0, x86.AOUTL},
	{"OUTW", LTYPE0, x86.AOUTW},
	{"OUTSB", LTYPE0, x86.AOUTSB},
	{"OUTSL", LTYPE0, x86.AOUTSL},
	{"OUTSW", LTYPE0, x86.AOUTSW},
	{"PAUSE", LTYPEN, x86.APAUSE},
	{"POPAL", LTYPE0, x86.APOPAL},
	{"POPAW", LTYPE0, x86.APOPAW},
	{"POPFL", LTYPE0, x86.APOPFL},
	{"POPFQ", LTYPE0, x86.APOPFQ},
	{"POPFW", LTYPE0, x86.APOPFW},
	{"POPL", LTYPE1, x86.APOPL},
	{"POPQ", LTYPE1, x86.APOPQ},
	{"POPW", LTYPE1, x86.APOPW},
	{"PUSHAL", LTYPE0, x86.APUSHAL},
	{"PUSHAW", LTYPE0, x86.APUSHAW},
	{"PUSHFL", LTYPE0, x86.APUSHFL},
	{"PUSHFQ", LTYPE0, x86.APUSHFQ},
	{"PUSHFW", LTYPE0, x86.APUSHFW},
	{"PUSHL", LTYPE2, x86.APUSHL},
	{"PUSHQ", LTYPE2, x86.APUSHQ},
	{"PUSHW", LTYPE2, x86.APUSHW},
	{"RCLB", LTYPE3, x86.ARCLB},
	{"RCLL", LTYPE3, x86.ARCLL},
	{"RCLQ", LTYPE3, x86.ARCLQ},
	{"RCLW", LTYPE3, x86.ARCLW},
	{"RCRB", LTYPE3, x86.ARCRB},
	{"RCRL", LTYPE3, x86.ARCRL},
	{"RCRQ", LTYPE3, x86.ARCRQ},
	{"RCRW", LTYPE3, x86.ARCRW},
	{"RDMSR", LTYPE0, x86.ARDMSR},
	{"RDPMC", LTYPE0, x86.ARDPMC},
	{"RDTSC", LTYPE0, x86.ARDTSC},
	{"REP", LTYPE0, x86.AREP},
	{"REPN", LTYPE0, x86.AREPN},
	{"RET", LTYPE0, obj.ARET},
	{"RETFL", LTYPERT, x86.ARETFL},
	{"RETFW", LTYPERT, x86.ARETFW},
	{"RETFQ", LTYPERT, x86.ARETFQ},
	{"ROLB", LTYPE3, x86.AROLB},
	{"ROLL", LTYPE3, x86.AROLL},
	{"ROLQ", LTYPE3, x86.AROLQ},
	{"ROLW", LTYPE3, x86.AROLW},
	{"RORB", LTYPE3, x86.ARORB},
	{"RORL", LTYPE3, x86.ARORL},
	{"RORQ", LTYPE3, x86.ARORQ},
	{"RORW", LTYPE3, x86.ARORW},
	{"RSM", LTYPE0, x86.ARSM},
	{"SAHF", LTYPE0, x86.ASAHF},
	{"SALB", LTYPE3, x86.ASALB},
	{"SALL", LTYPE3, x86.ASALL},
	{"SALQ", LTYPE3, x86.ASALQ},
	{"SALW", LTYPE3, x86.ASALW},
	{"SARB", LTYPE3, x86.ASARB},
	{"SARL", LTYPE3, x86.ASARL},
	{"SARQ", LTYPE3, x86.ASARQ},
	{"SARW", LTYPE3, x86.ASARW},
	{"SBBB", LTYPE3, x86.ASBBB},
	{"SBBL", LTYPE3, x86.ASBBL},
	{"SBBQ", LTYPE3, x86.ASBBQ},
	{"SBBW", LTYPE3, x86.ASBBW},
	{"SCASB", LTYPE0, x86.ASCASB},
	{"SCASL", LTYPE0, x86.ASCASL},
	{"SCASQ", LTYPE0, x86.ASCASQ},
	{"SCASW", LTYPE0, x86.ASCASW},
	{"SETCC", LTYPE1, x86.ASETCC}, /* see JCC etc above for condition codes */
	{"SETCS", LTYPE1, x86.ASETCS},
	{"SETEQ", LTYPE1, x86.ASETEQ},
	{"SETGE", LTYPE1, x86.ASETGE},
	{"SETGT", LTYPE1, x86.ASETGT},
	{"SETHI", LTYPE1, x86.ASETHI},
	{"SETLE", LTYPE1, x86.ASETLE},
	{"SETLS", LTYPE1, x86.ASETLS},
	{"SETLT", LTYPE1, x86.ASETLT},
	{"SETMI", LTYPE1, x86.ASETMI},
	{"SETNE", LTYPE1, x86.ASETNE},
	{"SETOC", LTYPE1, x86.ASETOC},
	{"SETOS", LTYPE1, x86.ASETOS},
	{"SETPC", LTYPE1, x86.ASETPC},
	{"SETPL", LTYPE1, x86.ASETPL},
	{"SETPS", LTYPE1, x86.ASETPS},
	{"SFENCE", LTYPE0, x86.ASFENCE},
	{"CDQ", LTYPE0, x86.ACDQ},
	{"CWD", LTYPE0, x86.ACWD},
	{"CQO", LTYPE0, x86.ACQO},
	{"SHLB", LTYPE3, x86.ASHLB},
	{"SHLL", LTYPES, x86.ASHLL},
	{"SHLQ", LTYPES, x86.ASHLQ},
	{"SHLW", LTYPES, x86.ASHLW},
	{"SHRB", LTYPE3, x86.ASHRB},
	{"SHRL", LTYPES, x86.ASHRL},
	{"SHRQ", LTYPES, x86.ASHRQ},
	{"SHRW", LTYPES, x86.ASHRW},
	{"STC", LTYPE0, x86.ASTC},
	{"STD", LTYPE0, x86.ASTD},
	{"STI", LTYPE0, x86.ASTI},
	{"STOSB", LTYPE0, x86.ASTOSB},
	{"STOSL", LTYPE0, x86.ASTOSL},
	{"STOSQ", LTYPE0, x86.ASTOSQ},
	{"STOSW", LTYPE0, x86.ASTOSW},
	{"SUBB", LTYPE3, x86.ASUBB},
	{"SUBL", LTYPE3, x86.ASUBL},
	{"SUBQ", LTYPE3, x86.ASUBQ},
	{"SUBW", LTYPE3, x86.ASUBW},
	{"SYSCALL", LTYPE0, x86.ASYSCALL},
	{"SYSRET", LTYPE0, x86.ASYSRET},
	{"SWAPGS", LTYPE0, x86.ASWAPGS},
	{"TESTB", LTYPE3, x86.ATESTB},
	{"TESTL", LTYPE3, x86.ATESTL},
	{"TESTQ", LTYPE3, x86.ATESTQ},
	{"TESTW", LTYPE3, x86.ATESTW},
	{"TEXT", LTYPET, obj.ATEXT},
	{"VERR", LTYPE2, x86.AVERR},
	{"VERW", LTYPE2, x86.AVERW},
	{"QUAD", LTYPE2, x86.AQUAD},
	{"WAIT", LTYPE0, x86.AWAIT},
	{"WBINVD", LTYPE0, x86.AWBINVD},
	{"WRMSR", LTYPE0, x86.AWRMSR},
	{"WORD", LTYPE2, x86.AWORD},
	{"XADDB", LTYPE3, x86.AXADDB},
	{"XADDL", LTYPE3, x86.AXADDL},
	{"XADDQ", LTYPE3, x86.AXADDQ},
	{"XADDW", LTYPE3, x86.AXADDW},
	{"XCHGB", LTYPE3, x86.AXCHGB},
	{"XCHGL", LTYPE3, x86.AXCHGL},
	{"XCHGQ", LTYPE3, x86.AXCHGQ},
	{"XCHGW", LTYPE3, x86.AXCHGW},
	{"XLAT", LTYPE2, x86.AXLAT},
	{"XORB", LTYPE3, x86.AXORB},
	{"XORL", LTYPE3, x86.AXORL},
	{"XORQ", LTYPE3, x86.AXORQ},
	{"XORW", LTYPE3, x86.AXORW},
	{"CMOVLCC", LTYPE3, x86.ACMOVLCC},
	{"CMOVLCS", LTYPE3, x86.ACMOVLCS},
	{"CMOVLEQ", LTYPE3, x86.ACMOVLEQ},
	{"CMOVLGE", LTYPE3, x86.ACMOVLGE},
	{"CMOVLGT", LTYPE3, x86.ACMOVLGT},
	{"CMOVLHI", LTYPE3, x86.ACMOVLHI},
	{"CMOVLLE", LTYPE3, x86.ACMOVLLE},
	{"CMOVLLS", LTYPE3, x86.ACMOVLLS},
	{"CMOVLLT", LTYPE3, x86.ACMOVLLT},
	{"CMOVLMI", LTYPE3, x86.ACMOVLMI},
	{"CMOVLNE", LTYPE3, x86.ACMOVLNE},
	{"CMOVLOC", LTYPE3, x86.ACMOVLOC},
	{"CMOVLOS", LTYPE3, x86.ACMOVLOS},
	{"CMOVLPC", LTYPE3, x86.ACMOVLPC},
	{"CMOVLPL", LTYPE3, x86.ACMOVLPL},
	{"CMOVLPS", LTYPE3, x86.ACMOVLPS},
	{"CMOVQCC", LTYPE3, x86.ACMOVQCC},
	{"CMOVQCS", LTYPE3, x86.ACMOVQCS},
	{"CMOVQEQ", LTYPE3, x86.ACMOVQEQ},
	{"CMOVQGE", LTYPE3, x86.ACMOVQGE},
	{"CMOVQGT", LTYPE3, x86.ACMOVQGT},
	{"CMOVQHI", LTYPE3, x86.ACMOVQHI},
	{"CMOVQLE", LTYPE3, x86.ACMOVQLE},
	{"CMOVQLS", LTYPE3, x86.ACMOVQLS},
	{"CMOVQLT", LTYPE3, x86.ACMOVQLT},
	{"CMOVQMI", LTYPE3, x86.ACMOVQMI},
	{"CMOVQNE", LTYPE3, x86.ACMOVQNE},
	{"CMOVQOC", LTYPE3, x86.ACMOVQOC},
	{"CMOVQOS", LTYPE3, x86.ACMOVQOS},
	{"CMOVQPC", LTYPE3, x86.ACMOVQPC},
	{"CMOVQPL", LTYPE3, x86.ACMOVQPL},
	{"CMOVQPS", LTYPE3, x86.ACMOVQPS},
	{"CMOVWCC", LTYPE3, x86.ACMOVWCC},
	{"CMOVWCS", LTYPE3, x86.ACMOVWCS},
	{"CMOVWEQ", LTYPE3, x86.ACMOVWEQ},
	{"CMOVWGE", LTYPE3, x86.ACMOVWGE},
	{"CMOVWGT", LTYPE3, x86.ACMOVWGT},
	{"CMOVWHI", LTYPE3, x86.ACMOVWHI},
	{"CMOVWLE", LTYPE3, x86.ACMOVWLE},
	{"CMOVWLS", LTYPE3, x86.ACMOVWLS},
	{"CMOVWLT", LTYPE3, x86.ACMOVWLT},
	{"CMOVWMI", LTYPE3, x86.ACMOVWMI},
	{"CMOVWNE", LTYPE3, x86.ACMOVWNE},
	{"CMOVWOC", LTYPE3, x86.ACMOVWOC},
	{"CMOVWOS", LTYPE3, x86.ACMOVWOS},
	{"CMOVWPC", LTYPE3, x86.ACMOVWPC},
	{"CMOVWPL", LTYPE3, x86.ACMOVWPL},
	{"CMOVWPS", LTYPE3, x86.ACMOVWPS},
	{"FMOVB", LTYPE3, x86.AFMOVB},
	{"FMOVBP", LTYPE3, x86.AFMOVBP},
	{"FMOVD", LTYPE3, x86.AFMOVD},
	{"FMOVDP", LTYPE3, x86.AFMOVDP},
	{"FMOVF", LTYPE3, x86.AFMOVF},
	{"FMOVFP", LTYPE3, x86.AFMOVFP},
	{"FMOVL", LTYPE3, x86.AFMOVL},
	{"FMOVLP", LTYPE3, x86.AFMOVLP},
	{"FMOVV", LTYPE3, x86.AFMOVV},
	{"FMOVVP", LTYPE3, x86.AFMOVVP},
	{"FMOVW", LTYPE3, x86.AFMOVW},
	{"FMOVWP", LTYPE3, x86.AFMOVWP},
	{"FMOVX", LTYPE3, x86.AFMOVX},
	{"FMOVXP", LTYPE3, x86.AFMOVXP},
	{"FCOMB", LTYPE3, x86.AFCOMB},
	{"FCOMBP", LTYPE3, x86.AFCOMBP},
	{"FCOMD", LTYPE3, x86.AFCOMD},
	{"FCOMDP", LTYPE3, x86.AFCOMDP},
	{"FCOMDPP", LTYPE3, x86.AFCOMDPP},
	{"FCOMF", LTYPE3, x86.AFCOMF},
	{"FCOMFP", LTYPE3, x86.AFCOMFP},
	{"FCOML", LTYPE3, x86.AFCOML},
	{"FCOMLP", LTYPE3, x86.AFCOMLP},
	{"FCOMW", LTYPE3, x86.AFCOMW},
	{"FCOMWP", LTYPE3, x86.AFCOMWP},
	{"FUCOM", LTYPE3, x86.AFUCOM},
	{"FUCOMP", LTYPE3, x86.AFUCOMP},
	{"FUCOMPP", LTYPE3, x86.AFUCOMPP},
	{"FADDW", LTYPE3, x86.AFADDW},
	{"FADDL", LTYPE3, x86.AFADDL},
	{"FADDF", LTYPE3, x86.AFADDF},
	{"FADDD", LTYPE3, x86.AFADDD},
	{"FADDDP", LTYPE3, x86.AFADDDP},
	{"FSUBDP", LTYPE3, x86.AFSUBDP},
	{"FSUBW", LTYPE3, x86.AFSUBW},
	{"FSUBL", LTYPE3, x86.AFSUBL},
	{"FSUBF", LTYPE3, x86.AFSUBF},
	{"FSUBD", LTYPE3, x86.AFSUBD},
	{"FSUBRDP", LTYPE3, x86.AFSUBRDP},
	{"FSUBRW", LTYPE3, x86.AFSUBRW},
	{"FSUBRL", LTYPE3, x86.AFSUBRL},
	{"FSUBRF", LTYPE3, x86.AFSUBRF},
	{"FSUBRD", LTYPE3, x86.AFSUBRD},
	{"FMULDP", LTYPE3, x86.AFMULDP},
	{"FMULW", LTYPE3, x86.AFMULW},
	{"FMULL", LTYPE3, x86.AFMULL},
	{"FMULF", LTYPE3, x86.AFMULF},
	{"FMULD", LTYPE3, x86.AFMULD},
	{"FDIVDP", LTYPE3, x86.AFDIVDP},
	{"FDIVW", LTYPE3, x86.AFDIVW},
	{"FDIVL", LTYPE3, x86.AFDIVL},
	{"FDIVF", LTYPE3, x86.AFDIVF},
	{"FDIVD", LTYPE3, x86.AFDIVD},
	{"FDIVRDP", LTYPE3, x86.AFDIVRDP},
	{"FDIVRW", LTYPE3, x86.AFDIVRW},
	{"FDIVRL", LTYPE3, x86.AFDIVRL},
	{"FDIVRF", LTYPE3, x86.AFDIVRF},
	{"FDIVRD", LTYPE3, x86.AFDIVRD},
	{"FXCHD", LTYPE3, x86.AFXCHD},
	{"FFREE", LTYPE1, x86.AFFREE},
	{"FLDCW", LTYPE2, x86.AFLDCW},
	{"FLDENV", LTYPE1, x86.AFLDENV},
	{"FRSTOR", LTYPE2, x86.AFRSTOR},
	{"FSAVE", LTYPE1, x86.AFSAVE},
	{"FSTCW", LTYPE1, x86.AFSTCW},
	{"FSTENV", LTYPE1, x86.AFSTENV},
	{"FSTSW", LTYPE1, x86.AFSTSW},
	{"F2XM1", LTYPE0, x86.AF2XM1},
	{"FABS", LTYPE0, x86.AFABS},
	{"FCHS", LTYPE0, x86.AFCHS},
	{"FCLEX", LTYPE0, x86.AFCLEX},
	{"FCOS", LTYPE0, x86.AFCOS},
	{"FDECSTP", LTYPE0, x86.AFDECSTP},
	{"FINCSTP", LTYPE0, x86.AFINCSTP},
	{"FINIT", LTYPE0, x86.AFINIT},
	{"FLD1", LTYPE0, x86.AFLD1},
	{"FLDL2E", LTYPE0, x86.AFLDL2E},
	{"FLDL2T", LTYPE0, x86.AFLDL2T},
	{"FLDLG2", LTYPE0, x86.AFLDLG2},
	{"FLDLN2", LTYPE0, x86.AFLDLN2},
	{"FLDPI", LTYPE0, x86.AFLDPI},
	{"FLDZ", LTYPE0, x86.AFLDZ},
	{"FNOP", LTYPE0, x86.AFNOP},
	{"FPATAN", LTYPE0, x86.AFPATAN},
	{"FPREM", LTYPE0, x86.AFPREM},
	{"FPREM1", LTYPE0, x86.AFPREM1},
	{"FPTAN", LTYPE0, x86.AFPTAN},
	{"FRNDINT", LTYPE0, x86.AFRNDINT},
	{"FSCALE", LTYPE0, x86.AFSCALE},
	{"FSIN", LTYPE0, x86.AFSIN},
	{"FSINCOS", LTYPE0, x86.AFSINCOS},
	{"FSQRT", LTYPE0, x86.AFSQRT},
	{"FTST", LTYPE0, x86.AFTST},
	{"FXAM", LTYPE0, x86.AFXAM},
	{"FXTRACT", LTYPE0, x86.AFXTRACT},
	{"FYL2X", LTYPE0, x86.AFYL2X},
	{"FYL2XP1", LTYPE0, x86.AFYL2XP1},
	{"ADDPD", LTYPE3, x86.AADDPD},
	{"ADDPS", LTYPE3, x86.AADDPS},
	{"ADDSD", LTYPE3, x86.AADDSD},
	{"ADDSS", LTYPE3, x86.AADDSS},
	{"ANDNPD", LTYPE3, x86.AANDNPD},
	{"ANDNPS", LTYPE3, x86.AANDNPS},
	{"ANDPD", LTYPE3, x86.AANDPD},
	{"ANDPS", LTYPE3, x86.AANDPS},
	{"CMPPD", LTYPEXC, x86.ACMPPD},
	{"CMPPS", LTYPEXC, x86.ACMPPS},
	{"CMPSD", LTYPEXC, x86.ACMPSD},
	{"CMPSS", LTYPEXC, x86.ACMPSS},
	{"COMISD", LTYPE3, x86.ACOMISD},
	{"COMISS", LTYPE3, x86.ACOMISS},
	{"CVTPL2PD", LTYPE3, x86.ACVTPL2PD},
	{"CVTPL2PS", LTYPE3, x86.ACVTPL2PS},
	{"CVTPD2PL", LTYPE3, x86.ACVTPD2PL},
	{"CVTPD2PS", LTYPE3, x86.ACVTPD2PS},
	{"CVTPS2PL", LTYPE3, x86.ACVTPS2PL},
	{"PF2IW", LTYPE3, x86.APF2IW},
	{"PF2IL", LTYPE3, x86.APF2IL},
	{"PF2ID", LTYPE3, x86.APF2IL}, /* syn */
	{"PI2FL", LTYPE3, x86.API2FL},
	{"PI2FD", LTYPE3, x86.API2FL}, /* syn */
	{"PI2FW", LTYPE3, x86.API2FW},
	{"CVTPS2PD", LTYPE3, x86.ACVTPS2PD},
	{"CVTSD2SL", LTYPE3, x86.ACVTSD2SL},
	{"CVTSD2SQ", LTYPE3, x86.ACVTSD2SQ},
	{"CVTSD2SS", LTYPE3, x86.ACVTSD2SS},
	{"CVTSL2SD", LTYPE3, x86.ACVTSL2SD},
	{"CVTSQ2SD", LTYPE3, x86.ACVTSQ2SD},
	{"CVTSL2SS", LTYPE3, x86.ACVTSL2SS},
	{"CVTSQ2SS", LTYPE3, x86.ACVTSQ2SS},
	{"CVTSS2SD", LTYPE3, x86.ACVTSS2SD},
	{"CVTSS2SL", LTYPE3, x86.ACVTSS2SL},
	{"CVTSS2SQ", LTYPE3, x86.ACVTSS2SQ},
	{"CVTTPD2PL", LTYPE3, x86.ACVTTPD2PL},
	{"CVTTPS2PL", LTYPE3, x86.ACVTTPS2PL},
	{"CVTTSD2SL", LTYPE3, x86.ACVTTSD2SL},
	{"CVTTSD2SQ", LTYPE3, x86.ACVTTSD2SQ},
	{"CVTTSS2SL", LTYPE3, x86.ACVTTSS2SL},
	{"CVTTSS2SQ", LTYPE3, x86.ACVTTSS2SQ},
	{"DIVPD", LTYPE3, x86.ADIVPD},
	{"DIVPS", LTYPE3, x86.ADIVPS},
	{"DIVSD", LTYPE3, x86.ADIVSD},
	{"DIVSS", LTYPE3, x86.ADIVSS},
	{"FXRSTOR", LTYPE2, x86.AFXRSTOR},
	{"FXRSTOR64", LTYPE2, x86.AFXRSTOR64},
	{"FXSAVE", LTYPE1, x86.AFXSAVE},
	{"FXSAVE64", LTYPE1, x86.AFXSAVE64},
	{"LDMXCSR", LTYPE2, x86.ALDMXCSR},
	{"MASKMOVOU", LTYPE3, x86.AMASKMOVOU},
	{"MASKMOVDQU", LTYPE3, x86.AMASKMOVOU}, /* syn */
	{"MASKMOVQ", LTYPE3, x86.AMASKMOVQ},
	{"MAXPD", LTYPE3, x86.AMAXPD},
	{"MAXPS", LTYPE3, x86.AMAXPS},
	{"MAXSD", LTYPE3, x86.AMAXSD},
	{"MAXSS", LTYPE3, x86.AMAXSS},
	{"MINPD", LTYPE3, x86.AMINPD},
	{"MINPS", LTYPE3, x86.AMINPS},
	{"MINSD", LTYPE3, x86.AMINSD},
	{"MINSS", LTYPE3, x86.AMINSS},
	{"MOVAPD", LTYPE3, x86.AMOVAPD},
	{"MOVAPS", LTYPE3, x86.AMOVAPS},
	{"MOVD", LTYPE3, x86.AMOVQ},    /* syn */
	{"MOVDQ2Q", LTYPE3, x86.AMOVQ}, /* syn */
	{"MOVO", LTYPE3, x86.AMOVO},
	{"MOVOA", LTYPE3, x86.AMOVO}, /* syn */
	{"MOVOU", LTYPE3, x86.AMOVOU},
	{"MOVHLPS", LTYPE3, x86.AMOVHLPS},
	{"MOVHPD", LTYPE3, x86.AMOVHPD},
	{"MOVHPS", LTYPE3, x86.AMOVHPS},
	{"MOVLHPS", LTYPE3, x86.AMOVLHPS},
	{"MOVLPD", LTYPE3, x86.AMOVLPD},
	{"MOVLPS", LTYPE3, x86.AMOVLPS},
	{"MOVMSKPD", LTYPE3, x86.AMOVMSKPD},
	{"MOVMSKPS", LTYPE3, x86.AMOVMSKPS},
	{"MOVNTO", LTYPE3, x86.AMOVNTO},
	{"MOVNTDQ", LTYPE3, x86.AMOVNTO}, /* syn */
	{"MOVNTPD", LTYPE3, x86.AMOVNTPD},
	{"MOVNTPS", LTYPE3, x86.AMOVNTPS},
	{"MOVNTQ", LTYPE3, x86.AMOVNTQ},
	{"MOVQOZX", LTYPE3, x86.AMOVQOZX},
	{"MOVSD", LTYPE3, x86.AMOVSD},
	{"MOVSS", LTYPE3, x86.AMOVSS},
	{"MOVUPD", LTYPE3, x86.AMOVUPD},
	{"MOVUPS", LTYPE3, x86.AMOVUPS},
	{"MULPD", LTYPE3, x86.AMULPD},
	{"MULPS", LTYPE3, x86.AMULPS},
	{"MULSD", LTYPE3, x86.AMULSD},
	{"MULSS", LTYPE3, x86.AMULSS},
	{"ORPD", LTYPE3, x86.AORPD},
	{"ORPS", LTYPE3, x86.AORPS},
	{"PACKSSLW", LTYPE3, x86.APACKSSLW},
	{"PACKSSWB", LTYPE3, x86.APACKSSWB},
	{"PACKUSWB", LTYPE3, x86.APACKUSWB},
	{"PADDB", LTYPE3, x86.APADDB},
	{"PADDL", LTYPE3, x86.APADDL},
	{"PADDQ", LTYPE3, x86.APADDQ},
	{"PADDSB", LTYPE3, x86.APADDSB},
	{"PADDSW", LTYPE3, x86.APADDSW},
	{"PADDUSB", LTYPE3, x86.APADDUSB},
	{"PADDUSW", LTYPE3, x86.APADDUSW},
	{"PADDW", LTYPE3, x86.APADDW},
	{"PAND", LTYPE3, x86.APAND},
	{"PANDB", LTYPE3, x86.APANDB},
	{"PANDL", LTYPE3, x86.APANDL},
	{"PANDSB", LTYPE3, x86.APANDSB},
	{"PANDSW", LTYPE3, x86.APANDSW},
	{"PANDUSB", LTYPE3, x86.APANDUSB},
	{"PANDUSW", LTYPE3, x86.APANDUSW},
	{"PANDW", LTYPE3, x86.APANDW},
	{"PANDN", LTYPE3, x86.APANDN},
	{"PAVGB", LTYPE3, x86.APAVGB},
	{"PAVGW", LTYPE3, x86.APAVGW},
	{"PCMPEQB", LTYPE3, x86.APCMPEQB},
	{"PCMPEQL", LTYPE3, x86.APCMPEQL},
	{"PCMPEQW", LTYPE3, x86.APCMPEQW},
	{"PCMPGTB", LTYPE3, x86.APCMPGTB},
	{"PCMPGTL", LTYPE3, x86.APCMPGTL},
	{"PCMPGTW", LTYPE3, x86.APCMPGTW},
	{"PEXTRW", LTYPEX, x86.APEXTRW},
	{"PINSRW", LTYPEX, x86.APINSRW},
	{"PINSRD", LTYPEX, x86.APINSRD},
	{"PINSRQ", LTYPEX, x86.APINSRQ},
	{"PMADDWL", LTYPE3, x86.APMADDWL},
	{"PMAXSW", LTYPE3, x86.APMAXSW},
	{"PMAXUB", LTYPE3, x86.APMAXUB},
	{"PMINSW", LTYPE3, x86.APMINSW},
	{"PMINUB", LTYPE3, x86.APMINUB},
	{"PMOVMSKB", LTYPE3, x86.APMOVMSKB},
	{"PMULHRW", LTYPE3, x86.APMULHRW},
	{"PMULHUW", LTYPE3, x86.APMULHUW},
	{"PMULHW", LTYPE3, x86.APMULHW},
	{"PMULLW", LTYPE3, x86.APMULLW},
	{"PMULULQ", LTYPE3, x86.APMULULQ},
	{"POR", LTYPE3, x86.APOR},
	{"PSADBW", LTYPE3, x86.APSADBW},
	{"PSHUFHW", LTYPEX, x86.APSHUFHW},
	{"PSHUFL", LTYPEX, x86.APSHUFL},
	{"PSHUFLW", LTYPEX, x86.APSHUFLW},
	{"PSHUFW", LTYPEX, x86.APSHUFW},
	{"PSHUFB", LTYPEM, x86.APSHUFB},
	{"PSLLO", LTYPE3, x86.APSLLO},
	{"PSLLDQ", LTYPE3, x86.APSLLO}, /* syn */
	{"PSLLL", LTYPE3, x86.APSLLL},
	{"PSLLQ", LTYPE3, x86.APSLLQ},
	{"PSLLW", LTYPE3, x86.APSLLW},
	{"PSRAL", LTYPE3, x86.APSRAL},
	{"PSRAW", LTYPE3, x86.APSRAW},
	{"PSRLO", LTYPE3, x86.APSRLO},
	{"PSRLDQ", LTYPE3, x86.APSRLO}, /* syn */
	{"PSRLL", LTYPE3, x86.APSRLL},
	{"PSRLQ", LTYPE3, x86.APSRLQ},
	{"PSRLW", LTYPE3, x86.APSRLW},
	{"PSUBB", LTYPE3, x86.APSUBB},
	{"PSUBL", LTYPE3, x86.APSUBL},
	{"PSUBQ", LTYPE3, x86.APSUBQ},
	{"PSUBSB", LTYPE3, x86.APSUBSB},
	{"PSUBSW", LTYPE3, x86.APSUBSW},
	{"PSUBUSB", LTYPE3, x86.APSUBUSB},
	{"PSUBUSW", LTYPE3, x86.APSUBUSW},
	{"PSUBW", LTYPE3, x86.APSUBW},
	{"PUNPCKHBW", LTYPE3, x86.APUNPCKHBW},
	{"PUNPCKHLQ", LTYPE3, x86.APUNPCKHLQ},
	{"PUNPCKHQDQ", LTYPE3, x86.APUNPCKHQDQ},
	{"PUNPCKHWL", LTYPE3, x86.APUNPCKHWL},
	{"PUNPCKLBW", LTYPE3, x86.APUNPCKLBW},
	{"PUNPCKLLQ", LTYPE3, x86.APUNPCKLLQ},
	{"PUNPCKLQDQ", LTYPE3, x86.APUNPCKLQDQ},
	{"PUNPCKLWL", LTYPE3, x86.APUNPCKLWL},
	{"PXOR", LTYPE3, x86.APXOR},
	{"RCPPS", LTYPE3, x86.ARCPPS},
	{"RCPSS", LTYPE3, x86.ARCPSS},
	{"RSQRTPS", LTYPE3, x86.ARSQRTPS},
	{"RSQRTSS", LTYPE3, x86.ARSQRTSS},
	{"SHUFPD", LTYPEX, x86.ASHUFPD},
	{"SHUFPS", LTYPEX, x86.ASHUFPS},
	{"SQRTPD", LTYPE3, x86.ASQRTPD},
	{"SQRTPS", LTYPE3, x86.ASQRTPS},
	{"SQRTSD", LTYPE3, x86.ASQRTSD},
	{"SQRTSS", LTYPE3, x86.ASQRTSS},
	{"STMXCSR", LTYPE1, x86.ASTMXCSR},
	{"SUBPD", LTYPE3, x86.ASUBPD},
	{"SUBPS", LTYPE3, x86.ASUBPS},
	{"SUBSD", LTYPE3, x86.ASUBSD},
	{"SUBSS", LTYPE3, x86.ASUBSS},
	{"UCOMISD", LTYPE3, x86.AUCOMISD},
	{"UCOMISS", LTYPE3, x86.AUCOMISS},
	{"UNPCKHPD", LTYPE3, x86.AUNPCKHPD},
	{"UNPCKHPS", LTYPE3, x86.AUNPCKHPS},
	{"UNPCKLPD", LTYPE3, x86.AUNPCKLPD},
	{"UNPCKLPS", LTYPE3, x86.AUNPCKLPS},
	{"XORPD", LTYPE3, x86.AXORPD},
	{"XORPS", LTYPE3, x86.AXORPS},
	{"CRC32B", LTYPE4, x86.ACRC32B},
	{"CRC32Q", LTYPE4, x86.ACRC32Q},
	{"PREFETCHT0", LTYPE2, x86.APREFETCHT0},
	{"PREFETCHT1", LTYPE2, x86.APREFETCHT1},
	{"PREFETCHT2", LTYPE2, x86.APREFETCHT2},
	{"PREFETCHNTA", LTYPE2, x86.APREFETCHNTA},
	{"UNDEF", LTYPE0, obj.AUNDEF},
	{"AESENC", LTYPE3, x86.AAESENC},
	{"AESENCLAST", LTYPE3, x86.AAESENCLAST},
	{"AESDEC", LTYPE3, x86.AAESDEC},
	{"AESDECLAST", LTYPE3, x86.AAESDECLAST},
	{"AESIMC", LTYPE3, x86.AAESIMC},
	{"AESKEYGENASSIST", LTYPEX, x86.AAESKEYGENASSIST},
	{"PSHUFD", LTYPEX, x86.APSHUFD},
	{"USEFIELD", LTYPEN, obj.AUSEFIELD},
	{"PCLMULQDQ", LTYPEX, x86.APCLMULQDQ},
	{"PCDATA", LTYPEPC, obj.APCDATA},
	{"FUNCDATA", LTYPEF, obj.AFUNCDATA},
}

func cinit() {
}

type yy struct{}

func (yy) Error(msg string) {
	asm.Yyerror("%s", msg)
}

func (yy) Lex(v *yySymType) int {
	var av asm.Yylval
	tok := asm.Yylex(&av)
	v.sym = av.Sym
	v.lval = av.Lval
	v.sval = av.Sval
	v.dval = av.Dval
	return tok
}

func yyparse() {
	yyParse(yy{})
}

func main() {
	cinit()

	asm.LSCONST = LSCONST
	asm.LCONST = LCONST
	asm.LFCONST = LFCONST
	asm.LNAME = LNAME
	asm.LVAR = LVAR
	asm.LLAB = LLAB

	asm.Thechar = '6'
	asm.Thestring = "amd64"
	asm.Thelinkarch = &x86.Linkamd64
	asm.Arches = map[string]*obj.LinkArch{
		"amd64p32": &x86.Linkamd64p32,
	}

	asm.Lexinit = lexinit
	asm.Cclean = cclean
	asm.Yyparse = yyparse

	asm.Main()
}
