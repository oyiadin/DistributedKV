package main

import (
	"fmt"
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

	resp := server.ExecuteOne("get k")
	fmt.Println(resp)
	resp = server.ExecuteOne("ge22t k")
	fmt.Println(resp)
	resp = server.ExecuteOne("set k v")
	fmt.Println(resp)
	resp = server.ExecuteOne("get k")
	fmt.Println(resp)

	//panic(server.ListenAndServe())
}
