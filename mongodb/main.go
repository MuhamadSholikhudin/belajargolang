package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Andy", 3})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Yulian", 4})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	csr, err := db.Collection("student").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("row  :", row, reflect.TypeOf(row), row.Grade)
		result = append(result, row)
	}
	fmt.Println(len(result))
	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}
}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "ANG John"}
	// var changes = student{"John", 500}
	var changes = bson.M{"Grade": 550}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	// var changes = student{"grade", 400}
	// _, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println("Update success!")
	// filter := bson.D{{Key: "name", Value: "grade"}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "grade", Value: 400}}}}

	// id, _ := primitive.ObjectIDFromHex("63d089f87d7c0dab5108c611")
	// filter := bson.D{{Key: "_id", Value: id}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "grade", Value: 5}}}}
	// _, err = db.Collection("student").UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Update success!")

}

func remove() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	var selector = bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Remove success!")
}

func main() {
	// insert()
	update()
	find()
	// remove()
}
