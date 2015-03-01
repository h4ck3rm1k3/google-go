GCCGO=gccgo-5

bootstrap: lex6 galign6 yacc6 asm6 dist6

clean :
	find -name *.o -exec rm {} \;

normal:
	bash ./all.bash 

go/build.o: go/build.o go/build/build.go go/build/doc.go go/build/read.go go/build/syslist.go
	$(GCCGO) -c -g -o $^

dist6: cmd/dist/build.go cmd/dist/buildgo.go cmd/dist/buildruntime.go cmd/dist/buildtool.go cmd/dist/main.go cmd/dist/sys_default.go cmd/dist/util.go
	$(GCCGO) -g -o dist.6 $^



cmd/internal/obj.o: cmd/internal/obj/ar.go cmd/internal/obj/data.go cmd/internal/obj/flag.go cmd/internal/obj/fmt.go cmd/internal/obj/funcdata.go cmd/internal/obj/go.go cmd/internal/obj/ld.go cmd/internal/obj/libc.go cmd/internal/obj/link.go cmd/internal/obj/mgc0.go cmd/internal/obj/obj.go cmd/internal/obj/objfile.go cmd/internal/obj/pass.go cmd/internal/obj/pcln.go cmd/internal/obj/stack.go cmd/internal/obj/sym.go cmd/internal/obj/textflag.go cmd/internal/obj/typekind.go cmd/internal/obj/util.go cmd/internal/obj/zbootstrap.go cmd/internal/obj/util_const.go
	$(GCCGO) -c -g -o $@ $^

#cmd/internal/obj/x86
cmd/internal/obj/x86.o : cmd/internal/obj.o  cmd/internal/obj/x86/6.out.go cmd/internal/obj/x86/anames6.go cmd/internal/obj/x86/asm6.go cmd/internal/obj/x86/list6.go cmd/internal/obj/x86/obj6.go cmd/internal/obj/x86/util.go 
	$(GCCGO) -L . -I . -c -g -o $@ $^

cmd/internal/obj/i386.o: cmd/internal/obj/i386/8.out.go  \
	cmd/internal/obj/i386/list8.go \
	cmd/internal/obj/i386/anames8.go \
	cmd/internal/obj/i386/obj8.go \
	cmd/internal/obj/i386/asm8.go  \
	cmd/internal/obj/i386/util.go
	$(GCCGO) -L . -I . -c -g -o $@ $^

cmd/internal/obj/ppc64.o: \
	cmd/internal/obj/ppc64/9.out.go \
	cmd/internal/obj/ppc64/list9.go \
	cmd/internal/obj/ppc64/anames9.go \
	cmd/internal/obj/ppc64/obj9.go \
	cmd/internal/obj/ppc64/asm9.go \
	cmd/internal/obj/ppc64/util.go
	$(GCCGO) -L . -I . -c -g -o $@ $^

cmd/internal/obj/arm.o : cmd/internal/obj.o  \
	cmd/internal/obj/arm/5.out.go \
	cmd/internal/obj/arm/anames5.go \
	cmd/internal/obj/arm/asm5.go \
	cmd/internal/obj/arm/list5.go \
	cmd/internal/obj/arm/obj5.go  \
	cmd/internal/obj/arm/util.go
	$(GCCGO) -L . -I . -c -g -o $@ $^


#cmd/internal/obj/data.go:func Symgrow


galign6 : cmd/6g/galign.go cmd/6g/gsubr.go cmd/6g/cgen.go cmd/6g/gg.go cmd/6g/ggen.go cmd/6g/peep.go cmd/6g/prog.go cmd/6g/reg.go cmd/6g/util.go cmd/internal/obj/x86.o cmd/internal/obj.o cmd/internal/gc.o
	$(GCCGO) -L . -I .  -g -o $@ $^ 

