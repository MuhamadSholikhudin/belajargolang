package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/xuri/excelize/v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnString4   = "root@/test"
	dbMaxIdleConns4 = 4
	dbMaxConns4     = 100
	totalWorker4    = 5
	xlsxFile        = "majestic_million.xlsx"
	dataHeaders4    = []string{
		"GlobalRank",
		"TldRank",
		"Domain",
		"TLD",
		"RefSubNets",
		"RefIPs",
		"IDN_Domain",
		"IDN_TLD",
		"PrevGlobalRank",
		"PrevTldRank",
		"PrevRefSubNets",
		"PrevRefIPs",
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

	db, err := openDbConnectionExcel()
	if err != nil {
		log.Fatal(err.Error())
	}

	excelReader, excelFile, err := openExcelFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer excelFile.Close()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	// fmt.Println(jobs)
	go dispatchWorkersExcel(db, jobs, wg)
	readCsvFilePerLineThenSendToWorkerExcel(excelReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func openDbConnectionExcel() (*sql.DB, error) {
	log.Println("=> open db connection")

	db, err := sql.Open("mysql", dbConnString4)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns4)
	db.SetMaxIdleConns(dbMaxIdleConns4)

	return db, nil
}

func openExcelFile() (*excelize.File, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(xlsxFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file majestic_million.csv tidak ditemukan. silakan download terlebih dahulu di https://blog.majestic.com/development/majestic-million-csv-daily")
		}
		return nil, nil, err
	}

	reader, _ := excelize.OpenFile(xlsxFile)
	return reader, f, nil
}

func dispatchWorkersExcel(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	//db *sql.DB
	for workerIndex := 0; workerIndex <= totalWorker4; workerIndex++ {
		go func(workerIndex int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJobExcel(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func readCsvFilePerLineThenSendToWorkerExcel(excelReader *excelize.File, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		// rows, err := excelReader.Rows("Sheet One")
		// if err != nil {
		// 	if err == io.EOF {
		// 		err = nil
		// 	}
		// 	break
		// }

		rowOrdered := make([]interface{}, 0)
		/*
			row, err := rows.Columns()
			if err != nil {
				fmt.Println(err)
			}
			if len(dataHeaders4) == 0 {
				dataHeaders4 = row
				continue
			}
			for _, colCell := range row {
				rowOrdered = append(rowOrdered, colCell)
			}

		*/

		// if len(dataHeaders4) == 0 {
		// 	dataHeaders4 = row
		// 	continue
		// }
		sheet1Name := "Sheet One"

		var mulai int
		for index, _ := range excelReader.GetRows(sheet1Name) {

			mulai = index + 1

			a := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))
			b := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("B%d", mulai))
			c := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("C%d", mulai))
			d := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("D%d", mulai))
			e := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("E%d", mulai))
			f := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("F%d", mulai))
			g := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("G%d", mulai))
			h := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("H%d", mulai))
			i := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("I%d", mulai))
			j := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("J%d", mulai))
			k := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("K%d", mulai))
			l := excelReader.GetCellValue(sheet1Name, fmt.Sprintf("L%d", mulai))

			rowOrdered = []interface{}{a, b, c, d, e, f, g, h, i, j, k, l}
			wg.Add(1)
			jobs <- rowOrdered
		}

	}
	close(jobs)

}

func doTheJobExcel(workerIndex, counter int, db *sql.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			// for i := 0; i < len(values); i++ {
			// 	fmt.Print(i, " ")
			// }

			conn, err := db.Conn(context.Background())
			query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)",
				strings.Join(dataHeaders4, ","),
				strings.Join(generateQuestionsMarkExcel(len(dataHeaders4)), ","),
			)

			// query := fmt.Sprintf("UPDATE domain SET %s = ? WHERE GlobalRank = %s",
			// 	strings.Join(dataHeaders4, " = ?,"),
			// 	values[0],
			// )
			// fmt.Println(query)

			// for i := 0; i < len(values); i++ {
			// 	fmt.Print(i, " ")
			// }
			// fmt.Println(values)

			_, err = conn.ExecContext(context.Background(), query, values...)
			if err != nil {
				log.Fatal(err.Error())
			}

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

func generateQuestionsMarkExcel(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, "?")
	}
	return s
}
