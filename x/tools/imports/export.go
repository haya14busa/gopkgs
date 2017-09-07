package imports

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"sync"
)

// export.go exports some type and func of golang.org/x/tools/imports

var (
	exportedGoPath   map[string]*Pkg
	exportedGoPathMu sync.RWMutex
)

// GoPath returns all importable packages (abs dir path => *Pkg).
func GoPath() map[string]*Pkg {
	exportedGoPathMu.Lock()
	defer exportedGoPathMu.Unlock()
	if exportedGoPath != nil {
		return exportedGoPath
	}
	populateIgnoreOnce.Do(populateIgnore)
	scanGoRootOnce.Do(scanGoRoot) // async
	scanGoPathOnce.Do(scanGoPath)
	<-scanGoRootDone
	dirScanMu.Lock()
	defer dirScanMu.Unlock()
	exportedGoPath = exportDirScan(dirScan)
	return exportedGoPath
}

func exportDirScan(ds map[string]*pkg) map[string]*Pkg {
	r := make(map[string]*Pkg)
	for path, pkg := range ds {
		p, err := exportPkg(pkg)
		if err != nil {
			continue
		}

		r[path] = p
	}
	return r
}

// Pkg represents exported type of pkg.
type Pkg struct {
	Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
	Name            string // package name ("http")
	ImportPath      string // full Pkg import path ("net/http", "foo/bar/vendor/a/b")
	ImportPathShort string // vendorless import path ("net/http", "a/b")
}

func exportPkg(p *pkg) (*Pkg, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, p.dir, notGoTestFile, parser.PackageClauseOnly)
	if err != nil {
		return nil, err
	}

	for name := range pkgs {
		// one package is valid, the test package ignored
		return &Pkg{Dir: p.dir, Name: name, ImportPath: p.importPath, ImportPathShort: p.importPathShort}, nil
	}

	return nil, fmt.Errorf("no exported package found on %s", p.dir)
}

func notGoTestFile(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(), "_test.go")
}
