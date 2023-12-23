go:
	go run ./cmd/go/main.go

test:
	go test -v -race -covermode=atomic ./...

lint: vet fmt optimise
# The disabled linters are deprecated
	golangci-lint run --enable-all --sort-results --tests --fix \
		--disable maligned --disable interfacer --disable scopelint --disable golint --disable exhaustivestruct \
		--disable varcheck --disable ifshort --disable nosnakecase --disable structcheck --disable deadcode \
		--disable forbidigo --disable depguard --disable ireturn --disable goerr113 ./...

vet:
	go vet -all ./...

fmt:
	go fmt ./cmd/*
	go fmt ./internal/*
	gofumpt -l -w .

optimise: structs

structs:
	structslop -fix -apply ./...

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/orijtech/structslop/cmd/structslop@latest
	go install github.com/mgechev/revive@latest