package database

import (
	"testing"
	"web-based-todo-list-backend/models"
)

func TestDatabase_GetTodoList(t *testing.T) {
	type fields struct {
		todolist models.DataResponse
	}

	tests := []struct {
		name   string
		fields fields
		want   *models.DataResponse
	}{
		{name: "when the todolist is in memory then get all",
			fields: fields{todolist: models.DataResponse{
				1: models.Todo{
					ID:          1,
					Description: "Dummy todo",
				},
			}},
			want: &models.DataResponse{
				1: models.Todo{
					ID:          1,
					Description: "Dummy todo",
				},
			},
		},
		{name: "when the todolist is not in memory then get all",
			fields: fields{todolist: models.DataResponse{}},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// need database

		})

	}
}
