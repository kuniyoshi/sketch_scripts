package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	// Allow directory to be specified via -dir option; default is "sketches".
	dir := flag.String("dir", "sketches", "sketch files directory")
	flag.Parse()

	err := listSketches(*dir)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
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

	// Determine the maximum filename width (in runes) for alignment
	maxName := 0
	for _, path := range files {
		w := utf8.RuneCountInString(filepath.Base(path))
		if w > maxName {
			maxName = w
		}
	}

	for _, path := range files {
		desc, err := extractSketchDescription(path)
		if err != nil {
			// Prefix the filename with a visual marker to distinguish it from wrapped descriptions
			fmt.Printf("▌%-*s  %s\n", maxName, filepath.Base(path), "(read error)")
			continue
		}
		if desc == "" {
			desc = "(no description)"
		}
		// Prefix the filename with a visual marker to distinguish it from wrapped descriptions
		fmt.Printf("▌%-*s  %s\n", maxName, filepath.Base(path), desc)
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
			val = strings.TrimSpace(val)
			return val, nil
		}
	}
	return "", nil
}
