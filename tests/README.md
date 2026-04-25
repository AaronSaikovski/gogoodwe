# GoGoodwe Tests

This directory contains black-box CLI tests using the `app_test` package. The majority of unit tests live alongside the source code in the `internal/` packages.

## Test Structure

```
tests/
└── cmd/gogoodwe/app/
    └── cmd_test.go          # CLI command structure and flag tests

internal/ (co-located tests)
├── shared/
│   ├── auth/
│   │   ├── login_test.go            # SemsLogin validation, context cancellation
│   │   ├── logincredentials_test.go  # Credential factory
│   │   ├── loginutils_test.go       # Login info validation, response checking, headers
│   │   └── utils_test.go            # Token JSON, header setting, powerstation JSON
│   ├── apihelpers/
│   │   └── callmonitorapi_test.go   # SSRF URL validation, nil/empty guards
│   └── utils/
│       ├── paramcheck_test.go       # Email and UUID validation
│       ├── datetime_test.go         # Date formatting
│       ├── dateutils_test.go        # Date range calculation
│       ├── jsonutils_test.go        # JSON marshal/unmarshal
│       ├── response_test.go         # Response body reading, size limits
│       ├── httpclient_test.go       # HTTP transport configuration
│       ├── output_test.go           # JSON parsing
│       └── processdata_test.go      # Data processing pipeline
└── features/fetchdata/
    ├── parsereporttype_test.go      # Report type string/int conversion
    ├── lookupmonitordata_test.go    # Factory returns correct types
    └── [7 report-type packages]/    # Constructor tests
```

## Running Tests

```bash
# Run all tests
go test -v ./...

# Run only CLI tests
go test -v ./tests/...

# Run only internal package tests
go test -v ./internal/...

# Run with coverage
go test -coverprofile=coverage.out ./internal/...
go tool cover -func=coverage.out

# Run a specific test
go test -v -run TestCheckValidEmail ./internal/shared/utils/...
```

## Test Cases Summary

### CLI Tests (`tests/cmd/gogoodwe/app/cmd_test.go`)

| Test | Subtests | Description |
|------|----------|-------------|
| TestParseReportType | 19 | String/numeric report types, invalid inputs |
| TestParseReportTypeEdgeCases | 5 | Whitespace, case sensitivity, partial matches |
| TestNewRootCmd | 3 | Command creation with version strings |
| TestRootCmdHasSubcommands | 1 | Verifies getdata and exporthistory exist |
| TestGetDataSubcommand | 1 | Subcommand name and description |
| TestGetDataFlags | 4 | Flag names and shortcuts |
| TestGetDataFlagDefaults | 1 | Default report type is "detail" |
| TestExportHistorySubcommand | 1 | Subcommand name and description |
| TestExportHistoryFlags | 6 | Flag names and shortcuts |
| TestRootCmdDescription | 1 | Command description validation |

### Internal Package Tests

| Package | Tests | Key Coverage |
|---------|-------|-------------|
| `utils` | 47 | Email/UUID validation, JSON, response body, HTTP transport, dates |
| `auth` | 18 | Credential validation, headers, token JSON, login flow |
| `apihelpers` | 13 | SSRF URL validation, parameter guards |
| `fetchdata` | 31 | Report type parsing, factory pattern, constructors |

**Total: 100+ test cases across all packages**

## Notes

- Tests use table-driven patterns with subtests throughout
- Black-box testing (`app_test` package) for CLI tests
- White-box testing (same package) for internal packages to test unexported functions
- Report-type `GetPowerData`/`GetMonitorData` methods are not unit-tested as they require live API access; they are thin wrappers around the well-tested `apihelpers.FetchMonitorAPIData`
