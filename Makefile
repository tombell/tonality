MODFLAGS=-mod=vendor
TESTFLAGS=-cover

test:
	go test ${MODFLAGS} ${TESTFLAGS} ./...

testmutation:
	go test ${MODFLAGS} ${TESTFLAGS} -tags=mutation ./...

.DEFAULT_GOAL := test
.PHONY: test testmutation
