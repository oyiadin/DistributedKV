package internal

import "strings"

type Server struct {
	DB DB
}

func NewServer() (*Server, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	return &Server{
		DB: *db,
	}, nil
}

func (s *Server) Init() (err error) {
	err = s.DB.Init()
	return
}

func (s *Server) ListenAndServe() {

}

func (s *Server) ExecuteOne(statement string) *Response {
	parts := strings.Split(statement, " ")
	return s.DB.ExecuteOne(parts)
}
