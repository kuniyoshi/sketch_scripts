sketch_scripts
================

Minimal Go command-line tool scaffold.

Quick start
-----------

- Build: `mkdir -p .gocache && GOCACHE=$PWD/.gocache go build -o sketch_scripts .`
- Run: `./sketch_scripts -v` or `./sketch_scripts hello -name Alice`

Notes
-----

- Module: `github.com/kuniyoshi/sketch_scripts` (set in `go.mod`).
- Version info: override at build time with ldflags:
  `GOCACHE=$PWD/.gocache go build -ldflags "-X main.version=0.1.1 -X main.commit=$(git rev-parse --short HEAD) -X main.date=$(date -u +%Y-%m-%d)" -o sketch_scripts .`
- Caching: this repo uses a local `.gocache` to avoid sandboxed writes to the default system cache.
