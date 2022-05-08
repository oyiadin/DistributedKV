package internal

import "strconv"

type Operator func(*DB, []string) *Response

var name2operator = map[string]Operator{
	"get": OpGet,
	"set": OpSet,
}

func OpGet(db *DB, arguments []string) (resp *Response) {
	resp = checkArgumentLength(arguments, 1)
	if resp != nil {
		return resp
	}

	key := arguments[0]
	value, err := db.Storage.InMem.Get(key)
	if err != nil {
		return MakeResponseOfFailure(err.Error())
	}
	return MakeResponseWithOneResult(value)
}

func OpSet(db *DB, arguments []string) (resp *Response) {
	resp = checkArgumentLength(arguments, 2)
	if resp != nil {
		return resp
	}

	key, value := arguments[0], arguments[1]
	err := db.Storage.InMem.Set(key, value)
	if err != nil {
		return MakeResponseOfFailure(err.Error())
	}
	return MakeResponseWithOneResult(value)
}

func checkArgumentLength(arguments []string, expectedLength int) *Response {
	if len(arguments) != expectedLength {
		return MakeResponseOfFailure("expected argument length: " + strconv.Itoa(expectedLength))
	}
	return nil
}
