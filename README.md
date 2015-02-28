# The GCCGO port of the Go language
This is an attempt to be able to bootstrap go with gccgo and not have to use
a running go program. 

See go/src/SConstruct for the start of a scons
see my patches for the scons go tools here : https://github.com/h4ck3rm1k3/scons-go-tools for the
scons go tools.

The dist tool builds, and I am working on the build command now.

Mike

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](doc/gopher/fiveyears.jpg)

For documentation about how to install and use Go,
visit https://golang.org/ or load doc/install-source.html
in your web browser.

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Please report issues here: https://golang.org/issue/new

Go is the work of hundreds of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

##### Please note that we do not use pull requests.

Unless otherwise noted, the Go source files are distributed
under the BSD-style license found in the LICENSE file.

--

## Binary Distribution Notes

If you have just untarred a binary Go distribution, you need to set
the environment variable $GOROOT to the full path of the go
directory (the one containing this file).  You can omit the
variable if you unpack it into /usr/local/go, or if you rebuild
from sources by running all.bash (see doc/install-source.html).
You should also add the Go binary directory $GOROOT/bin
to your shell's path.

For example, if you extracted the tar file into $HOME/go, you might
put the following in your .profile:

	export GOROOT=$HOME/go
	export PATH=$PATH:$GOROOT/bin

See https://golang.org/doc/install or doc/install.html for more details.
