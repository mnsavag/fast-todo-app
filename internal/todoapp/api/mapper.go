package api

import (
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service/dto"
	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"
)

func fromCreateListRequest(in *v1.CreateListRequest) dto.CreateList {
	return dto.CreateList{
		Title:       in.Title,
		Description: in.Description,
	}
}
