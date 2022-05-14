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
	Time time.Time
	Round int
}

type PlayerResults struct {
	ID primitive.ObjectID
	Alias string
	Points int 
}

type GameResult struct {
	GameID primitive.ObjectID
	Winner primitive.ObjectID
	Players [2]PlayerResults	
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

	game.Time = time.Now()

	result, err := Games.InsertOne(ctx, game)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func FinalizeGameResult(gameResult GameResult) {
	player1 := gameResult.Players[0]
	player2 := gameResult.Players[1]

	UpdateGame(gameResult)

	if player1.ID == gameResult.Winner {
		IncrementPlayerRecord(player1.ID, true, player1.Points)
		IncrementPlayerRecord(player2.ID, false, player2.Points)
	} else {
		IncrementPlayerRecord(player1.ID, false, player1.Points)
		IncrementPlayerRecord(player2.ID, true, player2.Points)
	}
}

func UpdateGame(gameResult GameResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := Games.UpdateOne(ctx, bson.M{"_id": gameResult.GameID}, bson.M{"$set": bson.M{"winner": gameResult.Winner, "players": gameResult.Players}})
	if err != nil {
		log.Fatal(err)
	}
}

func GetPlayerMatchHistory(id string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Print(objID)

	cur, err := Games.Find(ctx, bson.M{"players": bson.M{"$elemMatch": bson.M{"id": objID}}})
	if err !=nil {
		log.Fatal(err)
	}

	var games []bson.M 
	if err := cur.All(ctx, &games); err != nil {fmt.Println(err)}

	fmt.Print(games)

	return games
}