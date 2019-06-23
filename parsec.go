package parsec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var kvstore = KVStore{
	data: make(map[string][]byte),
}

type Context interface {
	KVStore(key sdk.StoreKey) KVStore
}

type MockContext struct {
	Context
}

func (c MockContext) KVStore(key sdk.StoreKey) KVStore {
	return kvstore
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
