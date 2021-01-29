
build_go:
	GOOS=js GOARCH=wasm go build -o build/go.wasm main.go

build_tiny:
	tinygo build -o build/tiny.wasm -target=wasm main.go

run:
	go run main.go