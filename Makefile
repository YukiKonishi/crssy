PACKAGE_LIST := $(shell go list ./...)
crssy: test
	go build -o crssy $(PACKAGE_LIST)
test:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
clean:
	rm -f crssy
