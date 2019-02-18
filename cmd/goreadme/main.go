// goreadme creates a README.md from your godoc
//
//
// Usage
//
//    goreadme [-out path/to/README.md] $PACKAGE
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/WillAbides/godoc2md/goreadme"
)

//go:generate ../../bin/goreadme github.com/WillAbides/godoc2md/cmd/goreadme

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: goreadme [options] package\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	out := flag.String("out", filepath.FromSlash("./README.md"), "path to README.md")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	pkgName := flag.Arg(0)

	err := goreadme.WriteReadme(pkgName, *out)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed creating README.md for %s", pkgName)
		os.Exit(1)
	}
}
