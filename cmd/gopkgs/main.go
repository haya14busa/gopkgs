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
	fullpath = flag.Bool("fullpath", false, `output absolute file path to package directory. ("/usr/lib/go/src/net/http")`)
	short    = flag.Bool("short", true, `output vendorless import path ("net/http", "foo/bar/vendor/a/b")`)
	format   = flag.String("format", "", "custom output format")
)

var usageInfo = `
Use -format to custom the output using template syntax. The struct being passed to template is:
	type Pkg struct {
		Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
		Name            string // package name ("http")
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

	tplFormat := "{{.ImportPath}}"
	if *format != "" {
		tplFormat = *format
	} else if *fullpath {
		tplFormat = "{{.Dir}}"
	} else if *short {
		tplFormat = "{{.ImportPathShort}}"
	}

	tpl, err := template.New("out").Parse(tplFormat)
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
