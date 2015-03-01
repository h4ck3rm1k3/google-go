// Inferno utils/6a/a.h and lex.c.
// http://code.google.com/p/inferno-os/source/browse/utils/6a/a.h
// http://code.google.com/p/inferno-os/source/browse/utils/6a/lex.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.	All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AoND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package asm holds code shared among the assemblers.
package asm

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"cmd/internal/obj"
	"bytes"
	"unicode/utf8"

)
func Yyerror3() {}
func Yyerror2(a string) {}
func Yyerror(a string, args ...interface{}) {
	/*
	 * hack to intercept message from yaccpar
	 */
	if a == "syntax error" || len(args) == 1 && a == "%s" && args[0] == "syntax error" {
		Yyerror("syntax error, last name: %s", last)
		return
	}

	prfile(Lineno)
	fmt.Printf("%s\n", fmt.Sprintf(a, args...))
	nerrors++
	if nerrors > 10 {
		fmt.Printf("too many errors\n")
		errorexit()
	}
}

// Initialized by client.
var (
	LSCONST int
	LCONST  int
	LFCONST int
	LNAME   int
	LVAR    int
	LLAB    int

	Thechar     rune
	Thestring   string
	Thelinkarch *obj.LinkArch

	Arches map[string]*obj.LinkArch

	Cclean  func()
	Yyparse func()
	Syminit func(*Sym)

	Lexinit []Lextab
)

type Lextab struct {
	Name  string
	Type  int
	Value int64
}

const (
	MAXALIGN = 7
	FPCHIP   = 1
	NSYMB    = 500
	BUFSIZ   = 8192
	HISTSZ   = 20
	EOF      = -1
	IGN      = -2
	NHASH    = 503
	STRINGSZ = 200
	NMACRO   = 10
)

const (
	CLAST = iota
	CMACARG
	CMACRO
	CPREPROC
)

type Macro struct {
	Text string
	Narg int
	Dots bool
}

type Sym struct {
	Link      *Sym
	Ref       *Ref
	Macro     *Macro
	Value     int64
	Type      int
	Name      string
	Labelname string
	Sym       int8
}

type Ref struct {
	Class int
}

type Io struct {
	Link *Io
	P    []byte
	F    *os.File
	B    [1024]byte
}

var fi struct {
	P []byte
}

func LabelLookup(s *Sym) *Sym {
	var p string
	var lab *Sym

	if thetext == nil {
		s.Labelname = s.Name
		return s
	}

	p = string(fmt.Sprintf("%s.%s", thetext.Name, s.Name))
	lab = Lookup(p)

	lab.Labelname = s.Name
	return lab

}

var 	include  []string;
var (
	debug    [256]int
	hash     = map[string]*Sym{}
	Dlist    []string
	newflag  int
	hunk     string
//	include  []string
	iofree   *Io
	ionext   *Io
	iostack  *Io
	Lineno   int32
	nerrors  int
	nhunk    int32
	ninclude int
	nsymb    int32
	nullgen  obj.Addr
	outfile  string
	Pass     int
	PC       int32
	peekc    int = IGN
	sym      int
	symb     string
	thunk    int32
	obuf     obj.Biobuf
	Ctxt     *obj.Link
	bstdout  obj.Biobuf
)

func setinclude(p string) {
	var i int
	if p == "" {
		return
	}
	for i = 1; i < len(include); i++ {
		if p == include[i] {
			return
		}
	}
	include = append(include, p)
}


func Lookup(symb string) *Sym {
	// turn leading · into ""·
	if strings.HasPrefix(symb, "·") {
		symb = `""` + symb
	}

	// turn · (U+00B7) into .
	// turn ∕ (U+2215) into /
	symb = strings.Replace(symb, "·", ".", -1)
	symb = strings.Replace(symb, "∕", "/", -1)

	s := hash[symb]
	if s != nil {
		return s
	}

	s = new(Sym)
	s.Name = symb
	syminit(s)
	hash[symb] = s
	return s
}

func errorexit() {
	obj.Bflush(&bstdout)
	if outfile != "" {
		os.Remove(outfile)
	}
	os.Exit(2)
}

func pushio() {
	var i *Io

	i = iostack
	if i == nil {
		Yyerror("botch in pushio")
		errorexit()
	}

	i.P = fi.P
}

func newio() {
	var i *Io
	var pushdepth int = 0

	i = iofree
	if i == nil {
		pushdepth++
		if pushdepth > 1000 {
			Yyerror("macro/io expansion too deep")
			errorexit()
		}
		i = new(Io)
	} else {
		iofree = i.Link
	}
	i.F = nil
	i.P = nil
	ionext = i
}

