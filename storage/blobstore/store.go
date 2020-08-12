package blobstore

import "github.com/freggy/wo/storage/blobstore/driver"

// Store is the interface for the blob storage layer abstracting away concrete storage
// implementation provided by specific drivers.
type Store struct {
	driver driver.Driver
}

// GetMap retrieves a map identified by the given digest from the underlying storage.
func (store *Store) GetMap(digest string) ([]byte, error) {}

// GetMap retrieves a configuration identified by the given digest from the underlying storage.
func (store *Store) GetConfig(digest string) (string, error) {}

// GetTag retrieves a tag for a specific map inside a repository form the underlying storage.
func (store *Store) GetTag(repository, mapName, tagName string) (Tag, error) {}

// PutMap saves a map to the underlying storage making it identifiable by its content.
func (store *Store) PutMap(mapblob []byte) error {}

// PutConfig saves a configuration to the underlying storage making it identifiable by its content.
func (store *Store) PutConfig(content string) error {}

// PutTag saves a tag for a map inside a repository to the underlying storage.
func (store *Store) PutTag(repository, mapName, tagName string, tag Tag) error {}