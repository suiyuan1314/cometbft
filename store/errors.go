package store

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeHeight      = errors.New("height must be greater than 0")
	ErrBlockPartIncomplete = errors.New("BlockStore can only save complete block part sets")
	ErrNilBlcokmeta        = errors.New("nil blockmeta")
)

type ErrExceedLatestHeight struct {
	Height int64
}

func (e ErrExceedLatestHeight) Error() string {
	return fmt.Sprintf("cannot prune beyond the latest height %v", e.Height)
}

type ErrExceedBaseHeight struct {
	Height int64
	Base   int64
}

func (e ErrExceedBaseHeight) Error() string {
	return fmt.Sprintf("cannot prune to height %v, it is lower than base height %v", e.Height, e.Base)
}

type ErrNonContiguousBlocks struct {
	Expected int64
	Actual   int64
}

func (e ErrNonContiguousBlocks) Error() string {
	return fmt.Sprintf("BlockStore can only save contiguous blocks. Wanted %v, got %v", e.Expected, e.Actual)
}

type ErrCommitHeightMismatch struct {
	Expected int64
	Actual   int64
}

func (e ErrCommitHeightMismatch) Error() string {
	return fmt.Sprintf("BlockStore cannot save seen commit of a different height (block: %d, commit: %d)", e.Actual, e.Expected)
}

type ErrMarshalCommit struct {
	Err error
}

func (e ErrMarshalCommit) Error() string {
	return fmt.Sprintf("unable to marshal commit: %v", e.Err)
}

func (e ErrMarshalCommit) Unwrap() error {
	return e.Err
}

type ErrDBOpt struct {
	Err error
}

func (e ErrDBOpt) Error() string {
	return e.Err.Error()
}

func (e ErrDBOpt) Unwrap() error {
	return e.Err
}
