package service

import (
	"errors"
	"github.com/Dimoonevs/task-api/internal/model"
	"github.com/Dimoonevs/task-api/internal/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) Create(input *model.Task) error {
	if err := s.validateInput(input); err != nil {
		logrus.Errorf("create: validation %v", err)
		return err
	}

	input.ID = uuid.NewString()
	now := uint64(time.Now().UnixMilli())
	input.CreatedAt, input.UpdatedAt = now, now

	return s.repo.Create(input)
}

func (s *TaskService) Get(id string) (*model.Task, error) {
	return s.repo.Get(id)
}

func (s *TaskService) List(status model.Status, page, pageSize int) ([]model.Task, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	if status != "" {
		if _, ok := model.AllStatuses[status]; !ok {
			logrus.Errorf("list: invalid status filter %s", status)
			return nil, errors.New("invalid status filter")
		}
	}
	return s.repo.List(status, page, pageSize)
}

func (s *TaskService) Update(id string, in *model.Task) error {
	if err := s.validateInput(in); err != nil {
		logrus.Errorf("update %s: validation %v", id, err)
		return err
	}
	existing, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("task not found")
	}

	in.ID = id
	in.CreatedAt = existing.CreatedAt
	in.UpdatedAt = uint64(time.Now().UnixMilli())

	return s.repo.Update(id, in)
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *TaskService) validateInput(t *model.Task) error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if _, ok := model.AllStatuses[t.Status]; !ok {
		return errors.New("unknown status")
	}
	return nil
}
