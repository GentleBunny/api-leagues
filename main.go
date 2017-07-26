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
			var user User
			if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Mail, &user.Age, &user.Gender); err != nil {
				fmt.Printf("Error in rows.Scan ", err.Error())
				return
			}
			fmt.Printf("%d %s %s %s %d %s \n", user.Id, user.Name, user.Username, user.Mail, user.Age, user.Gender)
		}
		c.JSON(200, "ok")
	})

	router.Run()
}
