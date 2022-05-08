package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	address = flag.String("address", "http://127.0.0.1:4233", "the address to connect to")
)

func main() {
	flag.Parse()

	_, err := url.Parse(*address)
	if err != nil {
		panic(err)
	}

	// ping server
	func() {
		_, err := http.Get(*address)
		if err != nil {
			panic(fmt.Sprintf("error connecting to server: %v", err))
		}
	}()

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}

	for {
		line, err := rl.Readline()
		if err == io.EOF || err == readline.ErrInterrupt {
			fmt.Println("goodbye!")
			os.Exit(0)
		}
		if err != nil {
			panic(err)
		}

		u, _ := url.Parse(*address)
		u.Path = path.Join(u.Path, "/command")
		resp, err := http.Post(
			u.String(),
			"text/plain",
			bytes.NewReader([]byte(line)))
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
			continue
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
