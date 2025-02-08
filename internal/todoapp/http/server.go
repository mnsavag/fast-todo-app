package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/config"
	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config     config.Config
	logger     *logrus.Logger
	httpServer *http.Server
}

func NewServer(config config.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Start() error {
	go func() {
		s.logger.Infof("http server started on port, :%s", s.config.HttpPort)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("start http server failed: %s", err.Error())
		}
	}()

	return nil
}

func (s *Server) Register(ctx context.Context) error {
	if err := s.newHttpServer(ctx); err != nil {
		s.logger.Fatalf("cant register http server: %s", err.Error())
		return err
	}
	return nil
}

func (s *Server) newHttpServer(ctx context.Context) error {
	handler := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := v1.RegisterTodoAppServiceV1HandlerFromEndpoint(
		ctx,
		handler,
		net.JoinHostPort(s.config.Host, s.config.GrpcPort),
		opts,
	)
	if err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Addr:         net.JoinHostPort(s.config.Host, s.config.HttpPort),
		Handler:      handler,
		IdleTimeout:  90 * time.Second,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	return nil
}

func (s *Server) Stop() error {
	graceCtx := context.Background()

	if err := s.httpServer.Shutdown(graceCtx); err != nil {
		return errors.Wrap(err, "could not gracefully shutdown http server")
	}

	s.logger.Info("http servers stopped...")

	return nil
}
