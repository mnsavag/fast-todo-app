package grpc

import (
	"net"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/config"
	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	config     config.Config
	logger     *logrus.Logger
}

func NewServer(config config.Config, logger *logrus.Logger) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		config:     config,
		logger:     logger,
	}
}

func (s *Server) RegisterTodoAppServiceV1(apiServer v1.TodoAppServiceV1Server) {
	v1.RegisterTodoAppServiceV1Server(s.grpcServer, apiServer)
}

func (s *Server) Start() error {
	listener, err := net.Listen(
		"tcp",
		net.JoinHostPort(s.config.Host, s.config.GrpcPort),
	)
	if err != nil {
		s.logger.Fatal("failed to start grpc server: %w", err)
		return err
	}

	s.logger.Infof("grpc server started on port, :%s", s.config.GrpcPort)
	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			s.logger.Fatal("failed to start grpc server: %w", err)
		}
	}()

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
	s.logger.Info("grpc server stopped...")
}
