package main

import (
	"fmt"
	"io"
	rlp "github.com/ethereum/go-ethereum/rlp"
)

type MyCoolType struct {
	Name string
	a, b uint
}

// EncodeRLP writes x as RLP list [a, b] that omits the Name field.
func (x *MyCoolType) EncodeRLP(w io.Writer) (err error) {
	return rlp.Encode(w, []uint{x.a, x.b})
}

func main() {
	var t = &MyCoolType{Name: "foobar", a: 5, b: 6}
	var bytes, _ = rlp.EncodeToBytes(t)
	fmt.Printf("%v → %X\n", t, bytes)
}
