package main

import (
	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}

	for {
		line, err := rl.Readline()
		if err != nil {
			panic(err)
		}
		server.ExecuteOne(line)
	}
}
