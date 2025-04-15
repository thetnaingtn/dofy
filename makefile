build:
	go build -o bin/dofy ./

goreleaser-check:
	goreleaser check

goreleaser-build:
	goreleaser release --skip=publish --snapshot --clean

