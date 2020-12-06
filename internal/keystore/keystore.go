package keystore

import (
	"errors"

	"github.com/devguyio/kvstore/internal/logging"
)

// KeyStore is a struct representing an in-memory key-value store
type KeyStore struct {
	data map[string]string
}

// NewKeyStore is a factory function to initialize the KeyStore
func NewKeyStore() (k *KeyStore) {
	return &KeyStore{
		data: make(map[string]string),
	}
}

// Put adds or update key/value pair
func (ks *KeyStore) Put(k string, v string) error {
	ks.data[k] = v
	return nil
}

// Get retrieves a value for key k or "" if it doesn't exist
func (ks *KeyStore) Get(k string) string {
	return ks.data[k]
}

// Del deletes a value for key k
func (ks *KeyStore) Del(k string) error {
	if _, ok := ks.data[k]; !ok {
		return errors.New("Invalid key")
	}
	logging.DefaultLogger().Debug.Printf("deleting key %s with value %s", k, ks.data[k])
	delete(ks.data, k)
	return nil
}
