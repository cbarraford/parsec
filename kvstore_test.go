package parsec

import (
	ctype "github.com/cosmos/cosmos-sdk/types"

	. "gopkg.in/check.v1"
)

type KVSuite struct{}

var _ = Suite(&KVSuite{})

func (s *KVSuite) TestKVStore(c *C) {
	key := []byte("foo")
	value := []byte("bar")
	storeKey := ctype.NewKVStoreKey("test")

	store := MockContext{}.KVStore(storeKey)
	store.Set(key, value)

	c.Assert(store.Has(key), Equals, true)
	c.Assert(store.Has(value), Equals, false)
	c.Assert(store.Get(key), DeepEquals, value)

	store.Delete(key)
	c.Assert(store.Has(key), Equals, false)

	// check that store persists
	store.Set(key, value)
	store = MockContext{}.KVStore(storeKey)
	c.Assert(store.Get(key), DeepEquals, value)
}
