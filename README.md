<div align="center">

# GoGoodwe

A command line tool and query the GOODWE SEMS Inverter APIs - written in 100% Go.

[![Build Status](https://github.com/AaronSaikovski/gogoodwe/workflows/build/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

## Installation

The toolchain is mainly driven by the Makefile.

```bash
help         - Display help about make targets for this Makefile
release      - Builds the project in preparation for (local)release
goreleaser   - Builds the project in preparation for release
docs         - updates the swagger docs
build        - Builds the project in preparation for debug
run          - builds and runs the program on the target platform
clean        - Remove the old builds and any debug information
test         - executes unit tests
deps         - fetches any external dependencies and updates
vet          - Vet examines Go source code and reports suspicious constructs
staticcheck  - Runs static code analyzer staticcheck - currently broken
seccheck     - Code vulnerability check
lint         - format code and tidy modules
```

To get started type,

- make dep - to fetch all dependencies
- make build - to build debug version for your target environment architecture
- make release - Builds a release version for your target environment architecture

## Usage

Determine the Station ID from the GOODWE site as follows. Open the [Sems Portal](https://www.semsportal.com). The Plant Status will reveal the Station ID in the URL. Example:

    https://www.semsportal.com/powerstation/powerstatussnmin/11112222-aaaa-bbbb-cccc-ddddeeeeeffff

Then the Station ID is `11112222-aaaa-bbbb-cccc-ddddeeeeeffff`.

From the command line the usage is pretty simple:

```bash
##Note the use of single quotes ''
./gogoodwe --account '<user@email.com>' --pwd '<password>' --powerstationid '<powerstation id>'

# Or
./gogoodwe -a '<user@email.com>' -p '<password>' -i '<powerstation id>'
```

To get the help on using the command line tool, type:

```bash
./gogoodwe --help

# Or
./gogoodwe -h
```

## Contributions

Please feel free to lodge an [issue or pull request on GitHub](https://github.com/AaronSaikovski/gogoodwe/issues).

## Thanks

- Originally based off the work of [Mark Ruys and his gw2pvo software](https://github.com/markruys/gw2pvo) - and [James Hodgkinson and his pygoodwe](https://github.com/yaleman/pygoodwe)

## Disclaimer

GOODWE access is based on the undocumented API used by mobile apps. This could break at any time.

## Known Issues

**None at time of release.**
