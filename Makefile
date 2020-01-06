.PHONY: all development-build production-build test clean

all:
	make development-build && ./sql-judge-api

development-build:
	export GIN_MODE=debug && go build

production-build:
	export GIN_MODE=release && go build

test:
	go test ./...

clean:
	rm sql-judge-api