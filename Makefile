BINARY_NAME=gitlab-teams-bot
BUILD_LDFLAGS=-ldflags="-w -s"

build: build-linux build-darwin

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${BUILD_LDFLAGS} -o ${BINARY_NAME}-linux *.go
	chmod 755 ${BINARY_NAME}-linux

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${BUILD_LDFLAGS} -o ${BINARY_NAME}-darwin *.go
	chmod 755 ${BINARY_NAME}-darwin

optimize:
	if [ -x /usr/local/bin/upx ]; then upx --brute ${BINARY_NAME}-*; fi

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux

