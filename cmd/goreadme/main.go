// goreadme creates a README.md from your godoc
//
//
// Usage
//
//    goreadme [-check] [-out path/to/README.md] $PACKAGE
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
	check := flag.Bool("check", false, "check whether readme is up to date")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	pkgName := flag.Arg(0)

	if *check {
		ok, err := goreadme.VerifyReadme(pkgName, *out)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error checking README.md for %s\n", pkgName)
			os.Exit(1)
		}
		if !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s is out of date\n", *out)
			os.Exit(1)
		}
	} else {
		err := goreadme.WriteReadme(pkgName, *out)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed creating README.md for %s\n", pkgName)
			os.Exit(1)
		}
	}
}
