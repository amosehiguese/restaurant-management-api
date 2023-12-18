.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: test
	go build -o bin/restuarant

.PHONY:start
start: build
	./bin/restuarant