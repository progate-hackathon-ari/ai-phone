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

// func WithWorkerService(ws *worker.WorkerService) Option {
// 	return func(s *Server) {
// 		s.ws = ws
// 	}
// }

// ...必要に応じて追加していく...
