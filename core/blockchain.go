package core

import (
	"fmt"
	"sync"

	"github.com/stefanlester/modularblockchain/crypto"
	"github.com/stefanlester/modularblockchain/types"
	"github.com/go-kit/log"
)

type Blockchain struct {
	logger log.Logger
	store  Storage
	// TODO: double check this!
	lock       sync.RWMutex
	headers    []*Header
	blocks     []*Block
	txStore    map[types.Hash]*Transaction
	blockStore map[types.Hash]*Block

	accountState *AccountState

	stateLock       sync.RWMutex
	collectionState map[types.Hash]*CollectionTx
	mintState       map[types.Hash]*MintTx
	validator       Validator

	// TODO: make this an interface.
	contractState *State
}

func NewBlockchain(l log.Logger, genesis *Block) (*Blockchain, error) {
	// All states inside the scope of the newblockchain.

	// TODO: read this from disk later on
	accountState := NewAccountState()

	coinbase := crypto.PublicKey{}
	accountState.CreateAccount(coinbase.Address())

	bc := &Blockchain{
		contractState:   NewState(),
		headers:         []*Header{},
		store:           NewMemorystore(),
		logger:          l,
		accountState:    accountState,
		collectionState: make(map[types.Hash]*CollectionTx),
		mintState:       make(map[types.Hash]*MintTx),
		blockStore:      make(map[types.Hash]*Block),
		txStore:         make(map[types.Hash]*Transaction),
	}
	
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}