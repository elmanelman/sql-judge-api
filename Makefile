dev:
	export GIN_MODE=debug && go run *.go

build:
	export GIN_MODE=release && go build *.go