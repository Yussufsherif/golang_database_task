package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.Run()
}

type postgres struct {
	DB *sql.DB
}

func (a *postgres) CreateConnection() {
	dbUname := "postgres"
	dbPass := "yushkongpost"
	dbHost := "localhost"
	dbName := " clients"
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUname, dbPass, dbHost, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}
