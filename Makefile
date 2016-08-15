SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

# Name for the binary output
BINARY=utilimath

# Values passed for Version and BuildTime
VERSION=0.1
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build, interpolate the variables
LDFLAGS=-ldflags "-X github.com/oldspiceland/utilimath/utilimath.Version=${VERSION} -X github.com/oldspiceland/utilimath/utilimath.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	go build ${LDFLAGS} -o ${BINARY} utilimath.go


.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY]; fi
