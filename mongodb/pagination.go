package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Definisikan struktur untuk data yang akan disimpan di MongoDB
type Item struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func main() {
	// Membuat koneksi ke MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Mendapatkan koneksi ke koleksi MongoDB
	collection := client.Database("mydb").Collection("items")

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		// Mengambil parameter page dan ukuran halaman dari permintaan
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
		name := r.URL.Query().Get("name") // Menambahkan parameter "name" dari permintaan

		// Menentukan opsional nilai default
		if page == 0 {
			page = 1
		}
		if perPage == 0 {
			perPage = 10
		}

		var params string = ""

		// Menentukan filter dan opsi pagination
		filter := bson.M{}
		if name != "" {
			filter["name"] = bson.M{
				"$regex":   name,
				"$options": "i", // Opsi "i" untuk pencocokan case-insensitive
			}
			params = fmt.Sprintf("&%sname=%s", params, name)

		}
		options := options.Find().SetSkip(int64((page - 1) * perPage)).SetLimit(int64(perPage))

		var total int64 = 0
		count, err := collection.CountDocuments(context.Background(), filter)
		if err != nil {
			w.Write([]byte("Koneksi query salah"))
			return
		}
		if count > 0 {
			total = count
		} else {
			w.Write([]byte("Data Tidak di temukan"))
			return
		}

		// Menjalankan kueri untuk mengambil data dengan filter dan opsi pagination
		cursor, err := collection.Find(context.Background(), filter, options)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		var items []Item
		for cursor.Next(context.Background()) {
			var item Item
			if err := cursor.Decode(&item); err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}

		// Mengembalikan data dalam format JSON
		// w.Header().Set("Content-Type", "application/json")
		// if err := json.NewEncoder(w).Encode(items); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		var totalminbyperpage, lastPage int64
		totalminbyperpage = total - ((total / 10) * 10)

		if totalminbyperpage == 0 {
			lastPage = (total / 10)
		} else {
			lastPage = ((total / 10) + 1)
		}

		var first, last, next, prev string
		first, last, next, prev = "", "", "", ""

		first = "1"
		last = strconv.Itoa(int(lastPage))

		next = strconv.Itoa(int(page + 1))
		if int(page+1) >= int(lastPage) {
			next = strconv.Itoa(int(lastPage))
		}

		prev = strconv.Itoa(int(page - 1))
		if int(page) == 1 {
			prev = strconv.Itoa(int(page))
		}

		links := map[string]interface{}{
			"first": fmt.Sprintf("?page=%s%s", first, params),
			"last":  fmt.Sprintf("?page=%s%s", last, params),
			"next":  fmt.Sprintf("?page=%s%s", next, params),
			"prev":  fmt.Sprintf("?page=%s%s", prev, params),
		}

		informationpages := map[string]interface{}{
			"currentPage": page,
			"from":        ((page - 1) * 10) + 1,
			"lastPage":    lastPage,
			"perPage":     perPage,
			"to":          ((page - 1) * 10) + len(items),
			"total":       total,
		}

		pages := map[string]interface{}{
			"page": informationpages,
		}

		result := map[string]interface{}{
			"code":  200,
			"meta":  pages,
			"data":  items,
			"links": links,
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte(resp))
		return

		// resp, err := json.Marshal(items)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
		// w.Write([]byte(resp))
		// return
	})

	http.ListenAndServe(":8080", nil)
}
