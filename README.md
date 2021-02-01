# Go WASM Tests

A few simple builds to compare compatibility of Go and TinyGo.

## Usage

```shell
TEST=<test> make build_go
TEST=<test> make build_tiny
```

where `<test-file>` is one of `rlp`, `scale`, `schnorr` or `sha`.

