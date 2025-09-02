package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
  list             List scripts and descriptions
  help             Show this help

Examples:
  sketch_scripts hello -name Alice
  sketch_scripts list -dir sketches

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
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		dir := listCmd.String("dir", "sketches", "Directory containing scripts")
		_ = listCmd.Parse(args[1:])
		if err := listSketches(*dir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", args[0])
		usage()
		os.Exit(2)
	}
}

// listSketches scans the given directory for files and prints
// "<filename>\t<description>" where description is taken from the first line
// that contains "SKETCH:". If none is found, it prints "(no description)".
func listSketches(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("reading dir %s: %w", dir, err)
	}

	// Collect file paths and sort by name for stable output
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		files = append(files, filepath.Join(dir, e.Name()))
	}
	sort.Strings(files)

	for _, path := range files {
		desc, err := extractSketchDescription(path)
		if err != nil {
			// Non-fatal: show an error marker for this file
			fmt.Printf("%s\t%s\n", filepath.Base(path), "(read error)")
			continue
		}
		if desc == "" {
			desc = "(no description)"
		}
		fmt.Printf("%s\t%s\n", filepath.Base(path), desc)
	}
	return nil
}

func extractSketchDescription(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	// Look for a line containing "SKETCH:" and return the rest of that line
	const key = "SKETCH:"
	for _, line := range strings.Split(string(data), "\n") {
		if idx := strings.Index(line, key); idx != -1 {
			// Take everything after the key
			val := line[idx+len(key):]
			// Trim common comment prefixes and whitespace
			val = strings.TrimSpace(val)
			return val, nil
		}
	}
	return "", nil
}
