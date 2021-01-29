package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	en := rlp.Encode()
	fmt.Println("Hello, world.")
}
