package server

import (
	"fmt"

	"github.com/andreylm/bmlabs/pkg/logger"
	"github.com/andreylm/bmlabs/pkg/router"
	"github.com/valyala/fasthttp"
)

type server struct {
	port   int
	host   string
	router *router.Router
}

func NewServer(host string, port int) *server {
	return &server{
		port:   port,
		host:   host,
		router: router.NewRouter(),
	}
}

func (s *server) Init() error {
	s.router.Init()
	return nil
}

func (s *server) Run() error {
	logger.Get().Infof("Starting server on %s:%d", s.host, s.port)
	return fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), s.router.HandleRequest)
}
