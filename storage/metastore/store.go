package metastore

import "github.com/freggy/wo/storage/metastore/driver"

// Store is the interface between the metadata storage layer abstracting away
// concrete storage implementation provided by specific drivers.
type Store struct {
	driver driver.Driver
}

func (store *Store) Get(repository, mapName, key string) (interface{}, error) {
	return store.driver.Get(repository, mapName, key)
}

func (store *Store) Set(repository, mapName, key string, value interface{}) error {
	return store.driver.Set(repository, mapName, key, value)
}
