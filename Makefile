# Binary name
BINARY=jzerp

# Builds the project
build:
	go build -o ${BINARY}
	go test -v

# Installs our project: copies binaries
install:
	go install
release:
	# Clean
	go clean

	# Build for mac
	go build -o ${BINARY}_darwin_arm64

	# Build for linux
	go clean
	rm -rf ${BINARY}_linux_386
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o  ${BINARY}_linux_386

	# Build for windows
	go clean
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${BINARY}_windows_386

	go clean

# Cleans our projects: deletes binaries
clean:
	go clean

.PHONY:  clean build