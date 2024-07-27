package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListCase struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Investigator        []string           `json:"investigator" bson:"investigator" validate:"required"`
	No_case             int                `json:"no_case" bson:"no_case" validate:"required"`
	Case_name           string             `json:"case_name" bson:"case_name" validate:"required"`
	Reporter_victim     []string           `json:"reporter_victim" bson:"reporter_victim" validate:"required"`
	Reported            []string           `json:"reported" bson:"reported" validate:"required"`
	Building            string             `json:"building" bson:"building" validate:"required"`
	Category_indication string             `json:"category_indication" bson:"category_indication" validate:"required"`
	Source              string             `json:"source" bson:"source" validate:"required"`
	Report_entry_date   time.Time          `json:"report_entry_date" bson:"report_entry_date" validate:"required"`
	Progress            string             `json:"progress" bson:"progress" validate:"required"`
	Results             string             `json:"results" bson:"results" validate:"required"`
	Date_of_completion  time.Time          `json:"date_of_completion" bson:"date_of_completion" validate:"required"`
	Sanction_date       time.Time          `json:"sanction_date" bson:"sanction_date" validate:"required"`
	Notes               string             `json:"notes" bson:"notes" validate:"required"`
	Document            string             `json:"document" bson:"document" validate:"required"`
	Created_at          time.Time          `json:"created_at" bson:"created_at" validate:"required"`
	Updated_at          time.Time          `json:"updated_at" bson:"updated_at" validate:"required"`
}

func main() {
	fmt.Println("OKE BANGET")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api.listcase.v1/fileupload", Fileupload)
	http.HandleFunc("/api.listcase.v1/open-file", OpenFile)
	fmt.Println("Server listening on port 1115...")
	http.ListenAndServe(":1115", nil)
}

var ctxx = func() context.Context {
	return context.Background()
}()

func ConnMdHi() (*mongo.Database, error) {
	ctxx := context.Background()
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

func Fileupload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	db, err := ConnMdHi()
	if err != nil {
		log.Fatal(err.Error())
	}
	textData := r.FormValue("textData")
	fmt.Println(textData)
	id, err := primitive.ObjectIDFromHex(textData)
	if err != nil {
		log.Fatal(err)
	}
	uploadedFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()
	csr, err := db.Collection("list_cases").Find(ctxx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctxx)
	result1 := make([]ListCase, 0)
	for csr.Next(ctxx) {
		var row1 ListCase
		err := csr.Decode(&row1)
		if err != nil {
			log.Fatal(err.Error())
		}
		result1 = append(result1, row1)
	}
	fileremove := fmt.Sprintf("files-caselist/%s", result1[0].Document)
	e := os.Remove(fileremove)
	if e != nil {
		fmt.Println("Data File Kasus tidak di temukan")
	}
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileLocation := filepath.Join(dir, "files-caselist/", fileHeader.Filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	update := bson.M{"$set": bson.M{"document": fileHeader.Filename}}
	var selector = bson.M{"_id": id}
	_, err = db.Collection("list_cases").UpdateOne(ctxx, selector, update)
	if err != nil {
		result := map[string]interface{}{
			"code":    400,
			"message": fmt.Sprintf("Data File tidak dapat di upload %T", err.Error()),
			"data":    fileHeader.Filename,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}
	result := map[string]interface{}{
		"code":    200,
		"message": "Data File Berhasil di upload",
		"data":    fileHeader.Filename,
	}
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
	return

	// // Respond with a success message
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "File and text uploaded successfully. File: %s, Text: %s", fileHeader.Filename, textData)
}

func OpenFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()

	// Get the value of the "search" parameter
	file := query.Get("file")

	fmt.Println(file)

	location := fmt.Sprintf("files-caselist/%s", file)

	fileBytes, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("File Kasus Tidak di temukan")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
