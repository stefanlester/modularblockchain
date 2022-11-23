package crypto

import (
	"fmt"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()

	fmt.Println(address)
}
