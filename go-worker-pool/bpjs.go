package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnStringbpjs   = "root:@tcp(127.0.0.1:3306)/hrd"
	dbMaxIdleConnsbpjs = 4
	dbMaxConnsbpjs     = 100
	totalWorkerbpjs    = 50
	csvFilebpjs        = "bpjs.csv"
	dataHeadersbpjs    = []string{
		"number_of_employees",
		"bpjs_ketenagakerjaan",
		"date_bpjs_ketenagakerjaan",
		"bpjs_kesehatan",
		"date_bpjs_kesehatan",
	}
)

// CREATE DATABASE IF NOT EXISTS test;
// USE test;
// CREATE TABLE IF NOT EXISTS domain (
//     GlobalRank int,
//     TldRank int,
//     Domain varchar(255),
//     TLD varchar(255),
//     RefSubNets int,
//     RefIPs int,
//     IDN_Domain varchar(255),
//     IDN_TLD varchar(255),
//     PrevGlobalRank int,
//     PrevTldRank int,
//     PrevRefSubNets int,
//     PrevRefIPs int
// );

func main() {
	start := time.Now()

	dbbpjs, err := openDbConnectionbpjs()
	if err != nil {
		log.Fatal(err.Error())
	}

	csvReaderbpjs, csvFilebpjs, err := openCsvFilebpjs()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFilebpjs.Close()

	jobsdbbpjs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	// fmt.Println(jobs)
	go dispatchWorkersbpjs(dbbpjs, jobsdbbpjs, wg)
	readCsvFilePerLineThenSendToWorkerbpjs(csvReaderbpjs, jobsdbbpjs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func openDbConnectionbpjs() (*sql.DB, error) {
	log.Println("=> open db connection")

	dbbpjs, err := sql.Open("mysql", dbConnStringbpjs)
	if err != nil {
		return nil, err
	}

	dbbpjs.SetMaxOpenConns(dbMaxConnsbpjs)
	dbbpjs.SetMaxIdleConns(dbMaxIdleConnsbpjs)

	return dbbpjs, nil
}

func openCsvFilebpjs() (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(csvFilebpjs)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file majestic_million.csv tidak ditemukan. silakan download terlebih dahulu di https://blog.majestic.com/development/majestic-million-csv-daily")
		}

		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

func dispatchWorkersbpjs(dbbpjs *sql.DB, jobsbpjs <-chan []interface{}, wg *sync.WaitGroup) {
	//db *sql.DB
	for workerIndex := 0; workerIndex <= totalWorkerbpjs; workerIndex++ {
		go func(workerIndex int, dbbpjs *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobsbpjs {
				doTheJobbpjs(workerIndex, counter, dbbpjs, job)
				wg.Done()
				counter++
			}
		}(workerIndex, dbbpjs, jobsbpjs, wg)
	}
}

func readCsvFilePerLineThenSendToWorkerbpjs(csvReader *csv.Reader, jobsbpjs chan<- []interface{}, wg *sync.WaitGroup) {
	isHeader := true
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if isHeader {
			isHeader = false
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}
		wg.Add(1)
		jobsbpjs <- rowOrdered
	}
	close(jobsbpjs)
}

func doTheJobbpjs(workerIndex, counter int, db *sql.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			conn, err := db.Conn(context.Background())
			// query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)",
			// 	strings.Join(dataHeaders, ","),
			// 	strings.Join(generateQuestionsMarkbpjs(len(dataHeaders)), ","),
			// )

			queryupdate := fmt.Sprintf("UPDATE employees SET %s = ? WHERE number_of_employees = '%s' ",
				strings.Join(dataHeadersbpjs, " = ?,"),
				values[0],
			)
			// fmt.Println(queryupdate)

			// for i := 0; i < len(values); i++ {
			// 	fmt.Print(i, " ")
			// }
			// fmt.Println(values)

			_, err = conn.ExecContext(context.Background(), queryupdate, values...)
			if err != nil {
				log.Fatal(err.Error())
			}

			// _, err = conn.ExecContext(context.Background(), query, values...)
			// if err != nil {
			// 	log.Fatal(err.Error())
			// }

			err = conn.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	if counter%100 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter, "data")
	}
}

func generateQuestionsMarkbpjs(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, "?")
	}
	return s
}