#	cmd/internal/asm.o \
# the boostrap contains :
asm6 :  cmd/asm/internal/arch.o \
	cmd/asm/internal/asm.o \
	cmd/asm/internal/flags.o \
	cmd/asm/internal/lex.o  \
	cmd/internal/gc.o \
	cmd/internal/obj.o \
	cmd/internal/obj/x86.o  \
	cmd/internal/obj/arm.o  \
	cmd/internal/obj/i386.o  \
	cmd/internal/obj/ppc64.o  \
	text/scanner.o \
	cmd/asm/main.go \
	cmd/asm/main.go 
	$(GCCGO) -I. -L. -g -o $@ $^ 

cmd/asm/internal/arch.o : ./cmd/asm/internal/arch/arch.go  \
	./cmd/asm/internal/arch/arm.go  \
	./cmd/asm/internal/arch/ppc64.go  \
	cmd/internal/obj/x86.o  \
	cmd/internal/obj/arm.o  \
	cmd/internal/obj/i386.o  \
	cmd/internal/obj/ppc64.o 
	$(GCCGO) -c -g -o $@ $^

cmd/asm/internal/flags.o: ./cmd/asm/internal/flags/flags.go
	$(GCCGO) -c -g -o $@ $^

#cmd/6a/lex.go  : missing variables
lex6:  cmd/6a/lex2.go \
	cmd/asm/internal/flags.o  \
	cmd/internal/obj.o      \
	text/scanner.o          \
	cmd/asm/internal/lex.o  \
	cmd/internal/obj/x86.o  \
	cmd/internal/obj/ppc64.o  \
	cmd/internal/obj/i386.o  \
	cmd/internal/obj/arm.o  \
	cmd/asm/internal/arch.o \
	cmd/internal/gc.o cmd/internal/asm.o
	$(GCCGO) -L. -I. -g -o $@ $^ 

# this one is broken
lexorig:  cmd/6a/lex.go \
	cmd/asm/internal/flags.o  \
	cmd/internal/obj.o      \
	text/scanner.o          \
	cmd/asm/internal/lex.o  \
	cmd/internal/obj/x86.o  \
	cmd/internal/obj/ppc64.o  \
	cmd/internal/obj/i386.o  \
	cmd/internal/obj/arm.o  \
	cmd/asm/internal/arch.o \
	cmd/internal/gc.o cmd/internal/asm.o
	$(GCCGO) -L. -I. -g -o $@ $^ 

text/scanner.o : text/scanner/scanner.go
	$(GCCGO) -c -g -o $@ $^  

cmd/asm/internal/lex.o: ./cmd/asm/internal/lex/input.go \
	 ./cmd/asm/internal/lex/lex.go \
	 ./cmd/asm/internal/lex/slice.go \
	 ./cmd/asm/internal/lex/stack.go \
	 ./cmd/asm/internal/lex/tokenizer.go \
	text/scanner.o
	$(GCCGO) -L. -I. -c -g -o $@ ./cmd/asm/internal/lex/input.go \
	 ./cmd/asm/internal/lex/lex.go \
	 ./cmd/asm/internal/lex/slice.go \
	 ./cmd/asm/internal/lex/stack.go \
	 ./cmd/asm/internal/lex/tokenizer.go



cmd/asm/internal/asm.o: cmd/asm/internal/lex.o cmd/asm/internal/flags.o cmd/asm/internal/asm/asm.go cmd/asm/internal/asm/parse.go cmd/asm/internal/lex.o cmd/asm/internal/arch.o
	$(GCCGO)  -g -c -o $@ cmd/asm/internal/asm/asm.go cmd/asm/internal/asm/parse.go

cmd/internal/asm.o: cmd/internal/obj.o cmd/internal/asm/asm2.go \
	cmd/internal/asm/asm.go  \
	cmd/internal/asm/lexbody.go \
	cmd/internal/asm/macbody.go
	$(GCCGO) -g -c -o $@ cmd/internal/asm/asm2.go \
	cmd/internal/asm/asm.go  \
	cmd/internal/asm/lexbody.go \
	cmd/internal/asm/macbody.go

cmd/internal/package.o:
	 ./cmd/internal/gc/builtins/runtime.go \
	 ./cmd/internal/gc/builtins/unsafe.go
	$(GCCGO) -c -g -o $@ $^

