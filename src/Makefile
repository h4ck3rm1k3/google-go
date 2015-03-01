GCCGO=gccgo-5

clean :
	find -name *.o -exec rm {} \;

normal:
	bash ./all.bash 

go/build.o: go/build.o go/build/build.go go/build/doc.go go/build/read.go go/build/syslist.go
	$(GCCGO) -c -g -o $^

dist.6: cmd/dist/build.go cmd/dist/buildgo.go cmd/dist/buildruntime.go cmd/dist/buildtool.go cmd/dist/main.go cmd/dist/sys_default.go cmd/dist/util.go
	$(GCCGO) -g -o dist.6 $^

cmd/internal/obj.o: cmd/internal/obj/ar.go cmd/internal/obj/data.go cmd/internal/obj/flag.go cmd/internal/obj/fmt.go cmd/internal/obj/funcdata.go cmd/internal/obj/go.go cmd/internal/obj/ld.go cmd/internal/obj/libc.go cmd/internal/obj/link.go cmd/internal/obj/mgc0.go cmd/internal/obj/obj.go cmd/internal/obj/objfile.go cmd/internal/obj/pass.go cmd/internal/obj/pcln.go cmd/internal/obj/stack.go cmd/internal/obj/sym.go cmd/internal/obj/textflag.go cmd/internal/obj/typekind.go cmd/internal/obj/util.go
	$(GCCGO) -c -g -o $@ $^

DEPS = cmd/internal/obj/x86/6.out.go cmd/internal/obj/x86/anames6.go cmd/internal/obj/x86/asm6.go cmd/internal/obj/x86/list6.go cmd/internal/obj/x86/obj6.go cmd/internal/obj/x86/util.go 

#cmd/internal/obj/x86
cmd/internal/obj/x86.o : cmd/internal/obj.o  $(DEPS)
	$(GCCGO) -L . -I . -c -g -o $@ $(DEPS)


#cmd/internal/obj/data.go:func Symgrow

cgen6 : cmd/6g/cgen.go 
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

#galign6: cmd/6g/galign.go
#	$(GCCGO) -c -g -o $@ $^

galign6 : cmd/6g/galign.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

gg6:  cmd/6g/gg.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

ggen6:  cmd/6g/ggen.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

gsubr6: cmd/6g/gsubr.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

peep6: cmd/6g/peep.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

prog6: cmd/6g/prog.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

reg6: cmd/6g/reg.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o

util6: cmd/6g/util.go
	$(GCCGO) -L . -I .  -g -o $@ $^ cmd/internal/obj.o


# the boostrap contains :
asm6 : ./cmd/asm/main.go
	$(GCCGO) -c -g -o $@ $^

#	 ./cmd/asm/internal/arch/arm.go \
#	 ./cmd/asm/internal/arch/ppc64.go

cmd/asm/internal/arch.o : ./cmd/asm/internal/arch/arch.go 
	$(GCCGO) -c -g -o $@ $^

cmd/asm/internal/flags.o: ./cmd/asm/internal/flags/flags.go
	$(GCCGO) -c -g -o $@ $^

# main
LEX_DEPS=cmd/asm/internal/flags.o  \
	cmd/internal/obj.o      \
	text/scanner.o          \
	cmd/asm/internal/lex.o  \
	cmd/internal/obj/x86.o  \
	cmd/asm/internal/arch.o \
	cmd/internal/asm.o      



lex6: cmd/6a/lex.go
	$(GCCGO) -L. -I. -g -o $@ $^ $(LEX_DEPS)


# cmd/asm/internal/flags.o text/scanner.o  cmd/internal/obj.o

lex : $(LEX_DEPS) lex6
	echo

text/scanner.o : text/scanner/scanner.go
	$(GCCGO) -c -g -o $@ $^  

cmd/asm/internal/lex.o: ./cmd/asm/internal/lex/input.go \
	 ./cmd/asm/internal/lex/lex.go \
	 ./cmd/asm/internal/lex/slice.go \
	 ./cmd/asm/internal/lex/stack.go \
	 ./cmd/asm/internal/lex/tokenizer.go
	$(GCCGO) -L. -I. -c -g -o $@ $^ text/scanner.o

cmd/internal/asm.o: cmd/asm/internal/asm/asm.go \
		    cmd/asm/internal/asm/parse.go
	$(GCCGO) -c -g -o $@ $^

cmd/internal/gc.o:  ./cmd/internal/gc/align.go \
	 ./cmd/internal/gc/array.go \
	 ./cmd/internal/gc/bits.go \
	 ./cmd/internal/gc/builtin.go \
	 ./cmd/internal/gc/builtins/runtime.go \
	 ./cmd/internal/gc/builtins/unsafe.go \
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
	 ./cmd/internal/gc/mkbuiltin.go \
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


bootstrap: lex galign6 cgen6 galign6 gg6 ggen6 gsubr6 peep6 prog6 reg6 util6 asm6



# these are the main routins
# cmd/5a/lex.go
# cmd/5g/galign.go
# cmd/6a/lex.go
# cmd/6g/galign.go
# cmd/8a/lex.go
# cmd/8g/galign.go
# cmd/9a/lex.go
# cmd/9g/galign.go
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
# cmd/yacc/yacc.go
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
