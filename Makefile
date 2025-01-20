.PHONY:
.SILENT:
.DEFAULT_GOAL := run

test:
	go test --count=1  -bench=. -v ./...

lint:
	golangci-lint run --fix --verbose