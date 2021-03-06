BINARY_NAME=tix
BUILD_FOLDER=build/release
LINUX_64_OUTPUT=$(BUILD_FOLDER)/linux64/$(BINARY_NAME)
LOCAL_OUTPUT=$(BUILD_FOLDER)/$(BINARY_NAME)
MAC_64_OUTPUT=$(BUILD_FOLDER)/mac64/$(BINARY_NAME)
WINDOWS_64_OUTPUT=$(BUILD_FOLDER)/windows64/$(BINARY_NAME).exe

GO_BUILD=go build -ldflags "-X main.version=${VERSION}" -o

all: deps test build-linux build-mac build-windows tar
build:
	$(GO_BUILD) $(LOCAL_OUTPUT)
clean:
	go clean
	rm -rf build
deps:
	go get -v -t -d ./...
build-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) $(LINUX_64_OUTPUT)
build-mac:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) $(MAC_64_OUTPUT)
build-windows:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) $(WINDOWS_64_OUTPUT)
tar:
	cd build && tar -zcvf release.tar.gz release
test:
	go test -v ./...
