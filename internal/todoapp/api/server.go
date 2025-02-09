package api

import (
	"context"
	"errors"

	"github.com/mnsavag/fast-todo-app.git/internal/lib/log"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service/dto"
	"github.com/mnsavag/fast-todo-app.git/pkg/api"
	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

const logFieldMethod = "method"

type Server struct {
	v1.UnimplementedTodoAppServiceV1Server
	logger         *logrus.Logger
	todoAppService TodoAppService
}

func NewServer(logger *logrus.Logger, todoAppService TodoAppService) *Server {
	return &Server{
		logger:         logger,
		todoAppService: todoAppService,
	}
}

func (s *Server) CreateList(ctx context.Context, in *v1.CreateListRequest) (*v1.CreateListResponse, error) {
	l := s.logger.WithField(logFieldMethod, "CreateList")

	mappedRequest := fromCreateListRequest(in)

	id, err := s.todoAppService.CreateList(ctx, mappedRequest)
	if err != nil {
		l.WithField(log.FieldError, err.Error())
		return nil, toGrpcError(err)
	}

	return &v1.CreateListResponse{
		ListId: id,
	}, nil
}

func (s *Server) DeleteList(ctx context.Context, in *v1.DeleteListRequest) (*emptypb.Empty, error) {
	l := s.logger.WithField(logFieldMethod, "DeleteList")

	err := s.todoAppService.DeleteList(ctx, in.ListId)
	if err != nil {
		l.WithField(log.FieldError, err.Error())
		return &emptypb.Empty{}, toGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetList(ctx context.Context, in *v1.GetListRequest) (*v1.GetListResponse, error) {
	l := s.logger.WithField(logFieldMethod, "GetList")

	list, err := s.todoAppService.GetList(ctx, in.ListId)
	if err != nil {
		if errors.Is(err, service.ErrListNotFound) {
			return &v1.GetListResponse{}, nil
		}

		l.WithField(log.FieldError, err.Error())
		return nil, toGrpcError(err)
	}

	return &v1.GetListResponse{
		List: &api.List{
			Id:          list.Id,
			Title:       list.Title,
			Description: list.Description,
		},
	}, nil
}

func (s *Server) CreateItem(ctx context.Context, in *v1.CreateItemRequest) (*v1.CreateItemResponse, error) {
	l := s.logger.WithField(logFieldMethod, "CreateItem")

	id, err := s.todoAppService.CreateItem(
		ctx,
		dto.CreateItem{
			Title:       in.Title,
			Description: in.Description,
		},
	)
	if err != nil {
		l.WithField(log.FieldError, err.Error())
		return nil, toGrpcError(err)
	}

	return &v1.CreateItemResponse{
		ItemId: id,
	}, nil
}

func (s *Server) DeleteItem(ctx context.Context, in *v1.DeleteItemRequest) (*emptypb.Empty, error) {
	l := s.logger.WithField(logFieldMethod, "DeleteItem")

	err := s.todoAppService.DeleteItem(ctx, in.ItemId)
	if err != nil {
		l.WithField(log.FieldError, err.Error())
		return &emptypb.Empty{}, toGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
