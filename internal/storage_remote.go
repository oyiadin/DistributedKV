package internal

type RemoteStorage struct {
	Raft Raft
}

func NewRemoteStorage() *RemoteStorage {
	return &RemoteStorage{}
}

func (s *RemoteStorage) Init() {

}

//func (s *RemoteStorage) Set(key, value string) (err error) {
//
//}
