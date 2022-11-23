package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	// address := pubKey.Address()

	msg := []byte("Hello, world!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	sig.Verify(pubKey, msg)

	assert.True(t, sig.Verify(pubKey, msg))

	fmt.Println(sig)
}
