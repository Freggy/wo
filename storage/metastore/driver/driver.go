package driver

type Driver interface {
	Get(repository, mapName, key string) (interface{}, error)

	Set(repository, mapName, key string, value interface{}) error
}
