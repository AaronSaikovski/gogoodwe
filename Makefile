# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"
TARGET = gogoodwe
MAINAPPPATH = ./cmd/${TARGET}/main.go

default: help

.PHONY: help
## help - Display help about make targets for this Makefile
help:
	@cat Makefile | grep '^## ' --color=never | cut -c4- | sed -e "`printf 's/ - /\t- /;'`" | column -s "`printf '\t'`" -t


.PHONY: release
## release - Builds the project in preparation for (local)release
release: vet lint seccheck
	go build $(GOFLAGS) -o bin/${TARGET} ${MAINAPPPATH}
	file bin/${TARGET}


.PHONY: goreleaser
## goreleaser - Builds the project in preparation for release
goreleaser:
	goreleaser release --snapshot --clean 


.PHONY: docs
## docs - updates the swagger docs
docs:	
	swag init


.PHONY: build
## build - Builds the project in preparation for debug
build:
	go build -o bin/${TARGET} ${MAINAPPPATH}
	file bin/${TARGET}


.PHONY: run
## run - builds and runs the program on the target platform
run:
	go run ${MAINAPPPATH}


.PHONY: clean
## clean - Remove the old builds and any debug information
clean:
	go clean -cache
	go clean
	rm -rf dist
	rm bin/${TARGET}


.PHONY: test
## test - executes unit tests
test:
	go test -v ./test/...


.PHONY: deps
## deps - fetches any external dependencies and updates
deps:
	go mod tidy
	go mod download
	go get -u ./...


.PHONY: vet
## vet - Vet examines Go source code and reports suspicious constructs
vet:
	go vet ./...


.PHONY: staticcheck
## staticcheck - Runs static code analyzer staticcheck - currently broken
staticcheck:
	staticcheck ./...


.PHONY: seccheck
## seccheck - Code vulnerability check
seccheck:	
	govulncheck ./...


.PHONY: lint
## lint - format code and tidy modules
lint:
	go fmt ./...
	go mod tidy -v
