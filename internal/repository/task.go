package repository

import "github.com/Dimoonevs/task-api/internal/model"

type TaskRepository interface {
	Create(t *model.Task) error
	Get(id string) (*model.Task, error)
	List(status model.Status, page, pageSize int) ([]model.Task, error)
	Update(id string, t *model.Task) error
	Delete(id string) error
}
