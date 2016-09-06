package gopkgs

import "github.com/haya14busa/gopkgs/x/tools/imports"

// Packages returns all importable Go packages.
// Packages uses [golang.org/x/tools/cmd/goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) implementation internally, so it's fast.
// e.g. it cares .goimportsignore
func Packages() []*imports.Pkg {
	gp := imports.GoPath()
	pkgs := make([]*imports.Pkg, 0, len(gp))
	return pkgs
}
