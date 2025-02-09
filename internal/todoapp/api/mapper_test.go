package api

import (
	"testing"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service/dto"
	v1 "github.com/mnsavag/fast-todo-app.git/pkg/api/v1"

	"github.com/stretchr/testify/assert"
)

func TestFromCreateListRequest(t *testing.T) {
	tests := []struct {
		name string
		in   *v1.CreateListRequest
		want dto.CreateList
	}{
		{
			name: "all field",
			in: &v1.CreateListRequest{
				Title:       "title",
				Description: "description",
			},
			want: dto.CreateList{
				Title:       "title",
				Description: "description",
			},
		},
		{
			name: "one field",
			in: &v1.CreateListRequest{
				Title: "title",
			},
			want: dto.CreateList{
				Title:       "title",
				Description: "",
			},
		},
	}

	for _, test := range tests {
		actual := fromCreateListRequest(test.in)
		assert.Equal(t, test.want, actual)
	}
}
