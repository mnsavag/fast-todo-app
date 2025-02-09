package service

import (
	"context"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/model"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service/dto"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) CreateList(ctx context.Context, createListData dto.CreateList) (uint32, error) {
	panic("implement me")
}
func (s *Service) DeleteList(ctx context.Context, listId uint32) error {
	panic("implement me")
}
func (s *Service) GetList(ctx context.Context, listId uint32) (model.List, error) {
	panic("implement me")
}

func (s *Service) CreateItem(ctx context.Context, createItemData dto.CreateItem) (uint32, error) {
	panic("implement me")
}
func (s *Service) DeleteItem(ctx context.Context, itemId uint32) error {
	panic("implement me")
}
