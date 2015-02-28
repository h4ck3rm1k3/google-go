normal:
	bash ./all.bash 

go/build.o: go/build.o go/build/build.go go/build/doc.go go/build/read.go go/build/syslist.go
	gccgo -c -g -o $^

dist.6: cmd/dist/build.go cmd/dist/buildgo.go cmd/dist/buildruntime.go cmd/dist/buildtool.go cmd/dist/main.go cmd/dist/sys_default.go cmd/dist/util.go
	gccgo -g -o dist.6 $^

cmd/internal/obj.o: cmd/internal/obj/ar.go cmd/internal/obj/data.go cmd/internal/obj/flag.go cmd/internal/obj/fmt.go cmd/internal/obj/funcdata.go cmd/internal/obj/go.go cmd/internal/obj/ld.go cmd/internal/obj/libc.go cmd/internal/obj/link.go cmd/internal/obj/mgc0.go cmd/internal/obj/obj.go cmd/internal/obj/objfile.go cmd/internal/obj/pass.go cmd/internal/obj/pcln.go cmd/internal/obj/stack.go cmd/internal/obj/sym.go cmd/internal/obj/textflag.go cmd/internal/obj/typekind.go cmd/internal/obj/util.go
	gccgo -c -g -o $^

DEPS = cmd/internal/obj/x86/6.out.go cmd/internal/obj/x86/anames6.go cmd/internal/obj/x86/asm6.go cmd/internal/obj/x86/list6.go cmd/internal/obj/x86/obj6.go cmd/internal/obj/x86/util.go 

cmd/internal/obj/x86.o : cmd/internal/obj.o  $(DEPS)
	gccgo -L . -I . -c -g -o $(DEPS)


#cmd/internal/obj/data.go:func Symgrow

cmd/6g.o : cmd/6g/cgen.go cmd/6g/galign.go cmd/6g/gg.go cmd/6g/ggen.go cmd/6g/gsubr.go cmd/6g/peep.go cmd/6g/prog.go cmd/6g/reg.go cmd/6g/util.go
	gccgo -L . -I .  -g -o cmd/6g.o $(DEPS) cmd/internal/obj.o


all :
	GOROOT=/home/mdupon/experiments/go PYTHONPATH=../../scons/src/engine/ ../../scons/build/scons/build/scripts/scons 
