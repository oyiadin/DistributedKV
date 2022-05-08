package internal

type Storage interface{}

type ReadStorage interface {
	Storage
	Get(key string) (value string, err error)
}

type WriteStorage interface {
	Storage
	Set(key, value string) (err error)
}

type ReadWriteStorage interface {
	ReadStorage
	WriteStorage
}
