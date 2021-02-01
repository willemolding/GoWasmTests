package main

import (
	"fmt"
	"bytes"
	scale "github.com/ChainSafe/gossamer/lib/scale"
)

func main() {
	buf := bytes.Buffer{}
	sd := scale.Decoder{Reader: &buf}
	buf.Write([]byte{0xff})
	output, _ := sd.ReadByte()
	fmt.Println(output)
}
