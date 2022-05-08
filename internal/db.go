package internal

type dbStorage struct {
	InMem  InMemoryStorage
	Local  LocalStorage
	Remote RemoteStorage
}

type DB struct {
	Config  Config
	Storage dbStorage
}

func NewDB() (*DB, error) {
	db := &DB{
		Config: Config{},
		Storage: dbStorage{
			InMem:  *NewInMemoryStorage(),
			Local:  LocalStorage{},
			Remote: RemoteStorage{},
		},
	}

	err := db.Init()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Init() (err error) {
	err = db.Config.LoadFrom([]string{"./", "./etc/config"})
	if err != nil {
		return err
	}
	db.Storage.InMem.Init()
	db.Storage.Local.Init()
	db.Storage.Remote.Init()
	return nil
}

func (db *DB) ExecuteOne(statement []string) *Response {
	command := statement[0]
	return db.DispatchCommand(command, statement[1:])
}

func (db *DB) DispatchCommand(command string, arguments []string) *Response {
	op, ok := name2operator[command]
	if !ok {
		return MakeResponseOfFailure("unknown command '" + command + "'")
	}

	return op(db, arguments)
}
