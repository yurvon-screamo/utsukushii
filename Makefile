all: fmt build unit

build: ui-build 
	go build

ui-build:
	cd utsukushii_ui && npm run build && cd ..

unit:
	go test ./junit_input ./model -v -coverprofile=coverage

fmt:
	go fmt
