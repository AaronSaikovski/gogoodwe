# GoGoodwe - CHANGELOG

## v2.0.2 (2024-04-62)

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
