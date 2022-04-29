package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID primitive.ObjectID  `bson:"_id,omitempty"`
	Alias string
	FirstName string
	LastName string
	Wins int
	Losses int
	Points int
}

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

func CreatePlayer(player Player) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Players.InsertOne(ctx, player)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func IncrementPlayerRecord(id primitive.ObjectID, isWinner bool, points int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if isWinner {

		_, err := Players.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$inc": bson.M{"wins": 1, "points": points}})
		if err != nil {
			panic(err)
		}

	} else {
		_, err := Players.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$inc": bson.M{"losses": 1, "points": points}})
		if err != nil {
			panic(err)
		}

	}
}