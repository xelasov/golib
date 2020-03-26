
build:
	go build ./pkg/...

test:
	go test -v ./pkg/...

build-examples:
	go build -o ./examples/conf/target/conf-example ./examples/conf/conf-example.go
