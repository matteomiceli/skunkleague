package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Players *mongo.Collection 
var Games *mongo.Collection 
var DbContext context.Context
var client *mongo.Client

func Init() {
	c, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	client = c

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	DbContext = ctx

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Players = client.Database("skunkleague").Collection("players")
	Games = client.Database("skungleague").Collection("games")

}

func GetAllPlayers() []primitive.M {
	cur, err := Players.Find(DbContext, bson.D{{}})
	if err !=nil {
		log.Fatal(err)
	}

	var players []bson.M
	
	if err = cur.All(DbContext, &players); err != nil {
		log.Fatal(err)
	}

	return players
}

func Disconnect() {
	client.Disconnect(DbContext)
}