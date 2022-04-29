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
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string
	// wins + losses
	// total points
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Players = client.Database("skunkleague").Collection("players")
	Games = client.Database("skungleague").Collection("games")
}

// Player functions
func GetAllPlayers() []Player{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cur, err := Players.Find(ctx, bson.D{{}})
	if err !=nil {
		log.Fatal(err)
	}

	var players []Player
	
	if err = cur.All(ctx, &players); err != nil {
		log.Fatal(err)
	}

	return players
}

func GetPlayerById(id string) Player {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	cur := Players.FindOne(ctx, bson.M{"_id": objID})

	var player Player
	
	if err := cur.Decode(&player); err != nil {fmt.Println(err)}

	return player
}

func AddNewPlayer(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Players.InsertOne(ctx, Player{Name: name})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// Game functions 
func GetAllGames() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := Games.Find(ctx, bson.D{{}})
	if err !=nil {
		log.Fatal(err)
	}

	var games []bson.M
	
	if err = cur.All(ctx, &games); err != nil {
		log.Fatal(err)
	}

	return games 
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Disconnect(ctx)
}