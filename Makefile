all: fmt build test

build: ui-build 
	go build

ui-build:
	cd utsukushii_ui && npm run build && cd ..

unit:
	go test ./model -v -coverprofile=coverage

fmt:
	go fmt
