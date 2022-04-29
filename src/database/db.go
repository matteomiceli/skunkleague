package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name string
	// wins + losses
	// 
}

var Players *mongo.Collection 
var Games *mongo.Collection 
var client *mongo.Client

func Init() {
	c, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	client = c

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Players = client.Database("skunkleague").Collection("players")
	Games = client.Database("skungleague").Collection("games")

}

func GetAllPlayers() []primitive.M {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := Players.Find(ctx, bson.D{{}})
	if err !=nil {
		log.Fatal(err)
	}

	var players []bson.M
	
	if err = cur.All(ctx, &players); err != nil {
		log.Fatal(err)
	}

	return players
}

func AddNewPlayer(name string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := Players.InsertOne(ctx, Player{Name: name})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Disconnect(ctx)
}