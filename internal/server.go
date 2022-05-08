package internal

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	DB          DB
	initialized bool
}

type H map[string]interface{}

func NewServer() (*Server, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	return &Server{
		DB: *db,
	}, nil
}

type handler func(c *handlerContext, w http.ResponseWriter, r *http.Request) (*H, error)
type middleware handler

type handlerContext struct {
	server      *Server
	handler     handler
	index       int
	middlewares []middleware
}

func (s *Server) Init() (err error) {
	if s.initialized {
		panic("server already initialized")
	}

	err = s.DB.Init()
	if err != nil {
		return
	}

	// init http
	func() {
		http.HandleFunc("/command", handlerFactory(&handlerContext{
			handler: handleCommand,
			server:  s,
			middlewares: []middleware{
				checkMethodMiddleware("POST"),
			},
		}))
	}()

	s.initialized = true
	return
}

func (c *handlerContext) Run(w http.ResponseWriter, r *http.Request) (*H, error) {
	return c.Next(w, r)
}

func (c *handlerContext) Next(w http.ResponseWriter, r *http.Request) (*H, error) {
	if c.index >= len(c.middlewares) {
		// execute the handler if we have no more handlerContext
		return c.handler(c, w, r)
	}
	mw := c.middlewares[c.index]
	c.index++
	return mw(c, w, r)
}

func handlerFactory(c *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var results []byte
		body, err := c.Run(w, r)
		if err != nil {
			if body != nil {
				log.Println("warning: err is not nil while body is not nil: ", body)
			}
			w.WriteHeader(http.StatusBadRequest)
			results = []byte(err.Error())
		} else {
			marshalled, err2 := json.MarshalIndent(body, "", " ")
			if err2 != nil {
				results = []byte("error occurred when marshalling results")
			} else {
				results = marshalled
			}
		}

		bytesWritten, err2 := w.Write(results)
		if err2 != nil {
			log.Println(
				"failed to write response, bytesWritten: ", bytesWritten,
				", err: ", err2, ", when dealing with original err: ", err)
		}
	}
}

func handleCommand(c *handlerContext, _ http.ResponseWriter, r *http.Request) (*H, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodyString := strings.TrimSpace(string(bodyBytes))
	resp, err := c.server.ExecuteOne(bodyString).ToMap()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func checkMethodMiddleware(methods ...string) middleware {
	return func(c *handlerContext, w http.ResponseWriter, r *http.Request) (*H, error) {
		targetMethod := r.Method
		for _, method := range methods {
			if targetMethod == method {
				return c.Next(w, r)
			}
		}
		return nil, errors.New("unsupported method")
	}
}

func (s *Server) ListenAndServe() error {
	log.Printf("listening on %s (%s)\n", s.DB.Config.peerSelf.URL, s.DB.Config.peerSelf.Name)
	return http.ListenAndServe(s.DB.Config.peerSelf.URL, nil)
}

func (s *Server) ExecuteOne(statement string) *Response {
	log.Printf("statement: %s\n", statement)
	parts := strings.Split(statement, " ")
	return s.DB.ExecuteOne(parts)
}
