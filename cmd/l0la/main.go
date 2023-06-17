package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/simonmittag/l0la"
	"os"
	"strconv"
)

type Mode uint8

const (
	Watch Mode = 1 << iota
	Usage
	Version
)

func main() {
	mode := Watch
	vM := flag.Bool("v", false, "print version")
	vU := flag.Bool("h", false, "print help")
	flag.Parse()

	if *vM {
		mode = Version
	} else if *vU {
		mode = Usage
	}

	pid, err := parsePid()

	switch mode {
	case Watch:
		if err == nil {
			l0la.Watch(pid)
		} else {
			printUsage()
		}
	case Usage:
		printUsage()
	case Version:
		printVersion()
	}
}

func parsePid() (int, error) {
	if len(flag.Args()) != 1 {
		return 0, errors.New("pid not supplied")
	}
	pid, err := strconv.Atoi(flag.Args()[0])
	return pid, err
}

func printVersion() {
	fmt.Println(l0la.Logo())
}

func printUsage() {
	fmt.Println(l0la.Logo())
	fmt.Fprintln(os.Stdout, "Usage: l0la [-v] [pid]")
	flag.PrintDefaults()
}

func printError(err error) {
	fmt.Printf("unable to watch: %v", err)
}
