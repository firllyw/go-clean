package main

import (
	"log"

	"github.com/goclean/internal/controller"
	"github.com/goclean/internal/repository/mongodb"
	"github.com/goclean/internal/services"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "members"
	database   = "members"
)

func main() {
	mongoConn := "mongodb://localhost:27017"
	if mongoConn == "" {
		panic("MONGO_CONN is empty")
	}

	client, err := mongo.Connect(nil, options.Client().ApplyURI(mongoConn))
	if err != nil {
		panic("Failed connecting to mongodb, err: " + err.Error())
	}
	e := echo.New()
	mongoColl := client.Database(database).Collection(collection)
	repo := mongodb.NewMemberRepo(mongoColl)
	svc := services.NewMemberService(repo)
	controller.NewMemberHandler(e, svc)
	log.Fatal(e.Start("localhost:8080"))

}
