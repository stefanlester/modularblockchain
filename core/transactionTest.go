package core

import (
	"testing"

	"github.com/stefanlester/modularblockchain/crypto"
	"github.com/stefanlester/modularblockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTransactionWithTamper(t *testing.T) {
	tx := NewTransaction(nil)

	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()
	hackerPrivKey := crypto.GeneratePrivateKey()

	tx.From = fromPrivKey.PublicKey()
	tx.To = toPrivKey.PublicKey()
	tx.Value = 666

	assert.Nil(t, tx.Sign(fromPrivKey))
	tx.hash = types.Hash{}

	tx.To = hackerPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}
