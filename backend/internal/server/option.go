package server

import (
	"time"
)

type Option func(s *Server)

func WithShutdownTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = t
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.srv.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.srv.WriteTimeout = t
	}
}

// ...必要に応じて追加していく...
