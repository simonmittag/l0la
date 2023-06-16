package main

import (
	"flag"
	"os"
	"testing"
)

func TestMainFuncWithVersion(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append([]string{""}, "-v")
	main()
}

func TestMainFuncWithUsage(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append([]string{""}, "-h")
	main()
}

//infinite loop
//func TestMainFuncWithPositionalArgsBeforeFlagArgs(t *testing.T) {
//	//first arg is command
//	os.Args = []string{"l0la", fmt.Sprintf("%v", os.Getpid())}
//	main()
//}

func TestPrintVersion(t *testing.T) {
	printVersion()
}

func TestPrintUsage(t *testing.T) {
	printUsage()
}
