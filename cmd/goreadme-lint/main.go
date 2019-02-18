package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/WillAbides/godoc2md/goreadme"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, " ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	excludes := arrayFlags([]string{})
	flag.Var(&excludes, "exclude", "directory to be excluded (may be used more than once)")
	readmeName := flag.String("name", "README.md", "filename for readmes")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	baseDir := flag.Arg(0)
	ok, outdated, err := goreadme.CheckReadmes(baseDir, *readmeName, excludes)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed checking readmes")
		os.Exit(1)
	}
	if !ok {
		for _, o := range outdated {
			fmt.Printf("%s is outdated\n", o)
		}
		os.Exit(1)
	}
}
