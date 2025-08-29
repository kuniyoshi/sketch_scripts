sketch_scripts
================

Minimal Go command-line tool scaffold.

Quick start
-----------

- Build: `mkdir -p .gocache && GOCACHE=$PWD/.gocache go build -o sketch_scripts .`
- Run: `./sketch_scripts -v` or `./sketch_scripts hello -name Alice`

Notes
-----

- Module: initialized as `sketch_scripts`. You can change this in `go.mod` later to a full path (e.g., `github.com/<you>/sketch_scripts`).
- Version info: override at build time with ldflags:
  `GOCACHE=$PWD/.gocache go build -ldflags "-X main.version=0.1.1 -X main.commit=$(git rev-parse --short HEAD) -X main.date=$(date -u +%Y-%m-%d)" -o sketch_scripts .`
- Caching: this repo uses a local `.gocache` to avoid sandboxed writes to the default system cache.

