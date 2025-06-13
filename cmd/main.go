package main

import (
	"context"
	"flag"
	"github.com/Dimoonevs/task-api/internal/handler"
	repoMongo "github.com/Dimoonevs/task-api/internal/repository/mongo"
	"github.com/Dimoonevs/task-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/vharitonsky/iniflags"
)

var (
	mongoConnectionString = flag.String("mongoConnectionString", "mongodb://mongo:27017", "MongoDB connection URI")
	mongoDatabase         = flag.String("mongoDatabase", "taskdb", "MongoDB database name")

	httpServerListenAddr = flag.String("httpServerListenAddr", ":8080", "Address the HTTP server should listen on")
)

func main() {
	iniflags.Parse()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*mongoConnectionString))
	if err != nil {
		log.Fatalf("mongo: %v", err)
	}
	db := client.Database(*mongoDatabase)

	repo := repoMongo.NewTaskRepo(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	router := gin.New()
	h.Register(router)

	if err = router.Run(*httpServerListenAddr); err != nil {
		log.Fatalf("server: %v", err)
	}
}
