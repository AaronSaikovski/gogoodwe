<div align="center">

## GoGoodwe v3.4.0

A high-performance command-line tool to query GOODWE SEMS (Solar Energy Management System) APIs - written in 100% Go.



[![Build Status](https://github.com/AaronSaikovski/gogoodwe/workflows/build/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

### Software Requirements:

- [Go v1.26.2](https://www.go.dev/dl/) or later needs to be installed to build the code.
- [Taskfile](https://taskfile.dev/) to run the build chain commands listed below.

### Dependencies:

GoGoodwe uses minimal external dependencies for a lightweight binary:
- `github.com/spf13/cobra` - CLI command framework
- `github.com/logrusorgru/aurora` - Terminal color output
- `github.com/valyala/fastjson` - High-performance JSON parsing

## Installation:

The toolchain is driven by using [Taskfile](https://taskfile.dev/) and all commands are managed via the file `Taskfile.yml`

The list of commands is as follows:

```bash
* build:             Builds the project in preparation for debug.
* clean:             Removes the old builds and any debug information from the source tree.
* deps:              Fetches any external dependencies and updates.
* destroy:           Destroy Azure resources for testing.
* docs:              Updates the swagger docs - For APIs.
* generate:          update binary build version using gogenerate.
* goreleaser:        Builds a cross platform release using goreleaser.
* lint:              Lint, format and tidy code.
* release:           Builds the project in preparation for (local)release.
* run:               Builds and runs the program on the target platform.
* seccheck:          Code vulnerability scanner check.
* staticcheck:       Runs static code analyzer staticcheck.
* test:              Executes unit tests.
* version:           Get the Go version.
* vet:               Vet examines Go source code and reports suspicious constructs.
* watch:             Use air server for hot reloading.
```

Execute using the taskfile utility:

```bash
task <command_from_above_list>
```

To get started type,

- `task deps` - to fetch all dependencies and update all dependencies.
- `task build` - to build debug version for your target environment architecture.
- `task release` - Builds a release version for your target environment architecture - outputs to /bin folder.

## Project Architecture

```
cmd/
  └── gogoodwe/         - Command-line application entry point
      ├── main.go       - Entry point, initializes Cobra root command
      └── app/          - Application command logic
          ├── cmd.go        - Cobra root and subcommand definitions
          ├── fetchdata.go  - GetData command: Login and API data fetching
          └── exportdata.go - ExportHistory command: Historical data export

internal/               - Internal application packages (not for external use)
  ├── shared/           - Shared utilities and helpers
  │   ├── auth/         - Authentication handling for SEMS API
  │   │   ├── login.go            - SEMS CrossLogin API authentication
  │   │   ├── logincredentials.go  - Login credential struct and factory
  │   │   ├── logininfo.go        - Combined login state (credentials + response)
  │   │   ├── loginresponse.go    - API response struct
  │   │   ├── loginutils.go       - Credential validation and header setup
  │   │   └── utils.go            - Token/header JSON generation
  │   ├── apihelpers/   - HTTP request/response handling
  │   │   └── callmonitorapi.go - API communication with SSRF protection
  │   └── utils/        - Common utilities
  │       ├── jsonutils.go      - JSON marshal/unmarshal helpers
  │       ├── httpclient.go     - HTTP transport with connection pooling
  │       ├── processdata.go    - Data processing pipeline (JSON -> output)
  │       ├── response.go       - Response body reading with size limits
  │       ├── output.go         - Colored terminal output via fastjson
  │       ├── paramcheck.go     - Email and UUID input validation
  │       ├── datetime.go       - Date formatting (YYYY-MM-DD)
  │       └── dateutils.go      - Date range calculations
  └── features/
      └── fetchdata/    - Data fetching feature
          ├── parsereporttype.go    - Report type string/int conversion
          ├── lookupmonitordata.go  - Factory: report type -> PowerData impl
          ├── common/               - Shared constants
          │   └── enums.go          - Report type constants (Detail..KPIData)
          ├── interfaces/           - Interface definitions
          │   ├── powerdata.go      - PowerData interface
          │   └── semslogin.go      - SemsLogin interface
          └── [report types]/       - One package per report type
              ├── currentkpidata/     - KPI monitoring data
              ├── inverterallpoint/   - All inverter point data
              ├── monitordetail/      - Detailed monitoring data
              ├── monitorsummary/     - Summary monitoring data
              ├── plantdetail/        - Plant detail data
              ├── plantpowerchart/    - Plant power chart data
              └── powerflow/          - Power flow data

tests/                  - Black-box CLI tests (app_test package)
  └── cmd/gogoodwe/app/cmd_test.go
```

### Performance Optimizations

- **HTTP Connection Pooling**: Reusable HTTP client with optimized transport settings (MaxIdleConns: 100, MaxConnsPerHost: 100)
- **Efficient JSON Parsing**: Uses `fastjson` for high-performance JSON processing
- **Optimized Timeouts**: HTTP response header timeout of 10s, TLS handshake timeout of 10s
- **Context Management**: 60-second root timeout with per-request 20-second timeouts

### Security Features

- **SSRF Protection**: Server-provided API URLs are validated against an HTTPS + domain allowlist before use
- **HTTP Status Checking**: Non-200 responses are rejected with informative errors in both login and data paths
- **Response Size Limiting**: Response bodies are capped at 10MB to prevent memory exhaustion
- **Input Validation**: Email format and UUID format validation on all user inputs

## Usage

GoGoodwe retrieves real-time and historical data from your GoodWe solar inverters via the SEMS (Solar Energy Management System) Portal API.

### Getting Your Station ID

1. Open the [Sems Portal](https://www.semsportal.com)
2. Navigate to your power plant status page
3. The Station ID (UUID format) will appear in the URL:

    `https://www.semsportal.com/powerstation/powerstatussnmin/11112222-aaaa-bbbb-cccc-ddddeeeeeffff`

Your Station ID is: `11112222-aaaa-bbbb-cccc-ddddeeeeeffff`

### Command Line Usage

GoGoodwe uses [Cobra](https://github.com/spf13/cobra) with subcommands for CLI argument parsing.

#### GetData Command

Retrieves real-time data from your SEMS inverter. The Report Type parameter specifies which type of data report to generate:

**Available Report Types:**
- `detail` or `0` - Fully detailed inverter monitoring report (default)
- `summary` or `1` - Summary Data report (reduced information)
- `point` or `2` - Inverter All points data
- `plant` or `3` - Plant Detail By Powerstation Id
- `plantchart` or `4` - Plant Chart data for use in Charts and Graphs (API support varies)
- `powerflow` or `5` - Powerflow Summary data
- `kpidata` or `6` - KPI (Key Performance Indicator) data including generation, power, income, and yield metrics

**Examples:**

```bash
# Using string-based report types (recommended)
./gogoodwe getdata --account 'user@email.com' \
                   --password 'password' \
                   --powerstationid '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
                   --reporttype 'detail'

./gogoodwe getdata -a 'user@email.com' \
                   -p 'password' \
                   -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
                   -r 'summary'

# KPI data report (Key Performance Indicators)
./gogoodwe getdata -a 'user@email.com' \
                   -p 'password' \
                   -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
                   -r 'kpidata'

# Using numeric report types (backward compatible)
./gogoodwe getdata -a 'user@email.com' -p 'password' -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' -r 0

# Report type is optional (defaults to 'detail')
./gogoodwe getdata -a 'user@email.com' -p 'password' -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff'
```

#### ExportHistory Command

Exports historical data from your SEMS inverter to Excel format (coming soon).

**Example:**

```bash
./gogoodwe exporthistory -a 'user@email.com' \
                         -p 'password' \
                         -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
                         --timestart '2024-01-01 00:00' \
                         --timeend '2024-01-31 23:59' \
                         --targets 'Vpv1,Vpv2,Ipv1'
```

### Getting Help

```bash
# Display main help message
./gogoodwe --help
./gogoodwe -h

# Display help for getdata subcommand
./gogoodwe getdata --help

# Display help for exporthistory subcommand
./gogoodwe exporthistory --help

# Display version
./gogoodwe --version
./gogoodwe -v
```

## Testing

GoGoodwe includes comprehensive unit tests across all internal packages.

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run with coverage report
go test -coverprofile=coverage.out ./internal/...
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out

# Run a single test
go test -v -run TestParseReportType ./internal/features/fetchdata/...
```

### Test Coverage

| Package | Coverage |
|---------|----------|
| `internal/shared/utils` | 97.4% |
| `internal/shared/auth` | 85.4% |
| `internal/features/fetchdata` | 100% |
| `internal/shared/apihelpers` | 44.2% |
| Report-type packages (x7) | 16.7% each |

### What's Tested

- **Input validation**: Email format, UUID format (case-insensitive), report type parsing
- **Authentication**: Credential validation, header generation, token JSON, login response checking
- **API helpers**: SSRF URL validation, nil/empty parameter guards, HTTP status handling
- **Utilities**: JSON marshal/unmarshal, response body reading, date formatting, HTTP transport config
- **CLI structure**: Command definitions, flag names, defaults, subcommand presence
- **Factory pattern**: `LookupMonitorData` returns correct types for all 7 report types

See `tests/README.md` for detailed test documentation.

## Recent Changes

### Version 3.4.0 (Current)
- **Security Hardening**:
  - Added SSRF protection: server-provided API URLs are validated against HTTPS + `*.semsportal.com` domain allowlist
  - Added HTTP status code checking in both login and data API paths
  - Removed credential echo from `exporthistory` debug output
  - UUID validation now accepts uppercase hex (case-insensitive)
- **Code Quality**:
  - Removed verbatim code duplication (`FetchMonitorAPIData`, `ProcessData`, `ProcessRawJSON` had identical copies in two locations)
  - Removed dead code: unused `HandleError()`, commented-out `getdata.go`, unused `SemsDataConstraint` interface, duplicate `LookupReportStruct()` factory
  - Cleaned commented-out code from all 7 report-type files
  - Added 60-second root context timeout (was `WithCancel` only, could hang indefinitely)
  - Removed redundant `ctx.Err()` check after service call
  - Extracted magic strings (`"v2.1.0"`, `"ios"`, `"en"`) to package-level constants
  - Fixed non-idiomatic error messages (removed `**Error:` decorations)
  - Fixed stale comment (`ApiLoginCredentials` -> `SemsLoginCredentials`)
  - Compiled regex patterns once at package level instead of per-call
- **Testing**: Comprehensive unit test suite added across all internal packages (64.7% overall coverage, 97.4% for utils, 85.4% for auth, 100% for fetchdata core)
- **Updated Go version** to 1.26.2

### Version 3.3.0
- **Project Restructuring** - Moved from `pkg/` to `internal/` directory structure following Go best practices
- **Command Structure** - Enhanced CLI with `getdata` and `exporthistory` subcommands
- **Code Cleanup** - Removed duplicate utility functions and improved import organization

### Version 3.2.1
- **KPI Data Report Type** - New `kpidata` report type with monthly generation, power, income, and yield metrics
- **Code Refactoring** - Unified data processing across all report types using `ProcessData()`
- **Dependency Updates** - Updated `fastjson` from v1.6.4 to v1.6.7

### Previous Versions
- v3.1.1 - Go 1.25.3, performance refactoring, improved context management
- v3.0.0 - Cobra CLI, string-based report types, comprehensive test suite

## Contributions

Please feel free to lodge an [issue or pull request on GitHub](https://github.com/AaronSaikovski/gogoodwe/issues).

## Thanks

- Originally based off the work of [Mark Ruys and his gw2pvo software](https://github.com/markruys/gw2pvo) - and [James Hodgkinson and his pygoodwe](https://github.com/yaleman/pygoodwe)

## Disclaimer

GOODWE access is based on the undocumented API used by mobile apps. This could break at any time.

## Known Issues

- The Powerchart report is returning no/blank values - investigating.
