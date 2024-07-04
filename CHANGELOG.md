# GoGoodwe - CHANGELOG

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
