set quiet

_default: _help

_help:
    just --list

# Format go files
fmt:
    golangci-lint fmt -v

# Lint go files
lint:
    golangci-lint run -v

# Generate files
generate:
    go generate ./...

# Run tests
test:
    go test -v -failfast -race -timeout 10m ./...
