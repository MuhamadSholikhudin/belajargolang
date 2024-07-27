package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctxmd = func() context.Context {
	return context.Background()
}()

type studentmd struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connectmd() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctxmd)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connectmd()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctxmd, studentmd{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctxmd, studentmd{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func main() {
	insert()
}
