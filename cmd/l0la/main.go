package main

import (
	"errors"
	"flag"
	"fmt"
	. "github.com/logrusorgru/aurora"
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

var pattern = "/mse6/"

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
			printVersion()
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

func printUsage() {
	printVersion()
	fmt.Fprintln(os.Stdout, "Usage: l0la [-v] [pid]")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("\n"+Red("<").String()+"l"+Blue("0").String()+"la"+Red(">").String()+" %s\n", l0la.Version)
}

func printError(err error) {
	fmt.Printf("unable to watch: %v", err)
}
