# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"
TARGET = gogoodwe
MAINAPPPATH = ./cmd/${TARGET}/main.go

## help - Display help about make targets for this Makefile
help:
	@cat Makefile | grep '^## ' --color=never | cut -c4- | sed -e "`printf 's/ - /\t- /;'`" | column -s "`printf '\t'`" -t

## localrelease -  Builds the project in preparation for (local) release
localrelease: vet lint seccheck unittest
	go build $(GOFLAGS) -o bin/${TARGET} ${MAINAPPPATH}
	file bin/${TARGET}

## release - Builds the project in preparation for release
release:
	goreleaser release --snapshot --clean 
	
## debug - Builds the project in preparation for debug
build:
	go build -o bin/${TARGET} ${MAINAPPPATH}
	file bin/${TARGET}

## buildandrun - builds and runs the program on the target platform
buildandrun: build
	./bin/${TARGET}

## run - runs main.go for testing
run: dep
	go run ${MAINAPPPATH}


## clean - Remove the old builds and any debug information
clean:
	go clean
	rm -rf dist
	rm bin/${TARGET}

## unittest - executes unit tests
unittest:
	go test -v ./test/...

## dep - fetches any external dependencies
dep:
	go mod tidy
	go mod download

## vet - Vet examines Go source code and reports suspicious constructs
vet:
	go vet ./...

## staticcheck - Runs static code analyzer staticcheck - currently broken
staticcheck: 	
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...

## seccheck - Code vulnerability check
seccheck:	
	govulncheck ./...

## lint - format code and tidy modules
lint:
	go fmt ./...
	go mod tidy -v