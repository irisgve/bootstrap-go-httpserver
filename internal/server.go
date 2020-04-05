package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	Logger *zap.Logger
	Config *Config

	*http.Server
}

func NewServer(cfg *Config) (*Server, error) {
	logger, err := zap.NewProduction(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		return nil, fmt.Errorf("can't initialize logger: %w", err)
	}
	defer logger.Sync()

	errorLogger, err := zap.NewStdLogAt(logger, zap.ErrorLevel)
	if err != nil {
		return nil, fmt.Errorf("can't initialize error logger for router: %w", err)
	}

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      newRouteHandler(logger),
		ErrorLog:     errorLogger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{logger, cfg, &srv}, nil
}

func (s *Server) Start() {
	defer s.Logger.Sync()

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Fatal("cannot listen on", zap.String("addr", s.Addr), zap.Error(err))
		}
	}()

	s.Logger.Info("server is ready to handle requests", zap.String("addr", s.Addr))
	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	s.Logger.Info("server is shutting down", zap.String("reason", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		s.Logger.Fatal("cannot gracefuly shutdown the server", zap.Error(err))
	}

	s.Logger.Info("server stopped")
}
