# GoGoodwe Tests

This directory contains test cases for the GoGoodwe CLI application.

## Test Structure

```
tests/
└── cmd/
    └── gogoodwe/
        └── app/
            └── cmd_test.go          # CLI command tests
```

## Running Tests

### Run all tests
```bash
go test ./tests/... -v
```

### Run specific package tests
```bash
go test ./tests/cmd/gogoodwe/app -v
```

### Run tests with coverage
```bash
go test ./tests/cmd/gogoodwe/app -cover
```

### Generate coverage report
```bash
go test ./tests/cmd/gogoodwe/app -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Cases

### cmd_test.go

**TestParseReportType** (17 subtests)
- Tests string-based report types (detail, summary, point, plant, plantchart, powerflow)
- Tests numeric values (0-5)
- Tests invalid inputs (invalid strings, numbers, empty, negative, wrong case)

**TestNewRootCmd** (3 subtests)
- Tests command creation with different version strings
- Verifies command name, descriptions, and version

**TestRootCmdFlags** (4 subtests)
- Verifies all CLI flags exist with correct names and shortcuts
- Checks account, password, powerstationid, and reporttype flags

**TestRootCmdFlagDefaults** (1 subtest)
- Validates default values for flags
- Ensures reporttype defaults to "detail"

**TestParseReportTypeEdgeCases** (5 subtests)
- Tests boundary conditions and edge cases
- Whitespace handling, case sensitivity, partial matches

**TestRootCmdHasRunE** (1 subtest)
- Verifies the root command has an execution handler

**TestRootCmdDescription** (1 subtest)
- Validates command descriptions are meaningful

## Test Results

- **Total Tests:** 28
- **Status:** All PASS ✅
- **Package:** `github.com/AaronSaikovski/gogoodwe/tests/cmd/gogoodwe/app`

## Notes

- Tests use black-box testing pattern (app_test package) for better isolation
- Tests validate the public API of the CLI command package
- ParseReportType function is exported to allow external testing
