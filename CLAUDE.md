# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go run main.go        # Run the application
go build              # Build binary
go test ./...         # Run all tests
go test -run TestName # Run a single test
```

## Architecture

Minimal single-package Go application. `main.go` is the entry point with a `main()` function in `package main`. Module name is `another_hello` (defined in `go.mod`).
