SRC := $(wildcard goasobi/*.go pkg/*.go)
PORT ?= 8080

bin/hello: $(SRC)
	go build -o bin/hello ./goasobi

.PHONY: run/hello
run/hello: bin/hello
	./bin/hello -p $(PORT)

.PHONY: fmt/pkg
fmt/pkg:
	go fmt ./pkg/...

.PHONY: test/pkg
test/pkg:
	go test ./pkg/...

.PHONY: cov/pkg
cov/pkg:
	mkdir -p bin
	go test -coverprofile=bin/pkg.cover ./pkg/...
	go tool cover -func=bin/pkg.cover

.PHONY: clean
clean:
	rm -f bin/hello bin/pkg.cover
