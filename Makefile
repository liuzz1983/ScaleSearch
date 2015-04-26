
TEST=./...
BENCH=.


default: build

bench:
    go test -v -test.run=NOTHINCONTAINSTHIS -test.bench=$(BENCH)

get:
	@go get -d ./...

build: get 
	@mkdir -p bin
	@go build -ldflags=$(GOLDFLAGS) -a -o bin/search ./server/

fmt:
	@go fmt ./...

test: fmt
	@go get github.com/stretchr/testify/assert
	@echo "=== TESTS ==="
	@go test -v -cover ./...