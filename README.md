<div align="center">

## GoGoodwe v3.0.0

A command line tool to query the GOODWE SEMS Solar Inverter APIs - written in 100% Go.

[![Build Status](https://github.com/AaronSaikovski/gogoodwe/workflows/build/badge.svg)](https://github.com/AaronSaikovski/gogoodwe/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/gogoodwe)](LICENSE)

</div>

### Software Requirements:

- [Go v1.22.X](https://www.go.dev/dl/) or later needs to be installed to build the code.
- [Azure CLI tools](https://learn.microsoft.com/en-us/cli/azure/) 2.50 or later
- [Taskfile](https://taskfile.dev/) to run the build chain commands listed below.

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

## Usage

Determine the Station ID from the GOODWE site as follows. Open the [Sems Portal](https://www.semsportal.com). The Plant Status will reveal the Station ID in the URL. Example:

    https://www.semsportal.com/powerstation/powerstatussnmin/11112222-aaaa-bbbb-cccc-ddddeeeeeffff

Then the Station ID is `11112222-aaaa-bbbb-cccc-ddddeeeeeffff`.

From the command line the usage is pretty simple:

The Report Type corresponds to the type of API call and Report that is generated:

- (0)-Detail - Fully detailed report.
- (1)-Summary - Summary Data report (reduced information).
- (2)-Point - Inverter All points data.
- (3)-Plant - Plant Detail By Powerstation Id.
- (4)-PlantChart - Plant Chart data for use in Charts and Graphs.
- (5)-PowerFlow - Powerflow Summary data

```bash
##Note the use of single quotes ''
./gogoodwe  --account '<user@email.com>' \
            --pwd '<password>' \
            --powerstationid '<powerstation id>' \
            --reporttype '<report type>'

# Or
./gogoodwe  -a '<user@email.com>' \
            -p '<password>' \
            -i '<powerstation id>' \
             -r '<report type>'

##w
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

- The Powerchart report is returning no/blank values - investigating.
