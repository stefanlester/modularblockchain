package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	// address := pubKey.Address()

	msg := []byte("Hello, world!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

// func TestKeyPairSignVerifyFail(t *testing.T) {
// 	privKey := GeneratePrivateKey()
// 	pubKey := privKey.PublicKey()
// 	// address := pubKey.Address()

// 	msg := []byte("Hello, world!")
// 	sig, err := privKey.Sign(msg)
// 	assert.Nil(t, err)

// 	assert.True(t, sig.Verify(pubKey, msg))

// }
