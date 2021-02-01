
build_go:
	GOOS=js GOARCH=wasm go build -o build/$(TEST)/go.wasm $(TEST).go

build_tiny:
	tinygo build -o build/$(TEST)/tiny.wasm -target=wasm $(TEST).go