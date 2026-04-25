# GoGoodwe - CHANGELOG

## v3.4.0 (2025-04-25)

### Security
- Added SSRF protection: validate server-provided API URLs against HTTPS + `*.semsportal.com` domain allowlist before use
- Added HTTP status code checking in login path (`auth/login.go`) and data path (`apihelpers/callmonitorapi.go`)
- Removed credential echo (account, powerstation ID) from `exporthistory` debug output
- UUID validation now accepts uppercase hex characters (case-insensitive)
- Compiled regex patterns once at package level instead of on every call

### Removed (Dead Code Cleanup)
- Deleted `internal/shared/utils/errorhandler.go` — `HandleError()` was defined but never called
- Deleted `internal/features/fetchdata/getdata.go` — 48 lines of fully commented-out code
- Deleted `internal/features/fetchdata/interfaces/semsdata.go` — unused `SemsDataConstraint` type union
- Deleted `internal/features/fetchdata/fetchdata.go` — verbatim duplicate of `apihelpers/callmonitorapi.go`
- Deleted `internal/features/fetchdata/processdata.go` — verbatim duplicate of `utils/processdata.go`
- Removed `LookupReportStruct()` — duplicate factory function never called
- Removed commented-out code from all 7 report-type `GetPowerData` implementations
- Removed commented-out `setPowerPlantHeaders` from `auth/loginutils.go`

### Fixed
- Added 60-second root context timeout; previously used `context.WithCancel()` only, which could hang indefinitely if the API didn't respond
- Removed redundant `ctx.Err()` check after `dataService.GetPowerData()` (errors already propagate from the service call)
- Fixed stale comment: `ApiLoginCredentials` -> `SemsLoginCredentials` in `logincredentials.go`
- Fixed non-idiomatic error messages: removed `**Error:` and `**` decorations from error strings
- Extracted magic strings (`"v2.1.0"`, `"ios"`, `"en"`) to named constants (`apiVersion`, `apiClient`, `apiLanguage`)

### Added
- Comprehensive unit test suite across all internal packages:
  - `internal/shared/utils` — 97.4% coverage (paramcheck, datetime, dateutils, jsonutils, response, httpclient, output, processdata)
  - `internal/shared/auth` — 85.4% coverage (loginutils, credentials, utils, login)
  - `internal/features/fetchdata` — 100% coverage (ParseReportType, LookupMonitorData)
  - `internal/shared/apihelpers` — 44.2% coverage (SSRF validation, parameter guards)
  - 7 report-type packages — constructor tests
  - Overall internal coverage: 64.7%

### Changed
- Updated Go version to 1.26.2

## v3.3.0 (2025-03-15)

### Changed
- Moved from `pkg/` to `internal/` directory structure following Go best practices
- Separated shared utilities (`internal/shared/`) from feature-specific code (`internal/features/`)
- Enhanced CLI with `getdata` and `exporthistory` subcommands
- Updated all tests to reflect new command structure

## v3.2.1 (2025-01-20)

### Added
- New `kpidata` report type (Key Performance Indicators): monthly generation, power, income, yield metrics
- Unified data processing across all report types using `ProcessData()`

### Changed
- Updated `fastjson` from v1.6.4 to v1.6.7

## v3.0.0 (2024-07-04)

- added new report types
- added new command line report types
- Minor performance and optimisation fixes

## v2.0.6 (2024-06-13)

- renamed SemsDataConstraint
- refactored API and monitor data into Internal
- Minor performance and optimisation fixes
- added contexts for better error signals between function calls

## v2.0.5 (2024-05-10)

- Optimised and refactored code to reduce CPU usage and memory pressure and increased speed.

## v2.0.3 (2024-04-26)

- Minor fixes
-

## v2.0.0 (2024-01-22)

- Major refactoring to cleanup project structure.
- Simplified package structure.
- refactored structs to use pointers more efficiently.

## v1.4.0 (2023-08-30)

- Major refactoring to move non-shared code to /internal folder
- Abstracted core away from main()

## v1.1.0 (2023-08-16)

- refactored code to make errors bubble back to main package and better error reporting/logging
- refactored main package to include run() method
- removed staticcheck as there is a bug with Go v1.21

## v1.0.0 (2023-08-10)

- initial version 1.0 release
