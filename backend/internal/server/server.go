package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/progate-hackathon-ari/backend/pkg/log"
	"golang.org/x/sync/errgroup"
)

const (
	DefaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	// srv server
	srv *http.Server
	// shutdown timeout
	shutdownTimeout time.Duration
	// logger
	// // worker service
	// ws *workerservice.WorkerService
}

// New はHTTPサーバを生成する
func New(addr string, handler http.Handler, opts ...Option) *Server {
	s := &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		shutdownTimeout: DefaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Run はHTTPサーバを起動する。
func (s *Server) Run(ctx context.Context) error {
	log.Info(ctx, "server listening at ...", "address", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error(ctx, "server failed to start", "error", err)
		return err
	}

	return nil
}

// Shutdown はhttp serverを停止する
func (s *Server) Shutdown(ctx context.Context) error {
	log.Error(ctx, "server shutting down ...")
	return s.srv.Shutdown(ctx)
}

// RunWithGraceful はサーバの起動とInterrupt,SIGTERMによる停止信号に対してのGracefulShutdownを提供する
func (s *Server) RunWithGraceful() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return s.Run(ctx)
	})

	group.Go(func() error {
		<-gCtx.Done()

		ctx, cancel = context.WithTimeout(context.TODO(), s.shutdownTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	})

	if err := group.Wait(); err != nil && err != context.Canceled {
		log.Error(ctx, "server shutdown failed, Error", "error", err)
		return err
	}

	log.Info(ctx, "server shutdown successfully")
	return nil
}
