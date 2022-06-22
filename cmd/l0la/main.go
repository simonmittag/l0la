package main

import (
	"flag"
	"fmt"
	"github.com/simonmittag/l0la"
	"os"
	"strconv"
)

type Mode uint8

const (
	Watch Mode = 1 << iota
	Version
)

var pattern = "/mse6/"

func main() {
	mode := Watch
	vM := flag.Bool("v", false, "print version")
	flag.Parse()

	if *vM {
		mode = Version
	}

	switch mode {
	case Watch:
		l0la.Watch(parsePid())
	case Version:
		printVersion()
	}
}

func parsePid() uint {
	if len(flag.Args()) != 1 {
		usage()
	}
	pid, _ := strconv.Atoi(flag.Args()[0])
	return uint(pid)
}

func usage() {
	fmt.Fprintf(os.Stdout, "Usage: l0la [-v] pid\n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("l0la %s\n", l0la.Version)
	os.Exit(0)
}
