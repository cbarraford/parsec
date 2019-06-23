package parsec

type Store interface {
	// Get returns nil iff key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)

	// Delete deletes the key. Panics on nil key.
	Delete(key []byte)

	// Iterator over a domain of keys in ascending order. End is exclusive.
	// Start must be less than end, or the Iterator is invalid.
	// Iterator must be closed by caller.
	// To iterate over entire domain, use store.Iterator(nil, nil)
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// Exceptionally allowed for cachekv.Store, safe to write in the modules.
	// Iterator(start, end []byte) Iterator

	// Iterator over a domain of keys in descending order. End is exclusive.
	// Start must be less than end, or the Iterator is invalid.
	// Iterator must be closed by caller.
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// Exceptionally allowed for cachekv.Store, safe to write in the modules.
	// ReverseIterator(start, end []byte) Iterator
}

type KVStore struct {
	Store
	data map[string][]byte
}

func NewKVStore() KVStore {
	return KVStore{
		data: make(map[string][]byte),
	}
}

func (k KVStore) Has(key []byte) bool {
	_, ok := k.data[string(key)]
	return ok
}

func (k KVStore) Get(key []byte) []byte {
	if !k.Has(key) {
		return nil
	}
	return k.data[string(key)]
}

func (k KVStore) Set(key, value []byte) {
	k.data[string(key)] = value
}

func (k KVStore) Delete(key []byte) {
	delete(k.data, string(key))
}
