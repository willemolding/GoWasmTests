package main 

import (
	"fmt"
	schnorrkel "github.com/ChainSafe/go-schnorrkel"
)

func main() {
	msg := []byte("hello friends")
	signingCtx := []byte("example")

	signingTranscript := schnorrkel.NewSigningContext(signingCtx, msg)
	verifyTranscript := schnorrkel.NewSigningContext(signingCtx, msg)

	priv, pub, err := schnorrkel.GenerateKeypair()
	if err != nil {
		fmt.Println(err)
		return
	}

	sig, err := priv.Sign(signingTranscript)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok := pub.Verify(sig, verifyTranscript)
	if !ok {
		fmt.Println("did not verify :(")
		return
	}
}