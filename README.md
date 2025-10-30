<div align="center">

## GoGoodwe v3.1.1

A high-performance command-line tool to query GOODWE SEMS (Solar Energy Management System) APIs - written in 100% Go.



[![Build Status](https://github.com/AaronSaikovski/gogoodwe/workflows/build/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

### Software Requirements:

- [Go v1.25.3](https://www.go.dev/dl/) or later needs to be installed to build the code.
- [Taskfile](https://taskfile.dev/) to run the build chain commands listed below.

### Dependencies:

GoGoodwe uses minimal external dependencies for a lightweight binary:
- `github.com/alexflint/go-arg` - CLI argument parsing
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
cmd/gogoodwe/           - Command-line application entry point
  └── app/              - Application logic for login and data fetching
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

The Report Type parameter specifies which type of data report to generate:

- (0)-Detail - Fully detailed report.
- (1)-Summary - Summary Data report (reduced information).
- (2)-Point - Inverter All points data.
- (3)-Plant - Plant Detail By Powerstation Id.
- (4)-PlantChart - Plant Chart data for use in Charts and Graphs - **Not Implemented due to an undocumented GoodWe API Change**
- (5)-PowerFlow - Powerflow Summary data

```bash
# Using long arguments (note the use of single quotes)
./gogoodwe --account 'user@email.com' \
           --pwd 'password' \
           --powerstationid '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
           --reporttype '0'

# Or using short arguments
./gogoodwe -a 'user@email.com' \
           -p 'password' \
           -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff' \
           -r '0'

# Report type is optional (defaults to 0 - Detail report)
./gogoodwe -a 'user@email.com' -p 'password' -i '11112222-aaaa-bbbb-cccc-ddddeeeeeffff'
```

To get the help on using the command line tool, type:

```bash
./gogoodwe --help

# Or
./gogoodwe -h
```

## Recent Changes

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
