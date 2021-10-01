package server

import (
	"context"
	"net/http"
	"time"

	"github.com/covrom/hex_arch_example/best_practice/internal/app/repos/user"
	"github.com/covrom/hex_arch_example/best_practice/internal/app/starter"
)

var _ starter.APIServer = &Server{}

type Server struct {
	srv http.Server
	us  *user.Users
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start(us *user.Users) {
	s.us = us
	go s.srv.ListenAndServe()
}