func newfile(s string, f *os.File) {
	var i *Io

	i = ionext
	i.Link = iostack
	iostack = i
	i.F = f
	if f == nil {
		var err error
		i.F, err = os.Open(s)
		if err != nil {
			Yyerror("%ca: %v", Thechar, err)
			errorexit()
		}
	}

	fi.P = nil
	obj.Linklinehist(Ctxt, int(Lineno), s, 0)
}

func dodef(p string) {
	Dlist = append(Dlist, p)
}

func usage() {
	fmt.Printf("usage: %ca [options] file.c...\n", Thechar)
	flag.PrintDefaults()
	errorexit()
}

func Main() {
	var p string

	// Allow GOARCH=Thestring or GOARCH=Thestringsuffix,
	// but not other values.
	p = obj.Getgoarch()

	if !strings.HasPrefix(p, Thestring) {
		log.Fatalf("cannot use %cc with GOARCH=%s", Thechar, p)
	}
	if p != Thestring {
		Thelinkarch = Arches[p]
		if Thelinkarch == nil {
			log.Fatalf("unknown arch %s", p)
		}
	}

	Ctxt = obj.Linknew(Thelinkarch)
	Ctxt.Diag = Yyerror
	Ctxt.Bso = &bstdout
	Ctxt.Enforce_data_order = 1
	bstdout = *obj.Binitw(os.Stdout)

	debug = [256]int{}
	cinit()
	outfile = ""
	setinclude(".")

	flag.Var(flagFn(dodef), "D", "name[=value]: add #define")
	flag.Var(flagFn(setinclude), "I", "dir: add dir to include path")
	flag.Var((*count)(&debug['S']), "S", "print assembly and machine code")
	flag.Var((*count)(&debug['m']), "m", "debug preprocessor macros")
	flag.StringVar(&outfile, "o", "", "file: set output file")
	flag.StringVar(&Ctxt.Trimpath, "trimpath", "", "prefix: remove prefix from recorded source file paths")

	flag.Parse()

	Ctxt.Debugasm = int32(debug['S'])

	if flag.NArg() < 1 {
		usage()
	}
	if flag.NArg() > 1 {
		fmt.Printf("can't assemble multiple files\n")
		errorexit()
	}

	if assemble(flag.Arg(0)) != 0 {
		errorexit()
	}
	obj.Bflush(&bstdout)
	if nerrors > 0 {
		errorexit()
	}
}

func assemble(file string) int {
	var i int

	if outfile == "" {
		outfile = strings.TrimSuffix(filepath.Base(file), ".s") + "." + string(Thechar)
	}

	of, err := os.Create(outfile)
	if err != nil {
		Yyerror("%ca: cannot create %s", Thechar, outfile)
		errorexit()
	}

	obuf = *obj.Binitw(of)
	fmt.Fprintf(&obuf, "go object %s %s %s\n", obj.Getgoos(), obj.Getgoarch(), obj.Getgoversion())
	fmt.Fprintf(&obuf, "!\n")

	for Pass = 1; Pass <= 2; Pass++ {
		pinit(file)
		for i = 0; i < len(Dlist); i++ {
			dodefine(Dlist[i])
		}
		Yyparse()
		Cclean()
		if nerrors != 0 {
			return nerrors
		}
	}

	obj.Writeobjdirect(Ctxt, &obuf)
	obj.Bflush(&obuf)
	return 0
}

func cinit() {
	for i := 0; i < len(Lexinit); i++ {
		s := Lookup(Lexinit[i].Name)
		if s.Type != LNAME {
			Yyerror("double initialization %s", Lexinit[i].Name)
		}
		s.Type = Lexinit[i].Type
		s.Value = Lexinit[i].Value
	}
}

func syminit(s *Sym) {
	s.Type = LNAME
	s.Value = 0
}

type flagFn func(string)

func (flagFn) String() string {
	return "<arg>"
}

func (f flagFn) Set(s string) error {
	f(s)
	return nil
}

type yyImpl struct{}

// count is a flag.Value that is like a flag.Bool and a flag.Int.
// If used as -name, it increments the count, but -name=x sets the count.
// Used for verbose flag -v.
type count int

func (c *count) String() string {
	return fmt.Sprint(int(*c))
}

func (c *count) Set(s string) error {
	switch s {
	case "true":
		*c++
	case "false":
		*c = 0
	default:
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid count %q", s)
		}
		*c = count(n)
	}
	return nil
}

func (c *count) IsBoolFlag() bool {
	return true
}



/*
 * common code for all the assemblers
 */
func pragpack() {

	for getnsc() != '\n' {

	}
}

func pragvararg() {
	for getnsc() != '\n' {

	}
}

