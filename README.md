<div align="center">

## GoGoodwe

A high-performance command-line tool to query GOODWE SEMS (Solar Energy Management System) APIs - written in 100% Go.

[![Test](https://github.com/AaronSaikovski/gogoodwe/actions/workflows/test.yml/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions/workflows/test.yml)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

### Software Requirements:

- [Go v1.26.4](https://www.go.dev/dl/) or later needs to be installed to build the code.
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
* docker-build:      Builds a Docker image based on the Dockerfile.
* docker-run:        Runs the Docker container.
* generate:          Update binary build version using gogenerate.
* goreleaser:        Builds a cross platform release using goreleaser.
* lint:              Lint, format, tidy code, and run go fix.
* release:           Builds the project in preparation for (local) release.
* run:               Builds and runs the program on the target platform.
* seccheck:          Code vulnerability scanner check (govulncheck).
* staticcheck:       Runs static code analyzer staticcheck.
* test:              Executes unit tests.
* vet:               Vet examines Go source code and reports suspicious constructs.
```

Execute using the taskfile utility:

```bash
task <command_from_above_list>
```

To get started type,

- `task deps` - to fetch all dependencies and update all dependencies.
- `task build` - to build debug version for your target environment architecture.
- `task release` - Builds a release version for your target environment architecture - outputs to /bin folder.

### macOS Security

If you downloaded a pre-built binary from a GitHub Release and macOS blocks it with "App can't be opened because Apple cannot check it for malicious software", run:

```bash
xattr -d com.apple.quarantine ./gogoodwe
```

Alternatively, right-click the binary and select **Open** from the context menu, then confirm when prompted.

## Project Architecture

```
cmd/gogoodwe/           - Command-line application entry point
  ├── main.go           - Entry point, initializes Cobra root command
  └── app/
      ├── cmd.go        - Cobra command definitions and flag configuration
      ├── fetchdata.go  - Login and API data fetching logic
      ├── enums.go      - Report type constants
      └── lookupmonitordata.go - Report type routing

tests/                  - Legacy test directory (CLI tests)
  └── cmd/
      └── gogoodwe/
          └── app/
              └── cmd_test.go - CLI command tests
  └── README.md         - Test documentation

pkg/                    - Core library packages
  ├── auth/             - Authentication handling for SEMS API
  ├── apihelpers/       - HTTP request/response handling and API communication
  ├── models/           - Data structures for each report type
  ├── interfaces/       - Interface definitions for type safety
  └── utils/            - Utilities for JSON, HTTP clients, and formatting
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

GoGoodwe uses [Cobra](https://github.com/spf13/cobra) for CLI argument parsing. The Report Type parameter specifies which type of data report to generate:

**Available Report Types:**

- `detail` or `0` - Fully detailed inverter monitoring report
- `summary` or `1` - Summary Data report (reduced information)
- `point` or `2` - Inverter All points data
- `plant` or `3` - Plant Detail By Powerstation Id
- `plantchart` or `4` - Plant Chart data for use in Charts and Graphs (API support varies)
- `powerflow` or `5` - Powerflow Summary data

**Examples:**

```bash
# Using string-based report types (recommended)
./gogoodwe --account 'user@email.com' \
           --password 'password' \
           --powerstationid '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
           --reporttype 'detail'

./gogoodwe -a 'user@email.com' \
           -p 'password' \
           -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
           -r 'summary'

# Using numeric report types (backward compatible)
./gogoodwe -a 'user@email.com' -p 'password' -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' -r 0

# Report type is optional (defaults to 'detail')
./gogoodwe -a 'user@email.com' -p 'password' -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff'
```

### Getting Help

```bash
# Display help message
./gogoodwe --help
./gogoodwe -h

# Display version
./gogoodwe --version
./gogoodwe -v
```

## Testing

GoGoodwe includes comprehensive test coverage across CLI functionality, utilities, authentication, and API helpers:

### Running Tests

```bash
# Run all tests
go test ./... -v

# Run with Taskfile
task test

# Run tests with coverage report
go test ./... -cover

# Generate HTML coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Static analysis
task staticcheck
task seccheck
```

### Test Suite

- **CLI Tests**: `tests/cmd/gogoodwe/app/cmd_test.go` — Report type parsing, flag validation, command initialization
- **Utils Tests**: `cmd/gogoodwe/utils/` and `pkg/utils/` — Email/powerstation validation, JSON marshaling, HTTP transport, date formatting
- **Auth Tests**: `pkg/auth/` — Login credentials, header setting, login response validation
- **API Tests**: `pkg/apihelpers/` — Token generation, header configuration, power station ID JSON

All tests use table-driven test patterns following Go best practices.

## Recent Changes

### Latest Updates

- **CI/CD Pipeline** — GitHub Actions workflows for automated testing (on every push/PR) and GoReleaser releases (on tags)
- **Expanded Test Suite** — Added 50+ test cases across utils, auth, and apihelpers packages using table-driven patterns
- **Static Analysis** — `staticcheck` and `govulncheck` integrated into CI pipeline and Taskfile
- **GoReleaser v2** — Updated to GoReleaser v2 with automatic cross-platform builds (Linux, macOS, Windows) and GitHub release creation
- **Taskfile Fixes** — `staticcheck` and `seccheck` tasks now use `go run` for zero-install setup
- **Cross-platform Taskfile** — macOS, Linux, and Windows support with platform-specific command routing
- **`go fix` integration** — automatic code fixes applied during lint
- **Docker tasks** — `docker-build` and `docker-run` for containerized builds
- **Security check task** — `seccheck` for `govulncheck` vulnerability scanning
- **Updated Go version** to 1.26.4
- **Cobra CLI Framework** - Replaced `go-arg` with [Cobra](https://github.com/spf13/cobra) for robust CLI argument parsing
- **String-based Report Types** - Support for human-readable report type names (`detail`, `summary`, `point`, etc.) alongside numeric values for backward compatibility
- **Improved Documentation** - Updated help text and command descriptions with Cobra

### Version 3.2.0 (Unreleased)

- **CI/CD Pipeline** — GitHub Actions for automated testing and GoReleaser v2 releases
- **Expanded Test Suite** — 50+ test cases across utils, auth, and apihelpers packages
- **Static Analysis** — `staticcheck` and `govulncheck` in CI pipeline
- **Cross-platform Taskfile** — fully compatible with macOS, Linux, and Windows
- **`go fix` integration** — automatic code fixes applied during lint
- **Docker tasks** — `docker-build` and `docker-run` for containerized builds
- **Security check task** — `seccheck` for `govulncheck` vulnerability scanning
- **Updated Go version** to 1.26.4
- **Updated fastjson dependency** to v1.6.10

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
