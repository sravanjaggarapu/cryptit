package main

import (
	"fmt"

	"github.com/sravanjaggarapu/cryptit.git/encrypt"

	"github.com/sravanjaggarapu/cryptit.git/decrypt"
)

func main() {
	str := encrypt.Encrypting("Sravan Kumar")
	fmt.Println(str)
	decrypted := decrypt.Decrypting(str)
	fmt.Println(decrypted)

}
