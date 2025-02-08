package api

import (
	"context"
	"fmt"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/config"
	"github.com/sirupsen/logrus"

	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	v1.UnimplementedTodoAppServiceV1Server
	cfg    config.Config
	logger *logrus.Logger
}

func NewServer(cfg config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) CreateItem(context.Context, *v1.CreateItemRequest) (*v1.CreateItemResponse, error) {
	fmt.Println("CreateItem HelloWorld")
	panic("implement me")
}
func (s *Server) CreateList(context.Context, *v1.CreateListRequest) (*v1.CreateListResponse, error) {
	panic("implement me")
}
func (s *Server) DeleteItem(context.Context, *v1.DeleteItemRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
func (s *Server) DeleteList(context.Context, *v1.DeleteListRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
func (s *Server) GetList(context.Context, *v1.GetListRequest) (*v1.GetListResponse, error) {
	panic("implement me")
}
