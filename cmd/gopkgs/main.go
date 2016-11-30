package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/haya14busa/gopkgs/x/tools/imports"
)

var (
	fullpath = flag.Bool("fullpath", false, `output absolute file path to package directory. ("/usr/lib/go/src/net/http")`)
	short    = flag.Bool("short", false, `output vendorless import path ("net/http", "foo/bar/vendor/a/b")`)
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, pkg := range imports.GoPath() {
		out := pkg.ImportPath
		if *fullpath {
			out = pkg.Dir
		} else if *short {
			out = pkg.ImportPathShort
		}
		fmt.Fprintln(w, out)
	}
}
