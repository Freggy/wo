package driver

type filesystemDriver struct {

}

func (fd *filesystemDriver) Get(repository, mapName, key string) (interface{}, error) {
	return nil, nil
}

func (fd *filesystemDriver) Set(repository, mapName, key string, value interface{}) error {
	return nil
}