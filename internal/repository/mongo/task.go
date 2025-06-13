package mongo

import (
	"context"
	"errors"
	"github.com/Dimoonevs/task-api/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepo struct {
	col *mongo.Collection
}

func NewTaskRepo(db *mongo.Database) *taskRepo {
	return &taskRepo{col: db.Collection("tasks")}
}

func (r *taskRepo) Create(t *model.Task) error {
	_, err := r.col.InsertOne(context.Background(), t)
	if err != nil {
		logrus.Errorf("mongo insert: %v", err)
	}
	return err
}

func (r *taskRepo) Get(id string) (*model.Task, error) {
	var res model.Task
	err := r.col.FindOne(context.Background(), bson.M{"_id": id}).Decode(&res)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return &res, err
}

func (r *taskRepo) List(status model.Status, page, pageSize int) ([]model.Task, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	cur, err := r.col.Find(context.Background(), filter, options.Find().SetSkip(int64((page-1)*pageSize)).SetLimit(int64(pageSize)))
	if err != nil {
		logrus.Errorf("mongo find: %v", err)
		return nil, err
	}
	defer cur.Close(context.Background())

	var tasks []model.Task
	if err := cur.All(context.Background(), &tasks); err != nil {
		logrus.Errorf("mongo cursor: %v", err)
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepo) Update(id string, t *model.Task) error {
	_, err := r.col.ReplaceOne(context.Background(), bson.M{"_id": id}, t)
	return err
}

func (r *taskRepo) Delete(id string) error {
	_, err := r.col.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
