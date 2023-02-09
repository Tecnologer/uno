formatVersion=+%y%m%d
version=`git describe --always --tags`

test:
	go test -coverprofile=c.out 
	go tool cover -html=c.out -o coverage.html

build:
	go build -ldflags "-X main.minversion=`date -u $(formatVersion)` -X main.version=$(version)" -buildmode=plugin -o ../../plugins/classic.so ./...