<div align="center">

## GoGoodwe v3.2.1

A high-performance command-line tool to query GOODWE SEMS (Solar Energy Management System) APIs - written in 100% Go.



[![Build Status](https://github.com/AaronSaikovski/gogoodwe/workflows/build/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

### Software Requirements:

- [Go v1.25.3](https://www.go.dev/dl/) or later needs to be installed to build the code.
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
  │   │   ├── login.go          - Login logic and credentials
  │   │   ├── loginutils.go     - Authentication utilities
  │   │   └── utils.go          - Auth-related helper functions
  │   ├── apihelpers/   - HTTP request/response handling
  │   │   └── callmonitorapi.go - API communication layer
  │   └── utils/        - Common utilities
  │       ├── jsonutils.go      - JSON parsing and formatting
  │       ├── httpclient.go     - HTTP client configuration
  │       ├── processdata.go    - Data processing pipeline
  │       └── datetime.go       - Date/time utilities
  └── features/
      └── fetchdata/    - Data fetching feature
          ├── fetchdata.go      - Core fetch logic
          ├── parsereporttype.go - Report type parsing
          ├── lookupmonitordata.go - Report type routing
          ├── common/           - Shared constants and types
          │   └── enums.go      - Report type constants
          ├── interfaces/       - Interface definitions
          │   ├── powerdata.go  - Data fetching interface
          │   ├── semsdata.go   - Type constraints
          │   └── semslogin.go  - Login interface
          └── [models]/         - Data structures for each report type
              ├── currentkpidata/     - KPI monitoring data
              ├── inverterallpoint/   - All inverter point data
              ├── monitordetail/      - Detailed monitoring data
              ├── monitorsummary/     - Summary monitoring data
              ├── plantdetail/        - Plant detail data
              ├── plantpowerchart/    - Plant power chart data
              └── powerflow/          - Power flow data

tests/                  - Comprehensive test suite
  └── cmd/
      └── gogoodwe/
          └── app/
              └── cmd_test.go - CLI command and subcommand tests
```

### Performance Optimizations

GoGoodwe includes several performance enhancements:
- **HTTP Connection Pooling**: Reusable HTTP client with optimized transport settings (MaxIdleConns: 100, MaxConnsPerHost: 100)
- **Efficient JSON Parsing**: Uses `fastjson` for high-performance JSON processing without double marshaling
- **Optimized Timeouts**: HTTP response header timeout of 10 seconds with TLS handshake timeout
- **Context Management**: Proper context handling with 60-second application timeout

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

GoGoodwe includes comprehensive test coverage for CLI functionality:

### Running Tests

```bash
# Run all tests
go test ./tests/... -v

# Run CLI command tests
go test ./tests/cmd/gogoodwe/app -v

# Run tests with coverage report
go test ./tests/cmd/gogoodwe/app -cover

# Generate HTML coverage report
go test ./tests/cmd/gogoodwe/app -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Suite

- **Location**: `tests/cmd/gogoodwe/app/cmd_test.go`
- **Coverage**: Comprehensive test cases for all functionality
- **Focus Areas**:
  - Report type parsing (string and numeric formats)
  - Root command structure and subcommands
  - CLI flag validation for getdata and exporthistory
  - Command initialization and configuration
  - Edge cases and error handling

See `tests/README.md` for detailed test documentation.

## Recent Changes

### Version 3.3.0 (Current)
- **Project Restructuring** - Major refactoring for better code organization:
  - Moved from `pkg/` to `internal/` directory structure following Go best practices
  - Separated shared utilities (`internal/shared/`) from feature-specific code (`internal/features/`)
  - Consolidated authentication utilities in `internal/shared/auth/`
  - Unified data processing in `internal/shared/utils/`
- **Command Structure** - Enhanced CLI with subcommands:
  - `getdata` - Retrieve real-time inverter data (replaces previous direct command)
  - `exporthistory` - Export historical data (foundation for future Excel export feature)
- **Code Cleanup** - Removed duplicate utility functions and improved import organization
- **Test Suite Updates** - Updated all tests to reflect new command structure and subcommands

### Version 3.2.1
- **KPI Data Report Type** - New `kpidata` report type providing Key Performance Indicator metrics:
  - Monthly generation data
  - Current power (Pac) and total power metrics
  - Income tracking (daily and total)
  - Yield rate calculations
  - Currency information
- **Code Refactoring** - Unified data processing across all report types using `ProcessData()` for improved consistency and maintainability
- **Dependency Updates** - Updated `fastjson` from v1.6.4 to v1.6.7 for improved JSON parsing performance

### Previous Updates
- **Cobra CLI Framework** - Replaced `go-arg` with [Cobra](https://github.com/spf13/cobra) for robust CLI argument parsing
- **String-based Report Types** - Support for human-readable report type names (`detail`, `summary`, `point`, etc.) alongside numeric values for backward compatibility
- **Comprehensive Test Suite** - Added 28 test cases covering CLI functionality with dedicated `tests/` directory
- **Improved Documentation** - Updated help text and command descriptions with Cobra

### Version 3.1.1
- **Updated Go version** to 1.25.3 for latest language improvements and security patches
- **Performance refactoring** with optimized HTTP client setup and connection pooling
- **Improved context management** for better timeout handling and resource cleanup
- **Code organization** - Renamed `root.go` to `main.go` for better clarity

## Contributions

Please feel free to lodge an [issue or pull request on GitHub](https://github.com/AaronSaikovski/gogoodwe/issues).

## Thanks

- Originally based off the work of [Mark Ruys and his gw2pvo software](https://github.com/markruys/gw2pvo) - and [James Hodgkinson and his pygoodwe](https://github.com/yaleman/pygoodwe)

## Disclaimer

GOODWE access is based on the undocumented API used by mobile apps. This could break at any time.

## Known Issues

- The Powerchart report is returning no/blank values - investigating.
