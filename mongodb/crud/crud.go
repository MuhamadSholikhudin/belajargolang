package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctxmdb = func() context.Context {
	return context.Background()
}()

type studentmdb struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connectmdb() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctxmdb)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insertmdb() {
	db, err := connectmdb()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctxmdb, studentmdb{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctxmdb, studentmdb{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func findmdb() {
	db, err := connectmdb()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctxmdb, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctxmdb)

	result := make([]studentmdb, 0)
	for csr.Next(ctxmdb) {
		var row studentmdb
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
		fmt.Println("result  :", result)
	}

	// if len(result) > 0 {
	// 	fmt.Println("Name  :", result[0].Name)
	// 	fmt.Println("Grade :", result[0].Grade)
	// }
}

func updatemdb() {
	db, err := connectmdb()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "John Wick"}
	var changes = studentmdb{"John Wick", 3}
	_, err = db.Collection("student").UpdateOne(ctxmdb, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func removemdb() {
	db, err := connectmdb()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctxmdb, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

func main() {
	// insertmdb()
	findmdb()
	// updatemdb()
	// findmdb()
	// removemdb()
	// findmdb()

}
