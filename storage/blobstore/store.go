package blobstore

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/freggy/wo/storage/blobstore/driver"
)

// Store is the interface for the blob storage layer abstracting away concrete storage
// implementation provided by specific drivers.
type Store struct {
	basePath string
	driver   driver.Driver
}

var (
	ErrBlobAlreadyExists = errors.New("blob already exists")
	ErrBlobNotFound      = errors.New("blob not found")

	mapPath    = "%s/maps/%s"
	configPath = "%s/configs/%s"
	tagPath    = "%s/tags/%s/%s/%s"
)

// GetMap retrieves a map identified by the given digest from the underlying storage.
func (store *Store) GetMap(digest string) ([]byte, error) {
	path := fmt.Sprintf(mapPath, store.basePath, digest)

	blob, err := store.get(path)
	if err != nil {
		return nil, err
	}

	return blob, nil
}

// GetMap retrieves a configuration identified by the given digest from the underlying storage.
func (store *Store) GetConfig(digest string) (string, error) {
	path := fmt.Sprintf(configPath, store.basePath, digest)

	blob, err := store.get(path)
	if err != nil {
		return "", err
	}

	return string(blob), nil
}

// GetTag retrieves a tag for a specific map inside a repository form the underlying storage.
func (store *Store) GetTag(repository, mapName, tagName string) (Tag, error) {
	path := fmt.Sprintf(tagPath, store.basePath, repository, mapName, tagName)

	blob, err := store.get(path)
	if err != nil {
		return Tag{}, err
	}

	var tag Tag
	if err := json.Unmarshal(blob, &tag); err != nil {
		return Tag{}, err
	}

	return tag, nil
}

// PutMap saves a map to the underlying storage making it identifiable by its content.
func (store *Store) PutMap(blob []byte) error {
	digest := sha256digest(blob)
	path := fmt.Sprintf(mapPath, store.basePath, digest)
	if err := store.put(path, blob); err != nil {
		return err
	}
	return nil
}

// PutConfig saves a configuration to the underlying storage making it identifiable by its content.
func (store *Store) PutConfig(content string) error {
	blob := []byte(content)
	digest := sha256digest(blob)
	path := fmt.Sprintf(configPath, store.basePath, digest)
	if err := store.put(path, blob); err != nil {
		return err
	}
	return nil
}

// PutTag saves a tag for a map inside a repository to the underlying storage.
func (store *Store) PutTag(repository, mapName, tagName string, tag Tag) error {
	path := fmt.Sprintf(tagPath, store.basePath, repository, mapName, tagName)

	blob, err := json.Marshal(tag)
	if err != nil {
		return err
	}

	if err := store.put(path, blob); err != nil {
		return err
	}

	return nil
}

func (store *Store) get(path string) ([]byte, error) {
	ok, err := store.driver.Exists(path)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrBlobNotFound
	}

	blob, err := store.driver.Read(path)
	if err != nil {
		return nil, err
	}

	return blob, nil
}

func (store *Store) put(path string, blob []byte) error {
	ok, err := store.driver.Exists(path)
	if err != nil {
		return err
	}

	if ok {
		return ErrBlobAlreadyExists
	}

	if err := store.driver.Write(path, blob); err != nil {
		return err
	}

	return nil
}
