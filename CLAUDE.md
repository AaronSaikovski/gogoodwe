# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GoGoodwe is a Go CLI tool that queries the GoodWe SEMS (Solar Energy Management System) Portal APIs to retrieve real-time and historical data from solar inverters. Built with Cobra for CLI, fastjson for JSON parsing, and aurora for colored terminal output.

## Build & Development Commands

This project uses [Taskfile](https://taskfile.dev/) (v3) for task automation:

```bash
task build          # Debug build → bin/gogoodwe.exe
task release        # Release build with stripped symbols
task run            # go run the app
task test           # Run tests: go test -v ./tests/...
task lint           # go fmt + go mod tidy
task vet            # go vet ./...
task staticcheck    # staticcheck ./... (may be broken)
task seccheck       # govulncheck ./...
task deps           # go mod tidy + download + update
task clean          # Clean build artifacts
task generate       # Update embedded version via get_version.sh
task goreleaser     # Cross-platform snapshot release
```

Run a single test: `go test -v -run TestParseReportType ./tests/...`

## Architecture

### Entry Point & CLI

- `cmd/gogoodwe/main.go` — entry point; embeds version from `version.txt` via `go:generate`
- `cmd/gogoodwe/app/cmd.go` — Cobra command definitions: root, `getdata`, `exporthistory`
- `cmd/gogoodwe/app/fetchdata.go` — `RunGetData()` orchestrates login → fetch → output

### Internal Packages (`internal/`)

**Features** (`internal/features/fetchdata/`):

- `interfaces/` — `PowerData` and `SemsLogin` interfaces for loose coupling
- `common/enums.go` — report type constants (Detail=0 through KPIData=6)
- `lookupmonitordata.go` — factory that returns the correct `PowerData` implementation based on report type
- `parsereporttype.go` — converts report type strings ("detail", "summary", etc.) to integers
- 7 report-type packages (monitordetail, monitorsummary, inverterallpoint, plantdetail, plantpowerchart, powerflow, currentkpidata) — each implements `PowerData` interface

**Shared** (`internal/shared/`):

- `auth/` — SEMS login credentials, API authentication (CrossLogin endpoint), token management
- `apihelpers/` — generic HTTP API caller (`FetchMonitorAPIData`)
- `utils/` — HTTP client config (connection pooling), JSON helpers, input validation (email/UUID regex), colored output

### Data Flow

1. Validate inputs (email format, UUID format)
2. Parse report type string → enum int
3. Login via SEMS CrossLogin API → get auth token + API endpoint
4. Factory lookup → get `PowerData` implementation for requested report type
5. Call data API with auth headers → unmarshal response → format and print colored JSON

### Testing

Tests live in `tests/cmd/gogoodwe/app/cmd_test.go` using black-box pattern (`package app_test`). 28 test cases covering report type parsing, command structure, and flag validation.

## Key Patterns

- **Interface-based design**: `PowerData` interface enables adding new report types by implementing `GetPowerData(ctx, loginInfo) error`
- **Factory pattern**: `LookupMonitorData()` maps report type int → concrete implementation
- **Embedded version**: version string embedded at build time via `go:generate` + `go:embed`
- **Global HTTP clients**: reused across requests with connection pooling for performance

## Workflow Orchestration

### 1. Plan Node Default

- Enter plan mode for ANY non-trivial task (3+ steps or architectural decisions)
- If something goes sideways, STOP and re-plan immediately – don't keep pushing
- Use plan mode for verification steps, not just building
- Write detailed specs upfront to reduce ambiguity

### 2. Subagent Strategy

- Use subagents liberally to keep main context window clean
- Offload research, exploration, and parallel analysis to subagents
- For complex problems, throw more compute at it via subagents
- One task per subagent for focused execution

### 3. Self-Improvement Loop

- After ANY correction from the user: update `tasks/lessons.md` with the pattern
- Write rules for yourself that prevent the same mistake
- Ruthlessly iterate on these lessons until mistake rate drops
- Review lessons at session start for relevant project

### 4. Verification Before Done

- Never mark a task complete without proving it works
- Diff behavior between main and your changes when relevant
- Ask yourself: "Would a staff engineer approve this?"
- Run tests, check logs, demonstrate correctness

### 5. Demand Elegance (Balanced)

- For non-trivial changes: pause and ask "is there a more elegant way?"
- If a fix feels hacky: "Knowing everything I know now, implement the elegant solution"
- Skip this for simple, obvious fixes – don't over-engineer
- Challenge your own work before presenting it

### 6. Autonomous Bug Fixing

- When given a bug report: just fix it. Don't ask for hand-holding
- Point at logs, errors, failing tests – then resolve them
- Zero context switching required from the user
- Go fix failing CI tests without being told how

## Task Management

1. **Plan First**: Write plan to `tasks/todo.md` with checkable items
2. **Verify Plan**: Check in before starting implementation
3. **Track Progress**: Mark items complete as you go
4. **Explain Changes**: High-level summary at each step
5. **Document Results**: Add review section to `tasks/todo.md`
6. **Capture Lessons**: Update `tasks/lessons.md` after corrections

## Core Principles

- **Simplicity First**: Make every change as simple as possible. Impact minimal code.
- **No Laziness**: Find root causes. No temporary fixes. Senior developer standards.
- **Minimal Impact**: Changes should only touch what's necessary. Avoid introducing bugs.
