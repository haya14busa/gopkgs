## gopkgs - List Go packages FAST by using the same implementation as [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

[![Travis Build Status](https://travis-ci.org/haya14busa/gopkgs.svg?branch=master)](https://travis-ci.org/haya14busa/gopkgs)
[![Appveyor Build status](https://ci.appveyor.com/api/projects/status/9tr7p8hclfypvwun?svg=true)](https://ci.appveyor.com/project/haya14busa/gopkgs)
[![Releases](https://img.shields.io/github/tag/haya14busa/gopkgs.svg)](https://github.com/haya14busa/gopkgs/releases)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/haya14busa/gopkgs?status.svg)](https://godoc.org/github.com/haya14busa/gopkgs)

gopkgs outputs list of importable Go packages.

By using the same implementation as [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports),
it's faster than [go list ...](https://golang.org/cmd/go/#hdr-List_packages) and it also has `-f` option.

gopkgs cares .goimportsignore which was introduced by https://github.com/golang/go/issues/16386
since it uses the same implementation as [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports).

![gopkgs_usage.gif (890×542)](https://raw.githubusercontent.com/haya14busa/i/3bdf4c81118c4f0261073a7f3144903623240edc/gopkgs/gopkgs_usage.gif)
Sample usage of gopkgs with other tools like [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) and filtering tools ([peco](https://github.com/peco/peco)).

### Installation

#### Install Binary from GitHub Releases

https://github.com/haya14busa/gopkgs/releases

#### go get

```
go get -u github.com/haya14busa/gopkgs/cmd/gopkgs
```

### SYNOPSIS

```
$ gopkgs -h
Usage of gopkgs:
  -format string
        custom output format
  -fullpath
        output absolute file path to package directory. ("/usr/lib/go/src/net/http")
  -include-name
        fill Pkg.Name which can be used with -format flag
  -short
        output vendorless import path ("net/http", "foo/bar/vendor/a/b") (default true)


Use -format to custom the output using template syntax. The struct being passed to template is:
    type Pkg struct {
        Dir             string // absolute file path to Pkg directory ("/usr/lib/go/src/net/http")
        ImportPath      string // full Pkg import path ("net/http", "foo/bar/vendor/a/b")
        ImportPathShort string // vendorless import path ("net/http", "a/b")

        // It can be empty. It's filled only when -include-name flag is true.
        Name string // package name ("http")
    }

```

### Vim

![gopkgs_vim.gif (890×542)](https://raw.githubusercontent.com/haya14busa/i/3bdf4c81118c4f0261073a7f3144903623240edc/gopkgs/gopkgs_vim.gif)

Sample usage of gopkgs in Vim: open godoc and import with [vim-go](https://github.com/fatih/vim-go) and [fzf](https://github.com/junegunn/fzf)

```vim
augroup gopkgs
  autocmd!
  autocmd FileType go command! -buffer Import exe 'GoImport' fzf#run({'source': 'gopkgs'})[0]
  autocmd FileType go command! -buffer Doc exe 'GoDoc' fzf#run({'source': 'gopkgs'})[0]
augroup END
```

Above Vim script is just a sample and isn't robust. I'm planning to create or contribute Vim plugins to include the same feature.

- https://github.com/rhysd/unite-go-import.vim
  - [unite.vim](https://github.com/Shougo/unite.vim) (and
    [vim-go](https://github.com/fatih/vim-go)) intergration for importing
    package or opening godoc in Vim.

### LICENSE

[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

#### under [x/tools/imports/](x/tools/imports/) directory

```
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

gopkgs copied and modified https://godoc.org/golang.org/x/tools/imports to export some interface
in accordance with LICENSE.

If The Go Authors provide goimports's directory scanning code as a library, I plan to use it.
ref: https://github.com/golang/go/issues/16427

### :bird: Author
haya14busa (https://github.com/haya14busa)

