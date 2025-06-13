package model

type Status string

const (
	StatusTodo         Status = "todo"
	StatusInProgress   Status = "in_progress"
	StatusInQA         Status = "in_qa"
	StatusReadyRelease Status = "ready_for_release"
	StatusDone         Status = "done"
)

var AllStatuses = map[Status]struct{}{
	StatusTodo:         {},
	StatusInProgress:   {},
	StatusInQA:         {},
	StatusReadyRelease: {},
	StatusDone:         {},
}

type Task struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      Status `json:"status" bson:"status"`
	CreatedAt   uint64 `json:"created_at" bson:"created_at"`
	UpdatedAt   uint64 `json:"updated_at" bson:"updated_at"`
}

type UpdateTask struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *Status `json:"status,omitempty"`
}
