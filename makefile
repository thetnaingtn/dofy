build:
	go build -o bin/dofy ./

goreleaser-check:
	goreleaser check

goreleaser-build:
	goreleaser build --snapshot --clean

goreleaser-dry-release:
	goreleaser release --skip=publish --snapshot --clean