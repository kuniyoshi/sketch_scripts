package main

import (
	"flag"
	"fmt"
	"os"
)

// These can be overridden at build time via -ldflags
var (
	version = "0.1.0"
	commit  = ""
	date    = ""
)

func printVersion() {
	if commit == "" && date == "" {
		fmt.Printf("sketch_scripts %s\n", version)
		return
	}
	fmt.Printf("sketch_scripts %s (%s %s)\n", version, commit, date)
}

func usage() {
	fmt.Fprintf(os.Stderr, `sketch_scripts — a tiny Go CLI

Usage:
  sketch_scripts [flags] <command> [args]

Flags:
  -h, --help       Show help
  -v, --version    Show version

Commands:
  hello            Print a friendly greeting
  help             Show this help

Examples:
  sketch_scripts hello -name Alice

`)
}

func main() {
	// Global flags
	showVersion := flag.Bool("version", false, "Show version")
	// Support -v as an alias for --version
	vAlias := flag.Bool("v", false, "Show version (alias)")

	flag.Usage = usage
	flag.Parse()

	if *showVersion || *vAlias {
		printVersion()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		usage()
		return
	}

	switch args[0] {
	case "help", "--help", "-h":
		usage()
	case "hello":
		helloCmd := flag.NewFlagSet("hello", flag.ExitOnError)
		name := helloCmd.String("name", "world", "Name to greet")
		_ = helloCmd.Parse(args[1:])
		fmt.Printf("Hello, %s!\n", *name)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", args[0])
		usage()
		os.Exit(2)
	}
}
