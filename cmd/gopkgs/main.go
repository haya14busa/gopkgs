package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"os"
	"text/tabwriter"

	"github.com/haya14busa/gopkgs/x/tools/imports"
)

var (
	format = flag.String("f", "{{.ImportPathShort}}", "output format of the package")
)

var usageInfo = `
Use -f to custom the output using template syntax. The struct being passed to template is:
	type Pkg struct {
		Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
		Name            string // package name ("http", "a")
		ImportPath      string // full Pkg import path ("net/http", "foo/bar/vendor/a/b")
		ImportPathShort string // vendorless import path ("net/http", "a/b")
	}
`

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	tw := tabwriter.NewWriter(os.Stderr, 0, 0, 4, ' ', tabwriter.AlignRight)
	fmt.Fprintln(tw, usageInfo)
}

func init() {
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		flag.Usage()
		os.Exit(2)
	}

	tpl, err := template.New("out").Parse(*format)
	if err != nil {
		fmt.Fprintln(os.Stderr)
		os.Exit(2)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, pkg := range imports.GoPath() {
		tpl.Execute(w, pkg)
		fmt.Fprintln(w)
	}
}
