package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Players [2]PlayerResults 
	Winner primitive.ObjectID `bson:"winner,omitempty"`
	Time primitive.Timestamp
	Round int
}

type PlayerResults struct {
	ID primitive.ObjectID
	Points int 
}

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

func CreateGame(game Game) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Games.InsertOne(ctx, game)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}