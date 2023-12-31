package eventindexer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Block is a database model representing simple header types
// to keep track of our most recently processed block number and hash.
type Block struct {
	ID      int    `json:"id"`
	Height  uint64 `json:"blockHeight" gorm:"column:block_height"`
	Hash    string `json:"hash"`
	ChainID int64  `json:"chainID"`
}

// SaveBlockOpts is required to store a new block
type SaveBlockOpts struct {
	Height  uint64
	Hash    common.Hash
	ChainID *big.Int
}

// BlockRepository defines methods necessary for interacting with
// the block store.
type BlockRepository interface {
	Save(opts SaveBlockOpts) error
	GetLatestBlockProcessed(chainID *big.Int) (*Block, error)
}
