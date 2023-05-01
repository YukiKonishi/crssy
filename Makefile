PACKAGE_LIST := $(shell go list ./...)
crssy:
	go build -o crssy $(PACKAGE_LIST)
test:
	go test $(PACKAGE_LIST)
clean:
	rm -f crssy
