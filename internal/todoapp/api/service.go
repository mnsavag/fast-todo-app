package api

import (
	"context"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/model"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service/dto"
)

type TodoAppService interface {
	ListExecutor
	ItemExecutor
}

type ListExecutor interface {
	CreateList(ctx context.Context, createListData dto.CreateList) (uint32, error)
	DeleteList(ctx context.Context, listId uint32) error
	GetList(ctx context.Context, listId uint32) (model.List, error)
}

type ItemExecutor interface {
	CreateItem(ctx context.Context, createItemData dto.CreateItem) (uint32, error)
	DeleteItem(ctx context.Context, itemId uint32) error
}
