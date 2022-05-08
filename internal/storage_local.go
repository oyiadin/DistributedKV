package internal

type LogType int

const (
	RawLog LogType = iota
	CompactedLog
)

type LogFile struct {
	Type     LogType
	Path     string
	Position int
	Capacity int
}

type LocalStorage struct {
	LogFiles             []*LogFile
	CurrentActiveLogFile *LogFile
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) Init() {

}

func (s *LocalStorage) Get(key string) (value string, err error) {
	return "", nil
}