func pragcgo(name string) {
	for getnsc() != '\n' {

	}
}

func pragfpround() {
	for getnsc() != '\n' {

	}
}

func pragtextflag() {
	for getnsc() != '\n' {

	}
}

func pragdataflag() {
	for getnsc() != '\n' {

	}
}

func pragprofile() {
	for getnsc() != '\n' {

	}
}

func pragincomplete() {
	for getnsc() != '\n' {

	}
}






var thetext *obj.LSym

func Settext(s *obj.LSym) {
	thetext = s
}


func isalnum(c int) bool {
	return isalpha(c) || isdigit(c)
}

func isalpha(c int) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func isspace(c int) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func ISALPHA(c int) bool {
	if isalpha(c) {
		return true
	}
	if c >= utf8.RuneSelf {
		return true
	}
	return false
}

var yybuf bytes.Buffer

func (yyImpl) Error(s string) {
	Yyerror("%s", s)
}

type Yylval struct {
	Sym  *Sym
	Lval int64
	Sval string
	Dval float64
}

func Yylex(yylval *Yylval) int {
	var c int
	var c1 int
	var s *Sym

	c = peekc
	if c != IGN {
		peekc = IGN
		goto l1
	}

l0:
	c = GETC()

l1:
	if c == EOF {
		peekc = EOF
		return -1
	}

	if isspace(c) {
		if c == '\n' {
			Lineno++
			return ';'
		}

		goto l0
	}

	if ISALPHA(c) {
		yybuf.Reset()
		goto aloop
	}
	if isdigit(c) {
		goto tnum
	}
	switch c {
	case '\n':
		Lineno++
		return ';'

	case '#':
		domacro()
		goto l0

	case '.':
		c = GETC()
		if ISALPHA(c) {
			yybuf.Reset()
			yybuf.WriteByte('.')
			goto aloop
		}

		if isdigit(c) {
			yybuf.Reset()
			yybuf.WriteByte('.')
			goto casedot
		}

		peekc = c
		return '.'

	case '_',
		'@':
		yybuf.Reset()
		goto aloop

	case '"':
		var buf bytes.Buffer
		c1 = 0
		for {
			c = escchar('"')
			if c == EOF {
				break
			}
			buf.WriteByte(byte(c))
		}
		yylval.Sval = buf.String()
		return LSCONST

	case '\'':
		c = escchar('\'')
		if c == EOF {
			c = '\''
		}
		if escchar('\'') != EOF {
			Yyerror("missing '")
		}
		yylval.Lval = int64(c)
		return LCONST

	case '/':
		c1 = GETC()
		if c1 == '/' {
			for {
				c = GETC()
				if c == '\n' {
					goto l1
				}
				if c == EOF {
					Yyerror("eof in comment")
					errorexit()
				}
			}
		}

		if c1 == '*' {
			for {
				c = GETC()
				for c == '*' {
					c = GETC()
					if c == '/' {
						goto l0
					}
				}

				if c == EOF {
					Yyerror("eof in comment")
					errorexit()
				}

				if c == '\n' {
					Lineno++
				}
			}
		}

	default:
		return int(c)
	}

	peekc = c1
	return int(c)

casedot:
	for {
		yybuf.WriteByte(byte(c))
		c = GETC()
		if !(isdigit(c)) {
			break
		}
	}

	if c == 'e' || c == 'E' {
		goto casee
	}
	goto caseout

casee:
	yybuf.WriteByte('e')
	c = GETC()
	if c == '+' || c == '-' {
		yybuf.WriteByte(byte(c))
		c = GETC()
	}

	for isdigit(c) {
		yybuf.WriteByte(byte(c))
		c = GETC()
	}

caseout:
	peekc = c
	if FPCHIP != 0 /*TypeKind(100016)*/ {
		last = yybuf.String()
		yylval.Dval = atof(last)
		return LFCONST
	}

	Yyerror("assembler cannot interpret fp constants")
	yylval.Lval = 1
	return LCONST

aloop:
	yybuf.WriteByte(byte(c))
	c = GETC()
	if ISALPHA(c) || isdigit(c) || c == '_' || c == '$' {
		goto aloop
	}
	peekc = c
	last = yybuf.String()
	s = Lookup(last)
	if s.Macro != nil {
		newio()
		ionext.P = macexpand(s)
		pushio()
		ionext.Link = iostack
		iostack = ionext
		fi.P = ionext.P
		if peekc != IGN {
			fi.P = append(fi.P, byte(peekc))
			peekc = IGN
		}

		goto l0
	}

	if s.Type == 0 {
		s.Type = LNAME
	}
	if s.Type == LNAME || s.Type == LVAR || s.Type == LLAB {
		yylval.Sym = s
		yylval.Sval = last
		return int(s.Type)
	}

	yylval.Lval = s.Value
	yylval.Sval = last
	return int(s.Type)

tnum:
	yybuf.Reset()
	if c != '0' {
		goto dc
	}
	yybuf.WriteByte(byte(c))
	c = GETC()
	c1 = 3
	if c == 'x' || c == 'X' {
		c1 = 4
		c = GETC()
	} else if c < '0' || c > '7' {
		goto dc
	}
	yylval.Lval = 0
	for {
		if c >= '0' && c <= '9' {
			if c > '7' && c1 == 3 {
				break
			}
			yylval.Lval = int64(uint64(yylval.Lval) << uint(c1))
			yylval.Lval += int64(c) - '0'
			c = GETC()
			continue
		}

		if c1 == 3 {
			break
		}
		if c >= 'A' && c <= 'F' {
			c += 'a' - 'A'
		}
		if c >= 'a' && c <= 'f' {
			yylval.Lval = int64(uint64(yylval.Lval) << uint(c1))
			yylval.Lval += int64(c) - 'a' + 10
			c = GETC()
			continue
		}

		break
	}

	goto ncu

dc:
	for {
		if !(isdigit(c)) {
			break
		}
		yybuf.WriteByte(byte(c))
		c = GETC()
	}

	if c == '.' {
		goto casedot
	}
	if c == 'e' || c == 'E' {
		goto casee
	}
	last = yybuf.String()
	yylval.Lval = strtoll(last, nil, 10)

ncu:
	for c == 'U' || c == 'u' || c == 'l' || c == 'L' {
		c = GETC()
	}
	peekc = c
	return LCONST
}

