BIN ?= sketch_scripts
PKG ?= .
GO ?= go
GOCACHE_DIR ?= $(CURDIR)/.gocache
XDG_CACHE_HOME_DIR ?= $(CURDIR)/.cache

# Ensure Go uses a local cache to avoid sandbox issues
export GOCACHE := $(GOCACHE_DIR)
export XDG_CACHE_HOME := $(XDG_CACHE_HOME_DIR)

.PHONY: build run fmt vet staticcheck

build:
	@mkdir -p "$(GOCACHE_DIR)"
	$(GO) build -o $(BIN) $(PKG)

run: fmt vet staticcheck build
	./$(BIN) $(ARGS)

fmt:
	gofmt -s -w .

vet:
	@mkdir -p "$(GOCACHE_DIR)"
	$(GO) vet ./...

# Runs staticcheck if available; instructs install otherwise
STATICCHECK := $(shell command -v staticcheck 2>/dev/null)
staticcheck:
ifdef STATICCHECK
	@mkdir -p "$(GOCACHE_DIR)" "$(XDG_CACHE_HOME_DIR)" "$(CURDIR)/.home/Library/Caches"
	HOME="$(CURDIR)/.home" $(STATICCHECK) ./...
else
	@echo "[warn] staticcheck not found. Install with:" 1>&2
	@echo "  GOBIN=$$HOME/go/bin $(GO) install honnef.co/go/tools/cmd/staticcheck@latest" 1>&2
	@echo "Skipping staticcheck for now." 1>&2
endif
