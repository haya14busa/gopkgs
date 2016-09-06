package imports

import (
	"testing"

	"regexp"
)

var githubReg = regexp.MustCompile("^github.com/")

func TestGoPath(t *testing.T) {
	ds := GoPath()
	GoPath() // confirm GoPath() can be called more than once
	foundStdLib := false
	foundGOPATH := false
	for _, pkg := range ds {
		if pkg.ImportPath == "net/http" {
			foundStdLib = true
		}
		if githubReg.MatchString(pkg.ImportPath) {
			foundGOPATH = true
		}
		if foundStdLib && foundGOPATH {
			break
		}
	}
	if !foundStdLib {
		t.Error("Standard Library net/http not found")
	}
	if !foundGOPATH {
		t.Error("GOPAH library which starts with github.com/ not found")
	}
}
