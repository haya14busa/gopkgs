package gopkgs

import "testing"

func TestPackages(t *testing.T) {
	pkgs := Packages(&Option{IncludeName: false})
	if len(pkgs) == 0 {
		t.Fatal("len(Packages()) == 0")
	}
	if pkgs[0].Name != "" {
		t.Errorf("Package.Name is not empty when IncludeName==false")
	}
}

func TestPackages_IncludeName(t *testing.T) {
	pkgs := Packages(&Option{IncludeName: true})
	if len(pkgs) == 0 {
		t.Fatal("len(Packages()) == 0")
	}
	if pkgs[0].Name == "" {
		t.Errorf("Package.Name is empty when IncludeName==true")
	}
}
