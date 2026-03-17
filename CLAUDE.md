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

Minimal single-package Go application. `main.go` is the entry point in `package main`. The `greet(name string) string` function holds the core logic and is tested directly in `main_test.go`. Module name is `another_hello` (defined in `go.mod`).

Accepts an optional name as the first CLI argument:

```
go run main.go          # Hello, World!
go run main.go Alice    # Hello, Alice!
```

## Pushing to GitHub

HTTPS push requires `gh` as the credential helper. If push fails with auth errors, run:

```
gh auth setup-git
```
