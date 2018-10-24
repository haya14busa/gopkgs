package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/haya14busa/gopkgs"
)

var (
	fullpath    = flag.Bool("fullpath", false, `output absolute file path to package directory. ("/usr/lib/go/src/net/http")`)
	short       = flag.Bool("short", true, `output vendorless import path ("net/http", "foo/bar/vendor/a/b")`)
	f           = flag.String("f", "", "alternate format for the output using the syntax of template package. e.g. {{.Name}};{{ImportPathShort}}")
	includeName = flag.Bool("include-name", false, "fill Pkg.Name which can be used with -f flag")
)

var usageInfo = `
Use -f to custom the output using template syntax. The struct being passed to template is:
	type Pkg struct {
		Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
		ImportPath      string // full Pkg import path ("net/http", "foo/bar/vendor/a/b")
		ImportPathShort string // vendorless import path ("net/http", "a/b")

		// It can be empty. It's filled only when -include-name flag is true.
		Name string // package name ("http")
	}
`

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, usageInfo)
}

func init() {
	flag.Usage = usage
}

type ByPath []*gopkgs.Pkg

func (pkgs ByPath) Len() int {
	return len(pkgs)
}
func (pkgs ByPath) Swap(i, j int) {
	pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
}
func (pkgs ByPath) Less(i, j int) bool {
	return strings.Compare(pkgs[i].ImportPath, pkgs[j].ImportPath) < 0
}

type ByShortPath []*gopkgs.Pkg

func (pkgs ByShortPath) Len() int {
	return len(pkgs)
}
func (pkgs ByShortPath) Swap(i, j int) {
	pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
}
func (pkgs ByShortPath) Less(i, j int) bool {
	return strings.Compare(pkgs[i].ImportPathShort, pkgs[j].ImportPathShort) < 0
}

type ByFullPath []*gopkgs.Pkg

func (pkgs ByFullPath) Len() int {
	return len(pkgs)
}
func (pkgs ByFullPath) Swap(i, j int) {
	pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
}
func (pkgs ByFullPath) Less(i, j int) bool {
	return strings.Compare(pkgs[i].Dir, pkgs[j].Dir) < 0
}

func uniq(s []*gopkgs.Pkg, f func(*gopkgs.Pkg) string) []*gopkgs.Pkg {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	u := make([]*gopkgs.Pkg, 0, l)
	prev := f(s[0])
	for _, p := range s[1:] {
		cur := f(p)
		if prev == cur {
			continue
		}
		u = append(u, p)
		prev = cur
	}

	return u
}

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		flag.Usage()
		os.Exit(2)
	}

	tplFormat := "{{.ImportPath}}"
	if *f != "" {
		tplFormat = *f
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

	opt := gopkgs.DefaultOption()
	opt.IncludeName = *includeName
	pkgs := gopkgs.Packages(opt)

	if *fullpath {
		sort.Sort(ByFullPath(pkgs))
		// Fullpaths is already unique
	} else if *short {
		sort.Sort(ByShortPath(pkgs))
		pkgs = uniq(pkgs, func(p *gopkgs.Pkg) string { return p.ImportPathShort })
	} else {
		sort.Sort(ByPath(pkgs))
		pkgs = uniq(pkgs, func(p *gopkgs.Pkg) string { return p.ImportPath })
	}

	for _, pkg := range pkgs {
		tpl.Execute(w, pkg)
		fmt.Fprintln(w)
	}
}
