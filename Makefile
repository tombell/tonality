MODFLAGS=-mod=vendor
TESTFLAGS=-cover

all: test

test:
	go test ${MODFLAGS} ${TESTFLAGS} ./...

testmutation:
	go test ${TESTFLAGS} -v -tags=mutation ./...

.PHONY: all test testv
