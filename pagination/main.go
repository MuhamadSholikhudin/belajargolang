package main

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/gin-gonic/gin"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pagination")
		defer db.Close()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}

		var version string

		err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

		if err2 != nil {
			c.JSON(500, gin.H{
				"error": err2,
			})
		}

		//fmt.Println(version)
		c.JSON(200, gin.H{
			"version_mysql": version,
		})
	})

	r.GET("/list", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pagination")
		defer db.Close()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		name, checkName := c.GetQuery("name")
		sql := "SELECT count(*) FROM cities"
		if checkName != false {
			sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", sql, name)
		}

		var total int64

		db.QueryRow(sql).Scan(&total)
		//c.String(200, "total %v", total)

		//paging
		sqlPaging := "SELECT * FROM cities"
		//paging
		if checkName != false {
			sqlPaging = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", sqlPaging, name)
		}
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "2"))

		sqlPaging = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlPaging, perPage, (page-1)*perPage)
		CityRows, err := db.Query(sqlPaging)

		defer CityRows.Close()
		cityList := City{}

		result := []City{}
		for CityRows.Next() {

			var id int

			var name string

			var population int
			err = CityRows.Scan(&id, &name, &population)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err,
				})
				return
			}

			cityList.Id = id

			cityList.Name = name

			cityList.Population = population

			result = append(result, cityList)

		}
		c.JSON(200, gin.H{
			"data":         result,
			"total_record": total,
			"current_page": page,
			"total_page":   math.Ceil(float64(total / int64(perPage))),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