cmd/internal/gc.o:  ./cmd/internal/gc/align.go \
	 ./cmd/internal/gc/array.go \
	 ./cmd/internal/gc/bits.go \
	 ./cmd/internal/gc/builtin.go \
	 ./cmd/internal/gc/bv.go \
	 ./cmd/internal/gc/closure.go \
	 ./cmd/internal/gc/const.go \
	 ./cmd/internal/gc/cplx.go \
	 ./cmd/internal/gc/dcl.go \
	 ./cmd/internal/gc/esc.go \
	 ./cmd/internal/gc/export.go \
	 ./cmd/internal/gc/fmt.go \
	 ./cmd/internal/gc/gen.go \
	 ./cmd/internal/gc/go.go \
	 ./cmd/internal/gc/gsubr.go \
	 ./cmd/internal/gc/init.go \
	 ./cmd/internal/gc/inl.go \
	 ./cmd/internal/gc/lex.go \
	 ./cmd/internal/gc/md5.go \
	 ./cmd/internal/gc/mparith1.go \
	 ./cmd/internal/gc/mparith2.go \
	 ./cmd/internal/gc/mparith3.go \
	 ./cmd/internal/gc/obj.go \
	 ./cmd/internal/gc/opnames.go \
	 ./cmd/internal/gc/order.go \
	 ./cmd/internal/gc/pgen.go \
	 ./cmd/internal/gc/plive.go \
	 ./cmd/internal/gc/popt.go \
	 ./cmd/internal/gc/racewalk.go \
	 ./cmd/internal/gc/range.go \
	 ./cmd/internal/gc/reflect.go \
	 ./cmd/internal/gc/reg.go \
	 ./cmd/internal/gc/select.go \
	 ./cmd/internal/gc/sinit.go \
	 ./cmd/internal/gc/subr.go \
	 ./cmd/internal/gc/swt.go \
	 ./cmd/internal/gc/typecheck.go \
	 ./cmd/internal/gc/unsafe.go \
	 ./cmd/internal/gc/util.go \
	 ./cmd/internal/gc/walk.go \
	 ./cmd/internal/gc/y.go
	$(GCCGO) -c -g -o $@ $^



cgo : cmd/cgo/main.go \
	cmd/cgo/ast.go \
	cmd/cgo/doc.go \
	cmd/cgo/gcc.go \
	cmd/cgo/godefs.go \
	cmd/cgo/out.go \
	cmd/cgo/util.go 
	$(GCCGO) -c -g -o $@ $^

yacc6: cmd/yacc/yacc.go
	$(GCCGO)  -g -o $@ $^

# these are the main routins
# cmd/addr2line/main.go
# cmd/api/goapi.go
# cmd/api/run.go
# cmd/asm/main.go
# cmd/cgo/doc.go
# cmd/cgo/main.go
# cmd/dist/util.go
# cmd/fix/main.go
# cmd/go/main.go
# cmd/go/test.go
# cmd/gofmt/gofmt.go
# cmd/internal/gc/mkbuiltin.go
# cmd/link/main.go
# cmd/nm/nm.go
# cmd/objdump/main.go
# cmd/pack/pack.go
# cmd/pprof/pprof.go
# cmd/trace/main.go
# compress/flate/gen.go
# crypto/md5/gen.go
# crypto/tls/generate_cert.go
# encoding/gob/decgen.go
# encoding/gob/dump.go
# encoding/gob/encgen.go
# go/doc/headscan.go
# image/color/palette/gen.go
# net/http/triv.go
# runtime/gengoos.go
# runtime/proc.go
# runtime/proc.go
# runtime/wbfat_gen.go
# runtime/wincallback.go
# strconv/makeisprint.go
# syscall/mksyscall_windows.go
# time/genzabbrs.go
# unicode/maketables.go


all :
	GOROOT=/home/mdupon/experiments/go PYTHONPATH=../../scons/src/engine/ ../../scons/build/scons/build/scripts/scons 
