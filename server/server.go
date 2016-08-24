package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port string
	path string
}

func NewServer(path_ string, port_ string) *Server {
	port_ = ":" + port_
	log.Println(port_)
	return &Server{
		port: port_,
		path: path_,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("request...")
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func (s *Server) Run() {
	http.HandleFunc(s.path, handler)
	http.ListenAndServe(s.port, nil)
}
