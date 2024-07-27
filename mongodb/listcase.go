package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctxx = func() context.Context {
	return context.Background()
}()

type listcase struct {
	Id                  int       `bson:"id"`
	Investigator        string    `bson:"investigator"`
	No_case             int       `bson:"no_case"`
	Case_name           string    `bson:"case_name"`
	Reporter_victim     string    `bson:"reporter_victim"`
	Reported            string    `bson:"reporter"`
	Building            string    `bson:"building"`
	Category_indication string    `bson:"category_indication"`
	Source              string    `bson:"source"`
	Report_entry_date   time.Time `bson:"report_entry_date"`
	Progress            string    `bson:"progress"`
	Results             string    `bson:"results"`
	Date_of_completion  time.Time `bson:"date_of_completion"`
	Sanction_date       time.Time `bson:"sanction_date"`
	Notes               string    `bson:"notes"`
	Created_at          time.Time `bson:"created_at"`
	Updated_at          time.Time `bson:"updated_at"`
}

func connectls() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctxx)
	if err != nil {
		return nil, err
	}

	return client.Database("hi"), nil
}

func insertmdb() {
	db, err := connectls()
	if err != nil {
		log.Fatal(err.Error())
	}

	// _, err = db.Collection("list_cases").InsertOne(ctxx, listcase{1, "Andy Yulian", 2, "Pencurian HP", "Jems", "Mita", "Main Office", "Pencurian", "Rekaman CCTV", time.Now(), "ON-GOING", "DI Lanjutkan Ranah Hukum", time.Now(), time.Now(), "Harus Teliti Dalam Kasus", time.Now(), time.Now()})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	_, err = db.Collection("list_cases").InsertOne(ctxx, listcase{2, "Wafa", 2, "Pencurian HP", "Ifa", "Rizal", "Main Office", "Membentak Atasan", "Saksi Karyawan Setempat", time.Now(), "PROCESS", "DI BErikan SP", time.Now(), time.Now(), "Di Sp dan di NAsihati", time.Now(), time.Now()})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func findmdb() {
	db, err := connectls()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("list_cases").Find(ctxx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctxx)

	result := make([]listcase, 0)
	for csr.Next(ctxx) {
		var row listcase
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}
	fmt.Println("result  :", result)

	// if len(result) > 0 {
	// 	fmt.Println("Name  :", result[0].Name)
	// 	fmt.Println("Grade :", result[0].Grade)
	// }
}

func updatemdb() {
	db, err := connectls()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"id": 2}
	var changes = listcase{2, "Rizma", 2, "Pencurian HP", "Isna", "Diah", "Main Office", "Membentak Atasan Dengan Keras", "Saksi Karyawan Setempat dan atasan korea", time.Now(), "SELESAI", "DI Berikan SP III", time.Now(), time.Now(), "Di Sp dan di NAsihati", time.Now(), time.Now()}
	_, err = db.Collection("list_cases").UpdateOne(ctxx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Update success!")

}

func removemdb() {
	db, err := connectls()
	if err != nil {
		log.Fatal(err.Error())
	}
	var selector = bson.M{"id": 2}
	_, err = db.Collection("list_cases").DeleteOne(ctxx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Remove success!")

}

func main() {
	insertmdb()
	// findmdb()
	// updatemdb()
	// findmdb()
	// removemdb()
	// findmdb()
}