func getc() int {
	var c int

	c = peekc
	if c != IGN {
		peekc = IGN
		if c == '\n' {
			Lineno++
		}
		return c
	}

	c = GETC()
	if c == '\n' {
		Lineno++
	}
	if c == EOF {
		Yyerror("End of file")
		errorexit()
	}

	return c
}

func getnsc() int {
	var c int

	for {
		c = getc()
		if !isspace(c) || c == '\n' {
			return c
		}
	}
}

func unget(c int) {
	peekc = c
	if c == '\n' {
		Lineno--
	}
}

func escchar(e int) int {
	var c int
	var l int

loop:
	c = getc()
	if c == '\n' {
		Yyerror("newline in string")
		return EOF
	}

	if c != '\\' {
		if c == e {
			return EOF
		}
		return c
	}

	c = getc()
	if c >= '0' && c <= '7' {
		l = c - '0'
		c = getc()
		if c >= '0' && c <= '7' {
			l = l*8 + c - '0'
			c = getc()
			if c >= '0' && c <= '7' {
				l = l*8 + c - '0'
				return l
			}
		}

		peekc = c
		unget(c)
		return l
	}

	switch c {
	case '\n':
		goto loop
	case 'n':
		return '\n'
	case 't':
		return '\t'
	case 'b':
		return '\b'
	case 'r':
		return '\r'
	case 'f':
		return '\f'
	case 'a':
		return 0x07
	case 'v':
		return 0x0b
	case 'z':
		return 0x00
	}

	return c
}

func pinit(f string) {
	Lineno = 1
	newio()
	newfile(f, nil)
	PC = 0
	peekc = IGN
	sym = 1
	for _, s := range hash {
		s.Macro = nil
	}
}

func filbuf() int {
	var i *Io
	var n int

loop:
	i = iostack
	if i == nil {
		return EOF
	}
	if i.F == nil {
		goto pop
	}
	n, _ = i.F.Read(i.B[:])
	if n == 0 {
		i.F.Close()
		obj.Linklinehist(Ctxt, int(Lineno), "<pop>", 0)
		goto pop
	}
	fi.P = i.B[1:n]
	return int(i.B[0]) & 0xff

pop:
	iostack = i.Link
	i.Link = iofree
	iofree = i
	i = iostack
	if i == nil {
		return EOF
	}
	fi.P = i.P
	if len(fi.P) == 0 {
		goto loop
	}
	tmp8 := fi.P
	fi.P = fi.P[1:]
	return int(tmp8[0]) & 0xff
}

var last string


func prfile(l int32) {
	obj.Linkprfile(Ctxt, int(l))
}

func GETC() int {
	var c int
	if len(fi.P) == 0 {
		return filbuf()
	}
	c = int(fi.P[0])
	fi.P = fi.P[1:]
	return c
}

func isdigit(c int) bool {
	return '0' <= c && c <= '9'
}

func strtoll(s string, p *byte, base int) int64 {
	if p != nil {
		panic("strtoll")
	}
	n, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return 0
	}
	return n
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
