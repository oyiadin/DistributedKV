package main

import (
	"github.com/oyiadin/DistributedKV/internal"
)

func main() {
	var err error

	server, err := internal.NewServer()
	if err != nil {
		panic(err)
	}

	err = server.Init()
	if err != nil {
		panic(err)
	}

	panic(server.ListenAndServe())
}
