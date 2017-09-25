package gopkgs

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/haya14busa/gopkgs/x/tools/imports"
)

// Pkg represents exported go packages.
// It's based on x/tools/imports.Pkg.
type Pkg struct {
	Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
	ImportPath      string // full Pkg import path ("net/http", "foo/bar/vendor/a/b")
	ImportPathShort string // vendorless import path ("net/http", "a/b")

	// It can be empty. It's filled only when Option.IncludeName is true.
	Name string // package name ("http")
}

type Option struct {
	// Fill package name in Pkg struct if it's true. Note that it needs to parse
	// package directory to get package name and it takes some costs.
	IncludeName bool
}

func DefaultOption() *Option {
	return &Option{
		IncludeName: false,
	}
}

// Packages returns all importable Go packages.
// Packages uses [golang.org/x/tools/cmd/goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) implementation internally, so it's fast.
// e.g. it cares .goimportsignore
func Packages(opt *Option) []*Pkg {
	gp := imports.GoPath()
	pkgs := make([]*Pkg, len(gp))
	i := 0
	for _, p := range gp {
		pkgs[i] = copyPkg(p)
		if opt.IncludeName {
			name, _ := packageName(p.Dir)
			if name != "" {
				pkgs[i].Name = name
			}
		}
		i++
	}
	return pkgs
}

func copyPkg(p *imports.Pkg) *Pkg {
	return &Pkg{
		Dir:             p.Dir,
		ImportPath:      p.ImportPath,
		ImportPathShort: p.ImportPathShort,
	}
}

func packageName(dir string) (string, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, notGoTestFile, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	for name := range pkgs {
		return name, nil
	}
	return "", fmt.Errorf("package not found in %s", dir)
}

func notGoTestFile(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(), "_test.go")
}
