package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// INIT POSTGRES
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/ultra_manager?sslmode=disable")
	if err != nil {
		fmt.Println("Can't connect to postgres:", err.Error())
		return
	}
	router := gin.Default()

	// TEST POSTGRES USAGE
	router.GET("/testUser", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM app_user")
		if err != nil {
			fmt.Println("Can't connect to postgres:", err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var username string
			var mail string
			var age int
			var gender string
			if err := rows.Scan(&id, &name, &username, &mail, &age, &gender); err != nil {
				fmt.Printf("Error in rows.Scan ", err.Error())
				return
			}
			fmt.Printf("%d %s %s %s %d %s \n", id, name, username, mail, age, gender)
		}
		c.JSON(200, "ok")
	})

	router.Run()
}
