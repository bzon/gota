BINARY = gota
VET_REPORT = vet.report
TEST_REPORT = tests.xml
GOARCH = amd64
BINDIR = bin

VERSION?=latest
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=bzon
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean getdep linux darwin windows
	cd ${BINDIR} && shasum -a 256 ** > shasum256.txt

create_bin_dir:
	rm -fr ${BINDIR}
	mkdir -p ${BINDIR}
	
linux: getdep create_bin_dir
	echo Build for linux ${GOARCH}
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINDIR}/${BINARY}-linux-${GOARCH} . 

darwin: getdep create_bin_dir
	echo Build for darwin ${GOARCH}
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINDIR}/${BINARY}-darwin-${GOARCH} .

windows: create_bin_dir
	echo Build for windows ${GOARCH}
	GOOS=windows go get -v -u github.com/spf13/cobra
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINDIR}/${BINARY}-windows-${GOARCH}.exe . 

test: getdep
	echo Performing a go test
	go test ./... -v -race

coverage: getdep
	echo Performing test with coverage
	go test ./... -v -race -coverprofile=coverage.txt -covermode=atomic

getdep:
	go get -v ./...

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VET_REPORT}
	-rm -fr ${BINDIR}
